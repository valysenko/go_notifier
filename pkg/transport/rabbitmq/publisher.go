package rabbitmq

import (
	"fmt"
	"go_notifier/internal/common"

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

	err = p.queueExists(channel, queueName)
	if err != nil {
		return err
	}

	err = channel.Publish(
		queueName, // exchange
		"",        // routing key
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

func (p *Publisher) PublishWithDelay(queueName string, event []byte, delayMs int) error {
	channel, err := p.connection.Channel()
	if err != nil {
		return err
	}

	exchangeType := common.QueueExchangeType[queueName]
	if exchangeType != common.DelayedExchangeType {
		return fmt.Errorf("not allowed to publish with delay to the %s exchange", queueName)
	}

	defer channel.Close()

	err = p.queueExists(channel, queueName)
	if err != nil {
		return err
	}

	err = channel.Publish(
		queueName, // exchange
		"",        // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        event,
			Headers: amqp.Table{
				"x-delay": delayMs,
			},
		},
	)
	if err != nil {
		return err
	}

	p.logger.Infof("[producer] queue=[%s] sent message=%s with delay=%d", queueName, string(event), delayMs)

	return nil
}

func (p *Publisher) queueExists(ch *amqp.Channel, queueName string) error {
	_, err := ch.QueueDeclarePassive(queueName, true, false, false, false, nil)
	return err
}
