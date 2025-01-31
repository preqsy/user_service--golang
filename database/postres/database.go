package database

import (
	"fmt"
	"log"
	"time"

	"user_service/models"
	"user_service/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresStore struct {
	client *gorm.DB
}

func (p PostgresStore) SaveUser(userData models.User) (*models.User, error) {

	userData.TimeCreated = time.Now()
	userData.Password = utils.EncryptPassword(userData.Password)
	result := p.client.Create(&userData)

	if result.Error != nil {
		return nil, fmt.Errorf("Error occured while saving: %v", result.Error)
	}

	return &userData, nil
}
func ConnectDB(host, user, password, dbName, port string) (*PostgresStore, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", host, user, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failure to connect to the database", err)
	}
	logrus.Info("Connected to the Database")

	db.AutoMigrate(&models.User{})
	return &PostgresStore{client: db}, err
}
