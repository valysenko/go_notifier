package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"go_notifier/internal/common"
	"go_notifier/pkg/transport/rabbitmq"

	log "github.com/sirupsen/logrus"
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

	fmt.Println(event)
	return rabbitmq.NewRetriableError(err, "error while processing message")
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

	fmt.Println(event)
	return rabbitmq.NewSkippableError(err, "error while processing message")
	return nil
}

type ScheduledNotificationHandler struct {
	logger log.FieldLogger
}

func NewScheduledNotificationHandler(logger log.FieldLogger) *ScheduledNotificationHandler {
	return &ScheduledNotificationHandler{
		logger: logger,
	}
}

func (snh *ScheduledNotificationHandler) Handle(ctx context.Context, b []byte) *rabbitmq.HandlerError {
	var event common.ScheduledNotification
	err := json.Unmarshal(b, &event)
	if err != nil {
		return rabbitmq.NewSkippableError(err, "error while processing message")
	}

	snh.logger.WithFields(log.Fields{
		"event": event,
	}).Info("ScheduledNotificationHandler - processing event")

	// TODO ...

	return nil
}
