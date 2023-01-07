package service

import (
	"github.com/graphql_example/internal/controller/storage"
	"github.com/graphql_example/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductService struct {
	Logger  *logger.Logger
	Storage storage.IStorage
}

func NewProductService(l *logger.Logger, stg *mongo.Client) *ProductService {
	return &ProductService{
		Logger:  l,
		Storage: storage.NewStoragePg(stg),
	}
}
