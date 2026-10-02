package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"example.com/go-backend/database"
	"example.com/go-backend/graph"
	"example.com/go-backend/routes"
	"github.com/99designs/gqlgen/graphql/handler"
)

func main() {

	db, err := database.Connect()

	if err != nil {
		fmt.Println("Database connection failed with error: ", err)
		return
	}

	defer db.Close(context.Background())

	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &graph.Resolver{
					DB: db,
				},
			},
		),
	)

	routes.SetupRoutes()

	http.Handle("/query", srv)

	fmt.Println("Server running on http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
