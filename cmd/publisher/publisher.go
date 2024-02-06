package main

import (
	"encoding/json"
	"go_notifier/pkg/rabbitmq"
)

func main() {
	amqpConnection := rabbitmq.NewConnection()
	defer amqpConnection.Close()
	publisher := rabbitmq.NewPublisher(amqpConnection)

	event := rabbitmq.OneEvent{TicketID: 1, CommentID: 2}
	body, err := json.Marshal(event)
	if err != nil {
	}
	publisher.Push(rabbitmq.RabbitFirstQueue, body)

	event2 := rabbitmq.TwoEvent{AuthorId: 12}
	body, err = json.Marshal(event2)
	if err != nil {
	}
	publisher.Push(rabbitmq.RabbitSecondQueue, body)
}
