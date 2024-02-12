package main

import (
	"encoding/json"
	"go_notifier/configs"
	"go_notifier/internal/common"
	"go_notifier/pkg/transport/rabbitmq"

	log "github.com/sirupsen/logrus"
)

func main() {
	cfg := configs.InitConfig()
	rabbitApp := rabbitmq.NewRabbitApp(&cfg.RabbitConfig)
	defer rabbitApp.Close()
	publisher := rabbitmq.NewPublisher(rabbitApp.Connection, log.New())

	event := rabbitmq.OneEvent{TicketID: 1, CommentID: 2}
	body, err := json.Marshal(event)
	if err != nil {
	}
	publisher.Publish(common.RabbitFirstQueue, body)

	// event2 := rabbitmq.TwoEvent{AuthorId: 12}
	// body, err = json.Marshal(event2)
	// if err != nil {
	// }
	// publisher.Publish(common.RabbitSecondQueue, body)
}
