package events

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

type RabbitmqService struct {
	channel *amqp.Channel
	conn    *amqp.Connection
}

func failOnError(err error, msg string) {
	if err != nil {
		logrus.Panicf("%s: %s", msg, err)
	}
}

func NewRabbitMqService(rabbitUrl string) *RabbitmqService {
	conn, err := amqp.Dial(rabbitUrl)
	failOnError(err, "Failed to connect to Rabbitmq")

	ch, err := conn.Channel()
	failOnError(err, "Failed to create a channel")

	return &RabbitmqService{channel: ch, conn: conn}

}

func (r RabbitmqService) PublishNewUserCreated(ctx context.Context, data []byte, queueName string) error {

	_, err := r.channel.QueueDeclare(
		queueName,
		true,  // durable
		false, // auto-delete
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to create a queue")
	r.channel.PublishWithContext(
		ctx,
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			Body:        data,
			ContentType: "application/json",
		},
	)
	logrus.Infof("[*] Published message: %s", queueName)
	return nil
}

func (r *RabbitmqService) Close() {
	if r.channel != nil {
		r.channel.Close()
	}
	if r.conn != nil {
		r.conn.Close()
	}
}
