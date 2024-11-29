package main

import (
	"log"
	"user_service/config"
	"user_service/core"
	datastore "user_service/database"
	database "user_service/database/postres"

	"user_service/models"
)

func main() {
	secrets := config.GetSecrets()
	var datastore datastore.Datastore

	datastore, err := database.ConnectDB(secrets.Host, secrets.Db_User, secrets.Password, secrets.DbName, secrets.Port)
	if err != nil {
		log.Fatal("Error connecting to the database")
	}

	service := core.CoreService(datastore)
	service.SaveUser(models.User{Email: "preciousohanyere08@gmail.com", Name: "Obinna", Password: "1111"})

}
