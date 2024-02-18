package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"go_notifier/internal/common"

	log "github.com/sirupsen/logrus"
)

type FirstQueueMessageHandler struct {
}

func (mh *FirstQueueMessageHandler) Handle(ctx context.Context, b []byte) error {
	fmt.Println("first handler")
	var event common.OneEvent
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
	var event common.TwoEvent
	err := json.Unmarshal(b, &event)
	if err != nil {
		return err
	}

	fmt.Println(event)
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

func (snh *ScheduledNotificationHandler) Handle(ctx context.Context, b []byte) error {
	var event common.ScheduledNotification
	err := json.Unmarshal(b, &event)
	if err != nil {
		return err
	}

	snh.logger.WithFields(log.Fields{
		"event": event,
	}).Info("ScheduledNotificationHandler - received event")

	// TODO ...

	snh.logger.WithFields(log.Fields{
		"event": event,
	}).Info("ScheduledNotificationHandler - processed event")

	return nil
}
