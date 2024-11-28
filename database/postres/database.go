package database

import (
	"fmt"
	"log"

	"user_service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConfig struct {
	host     string
	port     string
	password string
	user     string
	dbName   string
}
type PostgresStore struct {
	client *gorm.DB
}

func (p PostgresStore) SaveUser(user_data models.User) error {
	p.client.Create(user_data)
	return nil
}

func ConnectDB(host, user, password, dbName, port string) *PostgresStore {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbName=%s port=%v sslmode=disable TimeZone=UTC", host, user, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failure to connect to the database", err)
	}

	db.AutoMigrate()
	return &PostgresStore{client: db}
}
