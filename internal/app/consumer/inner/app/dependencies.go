package app

import (
	"go_notifier/internal/app/consumer/inner/handler"
	"go_notifier/internal/common"
	"go_notifier/internal/notifier"
	"go_notifier/pkg/firebase"
	"go_notifier/pkg/transport/rabbitmq"
	"go_notifier/pkg/twilio"
	"os"

	log "github.com/sirupsen/logrus"
)

func NewLogger() log.FieldLogger {
	logger := log.New()
	logger.SetOutput(os.Stdout)
	//logger.SetFormatter(&log.JSONFormatter{DataKey: "fields"})
	// logger.SetFormatter(&logrus.TextFormatter{
	// 	ForceColors:   true,
	// 	FullTimestamp: true,
	// })

	return logger
}

func NewTwilio(app *ConsumerApp) *twilio.Twilio {
	return twilio.NewTwilio(
		app.cfg.TwilioConfig.AccountSid,
		app.cfg.TwilioConfig.AuthToken,
		app.cfg.TwilioConfig.AccountNumber,
		app.logger,
	)
}

func NewFirebase(app *ConsumerApp) *firebase.Firebase {
	key, err := app.cfg.FirebaseConfig.GetDecodedFireBaseKey()
	if err != nil {
		panic(err)
	}
	return firebase.NewFirebase(key, app.logger)
}

func NewNotifierProvider(app *ConsumerApp) *notifier.NotifierProvider {
	return notifier.NewNotifierProvider(
		notifier.NewFirebaseNotifier(NewFirebase(app)),
		notifier.NewTwilioNotifier(NewTwilio(app)),
	)
}

func NewConsumersMessageHandlers(app *ConsumerApp) rabbitmq.MessageHandlersMap {
	mp := make(rabbitmq.MessageHandlersMap)
	mp[common.RabbitFirstQueue] = []rabbitmq.MessageHandler{
		&handler.FirstQueueMessageHandler{},
	}
	mp[common.RabbitSecondQueue] = []rabbitmq.MessageHandler{
		&handler.SecondQueueMessageHandler{},
	}
	mp[common.ScheduledNotificationsQueue] = []rabbitmq.MessageHandler{
		handler.NewScheduledNotificationHandler(app.logger, NewNotifierProvider(app)),
	}

	return mp
}
