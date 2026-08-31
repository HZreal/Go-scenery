package rabbitmqdemo

import amqp "github.com/rabbitmq/amqp091-go"

func Dial(url string) (*amqp.Connection, error) {
	return amqp.Dial(url)
}

func DeclareDemoQueue(ch *amqp.Channel, name string) (amqp.Queue, error) {
	return ch.QueueDeclare(
		name,
		true,
		false,
		false,
		false,
		amqp.Table{"x-queue-type": "quorum"},
	)
}
