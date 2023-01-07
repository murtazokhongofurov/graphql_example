package storage

import (
	mongodb "github.com/graphql_example/internal/controller/storage/mongo"
	"github.com/graphql_example/internal/controller/storage/repo"
	"go.mongodb.org/mongo-driver/mongo"
)

type IStorage interface {
	Product() repo.ProductStorageI
}

type StoragePg struct {
	Db          *mongo.Client
	ProductRepo repo.ProductStorageI
}

func NewStoragePg(db *mongo.Client) *StoragePg {
	return &StoragePg{
		Db:          db,
		ProductRepo: mongodb.NewProductRepo(db),
	}
}

func (s StoragePg) Product() repo.ProductStorageI {
	return s.ProductRepo
}
