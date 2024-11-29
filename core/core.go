package core

import (
	datastore "user_service/database"
	"user_service/models"
)

type Service struct {
	datastore datastore.Datastore
}

func (s Service) SaveUser(userData models.User) (*models.User, error) {
	return s.datastore.SaveUser(userData)
}
