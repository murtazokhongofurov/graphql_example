package main

import (
	"github.com/graphql_example/internal/app"
	"github.com/graphql_example/pkg/config"
)

func main() {
	cfg := config.Load()
	app.Run(cfg)
}
