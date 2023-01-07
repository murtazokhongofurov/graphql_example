package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/graphql_example/graph"
	"github.com/graphql_example/pkg/config"
	"github.com/graphql_example/pkg/db"
	"github.com/rs/cors"
)

func Run(cfg *config.Config) {
	_, err := db.ConnectToDb(cfg)
	if err != nil {
		fmt.Println("error connect db", err)
		return
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{fmt.Sprintf("http://%s:8080", cfg)},
		AllowCredentials: true,
	})

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Service: graph.Connect(cfg)}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", c.Handler(srv))

	log.Printf("connect to http://%s:%s/ for GraphQL playground", cfg.ProductServiceHost, cfg.ProductServicePort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.ProductServicePort), nil))
}
