package database

import (
	"fmt"
	"log"
	"time"

	"user_service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresStore struct {
	client *gorm.DB
}

func (p PostgresStore) SaveUser(user_data models.User) error {
	user_data.TimeCreated = time.Now()
	p.client.Create(user_data)
	return nil
}
func ConnectDB(host, user, password, dbName, port string) *PostgresStore {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", host, user, password, dbName, port)
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failure to connect to the database", err)
	}

	// db.AutoMigrate()
	return &PostgresStore{client: db}
}
