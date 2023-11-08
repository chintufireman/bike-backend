package db

import (
	"bike-website/model"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertNewBikeDetails(bike *model.Bike) (string, error) {
	client, err := InitDatabase()

	if err != nil {
		return "INTERNAL ISSUE OCCURED DURING DATA CONNECTION", err
	}
	defer client.Disconnect(context.Background())

	//creating context for db operation
	ctx := context.Background()

	// Specify the collection
	collection := client.Database("Bike").Collection("bike-details")

	var id *mongo.InsertOneResult
	// Insert the document into the collection
	id, err = collection.InsertOne(ctx, bike)
	if err != nil {
		return "", err
	}

	return "User inserted successfully with Id " + id.InsertedID.(primitive.ObjectID).Hex(), nil
}
