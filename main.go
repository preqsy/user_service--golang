package main

import (
	"fmt"
	"user_service/config"
	database "user_service/database/postres"
	"user_service/models"
)

func main() {
	secrets := config.GetSecrets()

	_ = database.ConnectDB(secrets.Host, secrets.Db_User, secrets.Password, secrets.DbName, secrets.Port).SaveUser(models.User{Email: "obbyprecious24@gmail.com", Name: "Obinna", Password: "1111"})

}
