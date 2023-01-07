package repo

import "github.com/graphql_example/graph/model"

type ProductStorageI interface {
	CreateProduct(*model.NewProduct) (*model.Product, error)
	GetProduct(*model.GetProductByName) (*model.Product, error)
}
