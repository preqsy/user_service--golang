package main

import (
	"fmt"
	"user_service/config"
	database "user_service/database/postres"

	// "user_service/datastore"
	"user_service/models"
)

func main() {
	secrets := config.GetSecrets()

	_, err := database.ConnectDB(secrets.Host, secrets.Db_User, secrets.Password, secrets.DbName, secrets.Port).SaveUser(models.User{Email: "obbypreciouss2yk44@gmail.com", Name: "Obinna", Password: "1111"})
	fmt.Println(err)

}
