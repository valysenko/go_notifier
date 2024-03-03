package handler

import (
	"context"
	"encoding/json"
	"go_notifier/internal/common"
	"go_notifier/internal/notifier"
	"go_notifier/pkg/transport/rabbitmq"

	log "github.com/sirupsen/logrus"
)

// go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name=NotifierProvider --case snake
type NotifierProvider interface {
	Provide(appType string) notifier.Notifier
}

type ScheduledNotificationHandler struct {
	logger           log.FieldLogger
	notifierProvider NotifierProvider
}

func NewScheduledNotificationHandler(logger log.FieldLogger, notifierProvider NotifierProvider) *ScheduledNotificationHandler {
	return &ScheduledNotificationHandler{
		logger:           logger,
		notifierProvider: notifierProvider,
	}
}

func (snh *ScheduledNotificationHandler) Handle(ctx context.Context, b []byte) *rabbitmq.HandlerError {
	var event common.ScheduledNotification
	err := json.Unmarshal(b, &event)
	if err != nil {
		return rabbitmq.NewSkippableError(err, "error while processing message")
	}

	notifier := snh.notifierProvider.Provide(event.AppType)
	if notifier == nil {
		return nil
	}

	err = notifier.Notify(ctx, event.AppIdentifier, event.Message)
	if err != nil {
		return rabbitmq.NewRetriableError(err, "error while processing message")
	}

	snh.logger.WithFields(log.Fields{
		"appIdentifier": event.AppIdentifier,
		"appType":       event.AppType,
		"message":       event.Message,
	}).Infof("sent notification")

	return nil
}
