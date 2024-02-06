package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Publisher struct {
	connection *amqp.Connection
}

func NewPublisher(conn *amqp.Connection) *Publisher {
	return &Publisher{
		connection: conn,
	}
}

func (e *Publisher) Push(queueName string, event string) error {
	channel, err := e.connection.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()

	queue, err := channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(queue.Name)

	err = channel.Publish(
		"",        // exchange
		queueName, // queue name
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(event),
		},
	)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Sending message: %s -> %s", event, queueName)
	return nil
}
