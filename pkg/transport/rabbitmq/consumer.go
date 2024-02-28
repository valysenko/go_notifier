package rabbitmq

import (
	"context"
	"fmt"
	"go_notifier/internal/common"
	"runtime"

	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"golang.org/x/sync/errgroup"
)

const maxRetries = 3 // 1 attmpt + 3 retries
const backoffBase = 1
const backoffCoefficient = 3
const numOfWorkers = 10

type ConsumersMap map[string]*Consumer

func NewConsumersMap(conn *amqp.Connection, l log.FieldLogger, hm MessageHandlersMap) ConsumersMap {
	consumers := make(ConsumersMap)
	for queue, handlers := range hm {
		for _, h := range handlers {
			consumers[queue] = NewConsumer(conn, queue, h, l)
		}
	}

	return consumers
}

func (cm ConsumersMap) Close() {
	for _, consumer := range cm {
		consumer.Close()
	}
}

/*
* map key is a queue name
* one queue can have several handlers
* one consumer has one handler and runs in a separate goroutine
 */
type MessageHandlersMap map[string][]MessageHandler

type MessageHandler interface {
	Handle(ctx context.Context, b []byte) *HandlerError
}

type Consumer struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	handler   MessageHandler
	logger    log.FieldLogger
	sleeper   *SecondsSleeper
	QueueName string
}

func NewConsumer(conn *amqp.Connection, queueName string, handler MessageHandler, logger log.FieldLogger) *Consumer {
	consumer := &Consumer{
		conn:      conn,
		QueueName: queueName,
		handler:   handler,
		logger:    logger,
		sleeper:   NewSecondsSleeper(backoffBase, backoffCoefficient),
	}

	return consumer
}

func (сon *Consumer) Listen(ctx context.Context) error {
	ch, err := сon.conn.Channel()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	сon.channel = ch
	defer ch.Close()

	logFields := log.Fields{"queue": сon.QueueName}

	exchangeType := common.QueueExchangeType[сon.QueueName]
	if exchangeType == common.DelayedExchangeType {
		err := сon.declareQueueWithDelayedExchange(ch)
		if err != nil {
			return err
		}
	} else {
		err := сon.declareQueueWithDirectExchange(ch)
		if err != nil {
			return err
		}
	}

	сon.LogInfo("started consuming messages", logFields)
	msgs, err := ch.Consume(
		сon.QueueName, // queue
		"",            // consumer tag
		false,         // auto-ack
		false,         // exclusive
		false,         // no-local
		false,         // no-wait
		nil,           // args
	)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	sem := make(chan struct{}, numOfWorkers)
	grp, ctx := errgroup.WithContext(ctx)
L:
	for {
		select {
		case <-ctx.Done():
			defer сon.Close()
			сon.LogInfo("context cancelled", logFields)
			break L
		case d, ok := <-msgs:
			delivery := d
			if !ok {
				сon.LogInfo("channel closed", logFields)
				break L
			}
			sem <- struct{}{}
			grp.Go(func() error {
				defer func() { <-sem }()
				return сon.processMessage(ctx, &delivery)
			})
		}
	}

	if err := grp.Wait(); err != nil {
		return fmt.Errorf("consumer error %w", err)
	}

	return nil
}

func (c *Consumer) LogInfo(msg string, fields log.Fields) {
	c.logger.WithFields(fields).Infof("[consumer] %s", msg)
}

func (c *Consumer) LogError(msg string, fields log.Fields) {
	c.logger.WithFields(fields).Errorf("[consumer] %s", msg)
}

func (c *Consumer) Close() error {
	if c.channel != nil {
		if err := c.channel.Close(); err != nil {
			return err
		}
	}

	if c.conn != nil {
		if err := c.conn.Close(); err != nil {
			return err
		}
	}

	return nil
}

func (con *Consumer) processMessage(ctx context.Context, d *amqp.Delivery) error {
	logFields := log.Fields{"queue": con.QueueName}
	logFields["body"] = string(d.Body)
	logFields["gorouitnes"] = runtime.NumGoroutine()
	con.LogInfo("received message", logFields)

	handlerError := con.handler.Handle(ctx, d.Body)
	if handlerError == nil {
		return con.ackMessage(d, logFields)
	}

	if handlerError.ErrorType == Skippable {
		con.LogError("error while processing message. skipping it", logFields)
		return nil
	} else if handlerError.ErrorType == Retriable {
		con.LogError("error while processing message. will retry", logFields)

		for i := 0; i < maxRetries; i++ {
			con.LogInfo(fmt.Sprintf("retry attempt=%d", i+1), logFields)
			err := con.handler.Handle(ctx, d.Body)
			if err == nil {
				return con.ackMessage(d, logFields)
			}
			con.sleeper.Sleep(i)
		}
		con.LogError("failed to retry message. skipping it", logFields)

		return con.ackMessage(d, logFields)
	}

	return nil
}

func (con *Consumer) ackMessage(d *amqp.Delivery, logFields log.Fields) error {
	if err := d.Ack(false); err != nil {
		return fmt.Errorf("error acknowledging message : %w", err)
	}

	con.LogInfo("acknowledged message", logFields)
	return nil
}

func (con *Consumer) declareQueueWithDirectExchange(ch *amqp.Channel) error {
	err := ch.ExchangeDeclare(
		con.QueueName,             // Exchange name
		common.DirectExchangeType, // Exchange type
		true,                      // Durable
		false,                     // Auto-deleted
		false,                     // Internal
		false,                     // No-wait
		nil,                       // Arguments
	)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	q, err := ch.QueueDeclare(
		con.QueueName, // Queue name
		true,          // Durable
		false,         // Delete when unused
		false,         // Exclusive
		false,         // No-wait
		nil,           // Arguments
	)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	err = ch.QueueBind(
		q.Name, // Queue name
		"",     // Routing key (empty for direct exchanges)
		q.Name, // Exchange name
		false,  // No-wait
		nil,    // Arguments
	)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (con *Consumer) declareQueueWithDelayedExchange(ch *amqp.Channel) error {
	args := make(amqp.Table)
	args["x-delayed-type"] = "direct"
	err := ch.ExchangeDeclare(
		con.QueueName,              // Exchange name
		common.DelayedExchangeType, // Exchange type
		true,                       // Durable
		false,                      // Auto-deleted
		false,                      // Internal
		false,                      // No-wait
		args,                       // Arguments
	)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	q, err := ch.QueueDeclare(
		con.QueueName, // Queue name
		true,          // Durable
		false,         // Delete when unused
		false,         // Exclusive
		false,         // No-wait
		nil,           // Arguments
	)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	err = ch.QueueBind(
		q.Name, // Queue name
		"",     // Routing key (empty for direct exchanges)
		q.Name, // Exchange name
		false,  // No-wait
		nil,    // Arguments
	)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}
