package app

import (
	"context"
	"fmt"
	"go_notifier/configs"
	"go_notifier/pkg/database"
	"go_notifier/pkg/transport/rabbitmq"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"golang.org/x/sync/errgroup"
)

type ConsumerApp struct {
	cfg              *configs.AppConfig
	logger           log.FieldLogger
	mysql            *database.AppDB
	rabbitConnection *amqp.Connection
}

func NewConsumerApp(ctx context.Context, cfg *configs.AppConfig) *ConsumerApp {
	appDB := database.InitDB(&cfg.DBConfig)
	rabbitApp := rabbitmq.NewRabbitApp(&cfg.RabbitConfig)
	logger := NewLogger()

	return &ConsumerApp{
		cfg:              cfg,
		logger:           logger,
		mysql:            appDB,
		rabbitConnection: rabbitApp.Connection,
	}
}

func (app *ConsumerApp) Run(ctx context.Context) error {
	msgHandlers := NewConsumersMessageHandlers(app)
	consumers := rabbitmq.NewConsumersMap(app.rabbitConnection, app.logger, msgHandlers)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	grp, ctx := errgroup.WithContext(ctx)
	for _, consumer := range consumers {
		consumer := consumer
		grp.Go(func() error {
			return consumer.Listen(ctx)
		})
	}
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		app.logger.Info("SIGTERM was send, closing consumers")
		consumers.Close()
		app.rabbitConnection.Close()
		cancel()
	}()

	numGoroutines := runtime.NumGoroutine()
	fmt.Println("Started app. Number of goroutines after:", numGoroutines)

	if err := grp.Wait(); err != nil {
		return fmt.Errorf("ConsumerApp Error: %s", err)
	}

	return nil
}

func (app *ConsumerApp) LogError(err error) {
	app.logger.Errorf(err.Error())
}

func (app *ConsumerApp) Close(ctx context.Context) {
	app.mysql.Mysql.Close()
	app.rabbitConnection.Close()
}
