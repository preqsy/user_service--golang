package main

import (
	"user_service/config"
	database "user_service/database/postres"
)

func main() {
	secrets := config.GetSecrets()

	_ = database.ConnectDB(secrets.Host, secrets.User, secrets.DbName, secrets.Password, secrets.Port)

}
