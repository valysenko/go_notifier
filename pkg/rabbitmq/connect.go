package rabbitmq

import "github.com/streadway/amqp"

const RabbitFirstQueue = "first"
const RabbitSecondQueue = "second"

type AppRabbit struct {
	connection      *amqp.Connection
	supportedQueues []string
}

func NewAppRabbit() (*AppRabbit, error) {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		return nil, err
	}

	return &AppRabbit{
		connection:      connection,
		supportedQueues: getQueueNames(),
	}, nil
}

func NewConnection() *amqp.Connection {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		panic(err)
	}

	return connection
}

func getQueueNames() []string {
	return []string{RabbitFirstQueue, RabbitSecondQueue}
}
