package graph

import (
	"fmt"

	"github.com/graphql_example/internal/controller/service"
	"github.com/graphql_example/pkg/config"
	"github.com/graphql_example/pkg/db"
	"github.com/graphql_example/pkg/logger"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Service *service.ProductService
}

func Connect(cfg *config.Config) *service.ProductService {
	l := logger.New(cfg.LogLevel)
	con, err := db.ConnectToDb(cfg)
	if err != nil {
		var someErr error
		fmt.Errorf("error connecting to mongo in resolver", someErr)
		l.Fatal(someErr, err.Error())
	}
	productService := service.NewProductService(l, con)
	return productService
}
