package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

type ConsumersMap map[string]*Consumer

func NewConsumersMap(conn *amqp.Connection) ConsumersMap {
	hm := NewMessageHandlers()
	consumers := make(ConsumersMap)
	for queue, handlers := range hm {
		for _, h := range handlers {
			consumers[queue] = NewConsumer(conn, queue, h)
		}
	}

	return consumers
}

func (cm ConsumersMap) Close() error {
	for _, consumer := range cm {
		if err := consumer.Close(); err != nil {
			return err
		}
	}
	return nil
}

type Consumer struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	handler   MessageHandler
	QueueName string
}

func NewConsumer(conn *amqp.Connection, queueName string, handler MessageHandler) *Consumer {
	consumer := &Consumer{
		conn:      conn,
		QueueName: queueName,
		handler:   handler,
	}

	return consumer
}

func (с *Consumer) Listen(ctx context.Context) error {
	ch, err := с.conn.Channel()
	if err != nil {
		return err
	}
	с.channel = ch
	defer ch.Close()

	q, err := ch.QueueDeclare(
		с.QueueName, // name
		true,        // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	done := ctx.Done()

	go func() {
		for {
			select {
			case <-done:
				с.Close()
				fmt.Println("Context cancelled, exiting Listen loop")
				return
			case d, ok := <-msgs:
				if !ok {
					fmt.Println("Channel closed, exiting Listen loop")
					return
				}
				fmt.Println(fmt.Sprintf("Received message: %s", d.Body))
				с.handler.Handle(ctx, d.Body)
				fmt.Println(fmt.Sprintf("Processed message: %s", d.Body))
			}
		}
	}()

	<-done
	fmt.Printf("[*] Waiting for message [Queue][%s]. To exit press CTRL+C", q.Name)
	return nil
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

// handlers
type MessageHandlersMap map[string][]MessageHandler

type MessageHandler interface {
	Handle(ctx context.Context, b []byte) error
}

func NewMessageHandlers() MessageHandlersMap {
	mp := make(MessageHandlersMap)
	mp[RabbitFirstQueue] = []MessageHandler{&FirstQueueMessageHandler{}}
	mp[RabbitSecondQueue] = []MessageHandler{&SecondQueueMessageHandler{}}

	return mp
}

type FirstQueueMessageHandler struct {
}

func (mh *FirstQueueMessageHandler) Handle(ctx context.Context, b []byte) error {
	fmt.Println("first handler")
	var event OneEvent
	err := json.Unmarshal(b, &event)
	if err != nil {
		return err
	}

	fmt.Println(event)
	return nil
}

type SecondQueueMessageHandler struct {
}

func (mh *SecondQueueMessageHandler) Handle(ctx context.Context, b []byte) error {
	fmt.Println("second handler")
	var event TwoEvent
	err := json.Unmarshal(b, &event)
	if err != nil {
		return err
	}

	fmt.Println(event)
	return nil
}
