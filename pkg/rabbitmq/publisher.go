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

func (e *Publisher) Push(queueName string, event []byte) error {
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
			ContentType: "application/json",
			Body:        event,
		},
	)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Sending message: %s -> %s", event, queueName)

	return nil
}

// example events
type OneEvent struct {
	TicketID  int `json:"ticketId"`
	CommentID int `json:"commentId"`
}

type TwoEvent struct {
	AuthorId int `json:"authorId"`
}
