package database

import (
	"Mereb3/constants"
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBInstance creates a new MongoDB client and connects to the database.
func DBInstance() (*mongo.Client, error) {
	MongoURI := os.Getenv("MONGO_URI")
	if MongoURI == "" {
		MongoURI = "mongodb://localhost:27017"
	}

	ctx, cancel := context.WithTimeout(context.Background(), constants.TIME_OUT)
	defer cancel()

	clientOptions := options.Client().ApplyURI(MongoURI)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", constants.DATABASE_CONNECTED_FAILED, err)
	}

	return client, nil
}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("mereb-task-v2").Collection(collectionName)
}

func CreateMongoClient() (*mongo.Client, error) {
	return DBInstance()
}
