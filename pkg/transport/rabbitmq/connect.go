package rabbitmq

import (
	"go_notifier/configs"

	"github.com/streadway/amqp"
)

type RabbitApp struct {
	Connection *amqp.Connection
	Config     *configs.RabbitConfig
}

func NewRabbitApp(cfg *configs.RabbitConfig) *RabbitApp {
	connection, err := amqp.Dial(cfg.ProvideDSN())
	if err != nil {
		panic(err)
	}

	return &RabbitApp{
		Connection: connection,
		Config:     cfg,
	}
}

func (r *RabbitApp) Close() {
	r.Connection.Close()
}
