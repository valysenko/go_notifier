package rabbitmq

import (
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type Publisher struct {
	connection *amqp.Connection
	logger     log.FieldLogger
}

func NewPublisher(conn *amqp.Connection, logger log.FieldLogger) *Publisher {
	return &Publisher{
		connection: conn,
		logger:     logger,
	}
}

func (p *Publisher) Publish(queueName string, event []byte) error {
	channel, err := p.connection.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()

	_, err = channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}

	err = channel.Publish(
		"",        // exchange
		queueName, // queue name
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        event,
		},
	)
	if err != nil {
		return err
	}

	p.logger.Infof("[producer] queue=[%s] sent message=%s", queueName, string(event))

	return nil
}
