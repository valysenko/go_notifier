package main

import (
	"encoding/json"
	"fmt"
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

	// event := common.OneEvent{TicketID: 1, CommentID: 2}
	// body, err := json.Marshal(event)
	// if err != nil {
	// }
	// publisher.Publish(common.RabbitFirstQueue, body)

	event2 := common.TwoEvent{AuthorId: 12}
	body, err := json.Marshal(event2)
	if err != nil {
	}
	publisher.Publish(common.RabbitSecondQueue, body)
	err = publisher.PublishWithDelay(common.RabbitSecondQueue, body, 5000)
	fmt.Println(err)

	// for i := 0; i < 1000; i++ {
	// 	event := common.OneEvent{TicketID: 1, CommentID: 2}
	// 	body, err := json.Marshal(event)
	// 	if err != nil {
	// 	}
	// 	publisher.Publish(common.RabbitFirstQueue, body)
	// }

	//////////////

	// event := common.OneEvent{TicketID: 1, CommentID: 2}
	// body, err := json.Marshal(event)
	// if err != nil {
	// }
	// publisher.Publish(common.RabbitFirstQueue, body)

	// event = common.OneEvent{TicketID: 2, CommentID: 2} // will retry
	// body, err = json.Marshal(event)
	// if err != nil {
	// }
	// publisher.Publish(common.RabbitFirstQueue, body)

	// event = common.OneEvent{TicketID: 3, CommentID: 2}
	// body, err = json.Marshal(event)
	// if err != nil {
	// }
	// publisher.Publish(common.RabbitFirstQueue, body)

	// event = common.OneEvent{TicketID: 4, CommentID: 2}
	// body, err = json.Marshal(event)
	// if err != nil {
	// }
	// publisher.Publish(common.RabbitFirstQueue, body)

	// event2 := common.TwoEvent{AuthorId: 12}
	// body, err = json.Marshal(event2)
	// if err != nil {
	// }
	// publisher.Publish(common.RabbitSecondQueue, body)

	// time.Sleep(time.Second * 2)

	// event = common.OneEvent{TicketID: 5, CommentID: 2}
	// body, err = json.Marshal(event)
	// if err != nil {
	// }
	// publisher.Publish(common.RabbitFirstQueue, body)

	// time.Sleep(time.Second * 10)

	// event = common.OneEvent{TicketID: 6, CommentID: 2}
	// body, err = json.Marshal(event)
	// if err != nil {
	// }
	// publisher.Publish(common.RabbitFirstQueue, body)
}
