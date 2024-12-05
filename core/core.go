package core

import (
	"context"
	"time"
	datastore "user_service/database"
	"user_service/events/topics"
	"user_service/models"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

type Service struct {
	datastore datastore.Datastore
}

func failOnError(err error, msg string) {
	if err != nil {
		logrus.Panicf("%s: %s", msg, err)
	}
}

func CoreService(datastore datastore.Datastore) *Service {
	return &Service{datastore: datastore}
}

func (s Service) SaveUser(userData models.User) (*models.User, error) {
	err := userData.Validate()
	if err != nil {
		logrus.Debug("Input correct details")
		return nil, err
	}
	user, err := s.datastore.SaveUser(userData)

	if err != nil {
		return nil, err
	}
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5673")
	failOnError(err, "Failed to connect to Rabbitmq")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	q, err := ch.QueueDeclare(
		topics.NewUserCreated,
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to open a Queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ch.PublishWithContext(
		ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			Body:        []byte("body"),
			ContentType: "text/plain",
		},
	)
	logrus.Info("[*] Queue message sent successfully")
	return user, nil
}
