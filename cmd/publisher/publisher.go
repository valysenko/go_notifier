package main

import "go_notifier/pkg/rabbitmq"

func main() {
	amqpConnection := rabbitmq.NewConnection()
	defer amqpConnection.Close()
	publisher := rabbitmq.NewPublisher(amqpConnection)
	// publisher.Push(rabbitmq.RabbitFirstQueue, "open ai 1")
	publisher.Push(rabbitmq.RabbitSecondQueue, "zendesk 1")
}
