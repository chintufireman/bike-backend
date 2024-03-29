package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDatabase() (*mongo.Client, error) {
	// Define the MongoDB connection string
	// connectionURI := config.Appconfig.DatabaseURL
	connectionURI := "mongodb://mongodb-bike-container:27017/"
	fmt.Println(connectionURI)
	// Set up the client options
	clientOptions := options.Client().ApplyURI(connectionURI)

	// Connect to the MongoDB server
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB")

	return client, nil
}
