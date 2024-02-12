package rabbitmq

import (
	"context"
	"errors"
	"go_notifier/internal/common"

	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

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
	Handle(ctx context.Context, b []byte) error
}

type Consumer struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	handler   MessageHandler
	logger    log.FieldLogger
	QueueName string
}

func NewConsumer(conn *amqp.Connection, queueName string, handler MessageHandler, logger log.FieldLogger) *Consumer {
	consumer := &Consumer{
		conn:      conn,
		QueueName: queueName,
		handler:   handler,
		logger:    logger,
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

	q, err := ch.QueueDeclare(
		сon.QueueName, // name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	logFields := log.Fields{"queue": сon.QueueName}

	сon.LogInfo("started consuming messages", logFields)
	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	for {
		select {
		case <-ctx.Done():
			defer сon.Close()
			сon.LogInfo("context cancelled", logFields)
			return nil
		case d, ok := <-msgs:
			if !ok {
				сon.LogInfo("channel closed", logFields)
				return nil
			}

			logFields["body"] = string(d.Body)
			сon.LogInfo("received message", logFields)

			var wrapedErr *common.WrappedError
			err = сon.handler.Handle(ctx, d.Body)
			if err != nil {
				if errors.As(err, &wrapedErr) {
					if wrapedErr.ErrorType == common.ErrNotFound {
						// don't need to retry - skip message
						сon.LogError("failed to process message", logFields)
					}
				} else {
					// TODO: retry
					сon.LogInfo("will be retry", logFields)
				}
			}
			сon.LogInfo("processed message", logFields)
		}
	}
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
