package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/graphql_example/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepo struct {
	Db *mongo.Client
}

func NewProductRepo(db *mongo.Client) *ProductRepo {
	return &ProductRepo{Db: db}
}

func (db *ProductRepo) CreateProduct(productReq *model.NewProduct) (*model.Product, error) {
	collection := db.Db.Database("productdb").Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()
	product := model.Product{}
	product.ID = uuid.New().String()
	_, err := collection.InsertOne(ctx, bson.D{
		{Key: "id", Value: product.ID},
		{Key: "name", Value: product.Name},
		{Key: "model", Value: product.Model},
		{Key: "price", Value: product.Price},
		{Key: "amount", Value: product.Amount},
	})
	fmt.Println("error while create product", err)
	if err != nil {
		return &model.Product{}, err
	}

	return &product, nil
}

func (r *ProductRepo) GetProduct(name *model.GetProductByName) (*model.Product, error) {
	collection := r.Db.Database("productdb").Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	product := &model.Product{}
	filter := bson.D{{"name", &name.Name}}
	err := collection.FindOne(ctx, filter).Decode(product)
	if err != nil {
		return &model.Product{}, err
	}
	return product, nil
}
