package service

import (
	"github.com/streadway/amqp"
)

type ChannelMessageBrokerType struct {
	AMPQ *amqp.Channel
}

var ChannelMessageBroker = &ChannelMessageBrokerType{}

func NewProducer(config *ConfigType) (*amqp.Connection, error) {

	connectRabbitMQ, err := newConnectMessageBroker(config)
	if err != nil {
		return nil, err
	}

	if err = queueDeclare(ChannelMessageBroker.AMPQ); err != nil {
		return nil, err
	}

	return connectRabbitMQ, nil
}

func newConnectMessageBroker(config *ConfigType) (*amqp.Connection, error) {
	connectRabbitMQ, err := amqp.Dial(config.AMQPURL)
	if err != nil {
		return nil, err
	}

	ChannelMessageBroker.AMPQ, err = connectRabbitMQ.Channel()
	if err != nil {
		return nil, err
	}
	return connectRabbitMQ, nil
}

func queueDeclare(channel *amqp.Channel) error {

	_, err := channel.QueueDeclare(
		"SendRequest",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	return nil
}

func NewConsumerChanel(config *ConfigType) (*amqp.Connection, error) {
	return newConnectMessageBroker(config)
}
