package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"go_notifier/internal/common"
	"go_notifier/pkg/transport/rabbitmq"
	"time"
)

type FirstQueueMessageHandler struct {
}

func (mh *FirstQueueMessageHandler) Handle(ctx context.Context, b []byte) *rabbitmq.HandlerError {
	fmt.Println("first handler")
	var event common.OneEvent
	err := json.Unmarshal(b, &event)
	if err != nil {
		return rabbitmq.NewSkippableError(err, "error while processing message")
	}

	if event.TicketID == 2 {
		// time.Sleep(time.Microsecond * 30)
		return rabbitmq.NewRetriableError(err, "error while processing message")
	}

	time.Sleep(time.Second * 1)

	return nil
}

type SecondQueueMessageHandler struct {
}

func (mh *SecondQueueMessageHandler) Handle(ctx context.Context, b []byte) *rabbitmq.HandlerError {
	fmt.Println("second handler")
	var event common.TwoEvent
	err := json.Unmarshal(b, &event)
	if err != nil {
		return rabbitmq.NewSkippableError(err, "error while processing message")
	}

	return nil
}
