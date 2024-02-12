package app

import (
	"go_notifier/internal/common"
	"go_notifier/internal/consumer/inner/handler"
	"go_notifier/pkg/transport/rabbitmq"
	"os"

	log "github.com/sirupsen/logrus"
)

func NewLogger() log.FieldLogger {
	logger := log.New()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&log.JSONFormatter{DataKey: "fields"})
	// logger.SetFormatter(&logrus.TextFormatter{
	// 	ForceColors:   true,
	// 	FullTimestamp: true,
	// })

	return logger
}

func NewConsumersMessageHandlers(app *ConsumerApp) rabbitmq.MessageHandlersMap {
	mp := make(rabbitmq.MessageHandlersMap)
	mp[common.RabbitFirstQueue] = []rabbitmq.MessageHandler{
		&handler.FirstQueueMessageHandler{},
	}
	mp[common.RabbitSecondQueue] = []rabbitmq.MessageHandler{
		&handler.SecondQueueMessageHandler{},
	}

	return mp
}
