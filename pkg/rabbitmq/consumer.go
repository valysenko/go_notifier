package rabbitmq

import (
	"context"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type ConsumersMap map[string]*Consumer

func NewConsumersMap(conn *amqp.Connection) ConsumersMap {
	queueNames := getQueueNames()
	consumers := make(ConsumersMap)
	for _, qName := range queueNames {
		consumers[qName] = NewConsumer(conn, qName)
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
	QueueName string
}

func NewConsumer(conn *amqp.Connection, queueName string) *Consumer {
	consumer := &Consumer{
		conn:      conn,
		QueueName: queueName,
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
				log.Printf("Received a message: %s", d.Body)
			}
		}
	}()

	<-done
	log.Printf("[*] Waiting for message [Queue][%s]. To exit press CTRL+C", q.Name)
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
