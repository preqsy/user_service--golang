package main

import (
	"log"
	"user_service/config"
	"user_service/core"
	datastore "user_service/database"
	database "user_service/database/postres"

	"user_service/models"

	"net/http"
	"user_service/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	secrets := config.GetSecrets()
	var datastore datastore.Datastore
	const defaultPort = "8080"

	datastore, err := database.ConnectDB(secrets.Host, secrets.Db_User, secrets.Password, secrets.DbName, secrets.Port)
	if err != nil {
		log.Fatal("Error connecting to the database")
	}

	service := core.CoreService(datastore)
	service.SaveUser(models.User{Email: "preciousohanyere088@gmail.com", Name: "Obinna", Password: "1111"})

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, nil))

}
