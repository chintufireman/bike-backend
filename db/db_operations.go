package db

import (
	"bike-website/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
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

func FetchAllBikes() ([]*model.Bike, error) {
	client, error := InitDatabase()

	if error != nil {
		return nil, error
	}

	ctx := context.Background()
	collection := client.Database("Bike").Collection("bike-details")
	filter := bson.M{}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var bikes []*model.Bike
	for cursor.Next(ctx) {
		var bike model.Bike
		if err := cursor.Decode(&bike); err != nil {
			return nil, err
		}
		bikes = append(bikes, &bike)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return bikes, nil
}

func UploadImageData(image *model.Image) (string, error) {
	client, err := InitDatabase()

	if err != nil {
		return "", err
	}
	defer client.Disconnect(context.Background())
	ctx := context.Background()
	collection := client.Database("Bike").Collection("bike-details")

	var id *mongo.InsertOneResult
	id, err = collection.InsertOne(ctx, image)

	if err != nil {
		return "", err
	}
	
	return "Image inserted successfully with Id " + id.InsertedID.(primitive.ObjectID).Hex(), nil

}

func FetchImage() ([]byte, error) {
	client, err := InitDatabase()

	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())
	ctx := context.Background()
	collection := client.Database("Bike").Collection("bike-details")

	filter := bson.M{}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var image model.Image
    if cursor.Next(ctx) {
        if err := cursor.Decode(&image); err != nil {
            return nil, err
        }
        return image.Data, nil
    }

	return image.Data,nil
}
