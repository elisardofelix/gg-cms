package DB

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var client *mongo.Client

func GetMongoDBClient() (*mongo.Client, error) {
	var err error
	// Set client options
	clientOptions := options.Client().ApplyURI(Conf.connectionString)


	if client == nil {
		// Connect to MongoDB
		client, err = mongo.Connect(context.TODO(), clientOptions)
	}

	return client, err
}

func DropMongoTestDB(client *mongo.Client) error {
	err := client.Database(Conf.TestDataBase).Drop(context.TODO())
	return err
}