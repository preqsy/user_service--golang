package datastore

import "user_service/models"

type Datastore interface {
	SaveUser(user models.User) (*models.User, error)
}
