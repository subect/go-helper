package rabbitMQ

import (
	"github.com/streadway/amqp"
)

func InitializeRabbitMQConnection(url string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func PublishMessage(channel *amqp.Channel, exchange, routingKey, message string) error {
	err := channel.Publish(
		exchange,   // Exchange
		routingKey, // Routing key
		false,      // Mandatory
		false,      // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		return err
	}
	return nil
}
