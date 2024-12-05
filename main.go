package main

import (
	"log"
	"user_service/config"
	"user_service/core"
	database "user_service/database/postres"
	"user_service/events"

	"net/http"
	"user_service/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	secrets := config.GetSecrets()

	const defaultPort = "8080"

	datastore, _ := database.ConnectDB(secrets.Host, secrets.Db_User, secrets.Password, secrets.DbName, secrets.Port)
	rabbitService := events.NewRabbitMqService("amqp://guest:guest@localhost:5673")
	service := core.CoreService(datastore, rabbitService)
	resolver := graph.NewResolver(service)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, nil))

}
