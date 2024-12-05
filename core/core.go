package core

import (
	"context"
	"time"
	datastore "user_service/database"
	"user_service/events"
	"user_service/events/topics"
	"user_service/models"

	"encoding/json"

	"github.com/sirupsen/logrus"
)

type Service struct {
	datastore datastore.Datastore
	rabbitmq  *events.RabbitmqService
}

func CoreService(datastore datastore.Datastore, rabbitmq *events.RabbitmqService) *Service {
	return &Service{datastore: datastore, rabbitmq: rabbitmq}
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	data, err := json.Marshal(&user)
	if err != nil {
		logrus.Error("Error while converting data to json")
		return nil, err
	}

	s.rabbitmq.PublishNewUserCreated(ctx, data, topics.NewUserCreated)

	return user, nil
}
