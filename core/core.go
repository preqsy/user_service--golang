package core

import (
	datastore "user_service/database"
	"user_service/models"

	"github.com/sirupsen/logrus"
)

type Service struct {
	datastore datastore.Datastore
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
	return s.datastore.SaveUser(userData)
}
