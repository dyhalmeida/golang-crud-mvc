package mongodb

import (
	"context"
	"os"

	"github.com/dyhalmeida/golang-crud-mvc/src/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGO_URI = "MONGO_URI"
	MONGO_DATABASE = "MONGO_DATABASE"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {

	mongodbUri := os.Getenv(MONGO_URI)
	mongodbDatabase := os.Getenv(MONGO_DATABASE)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbUri))

	if utils.HasError(err) {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if utils.HasError(err) {
		return nil, err
	}
	

	return client.Database(mongodbDatabase), nil

}