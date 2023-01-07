package db

import (
	"context"
	"fmt"

	"github.com/graphql_example/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDb(cfg *config.Config) (*mongo.Client, error) {
	return mongo.Connect(context.Background(), options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", cfg.MongoHost, cfg.MongoPort)))
}
