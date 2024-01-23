package db

import (
	"bike-website/model"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	// fmt.Println(collection.Database().Name())
	
	filter := bson.M{}
	fmt.Println(client.ListDatabaseNames(ctx,filter))
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

func FetchImage() ([]*model.Bike, error) {
	client, err := InitDatabase()

	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())
	ctx := context.Background()
	collection := client.Database("Bike").Collection("bike-details")

	filter := bson.M{}
	projection := bson.M{"image": 1} // Include only the 'image' field

	cursor, err := collection.Find(ctx, filter, options.Find().SetProjection(projection))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var images []*model.Bike

	for cursor.Next(ctx) {
		var image *model.Bike
		image = &model.Bike{}

		if err := cursor.Decode(&image); err != nil {
			return nil, err
		}
		// fmt.Println(reflect.ValueOf(image))
		images = append(images, image)
	}

	if len(images) == 0 {
		return nil, errors.New("No documents found with image field")
	}

	return images, nil
}

func FetchOneBikeDetails(name string) (*model.Bike, error) {
	client, err := InitDatabase()
	if err != nil {
		return nil, err
	}
	
	collection := client.Database("Bike").Collection("bike-details")
	defer client.Disconnect(context.Background())
	ctx := context.Background()
	filter:=bson.D{{"name",name}}
	var bikeDetails model.Bike = model.Bike{};
	var bikeDetailsPtr *model.Bike
	bikeDetailsPtr=&bikeDetails

	err2:=collection.FindOne(ctx,filter).Decode(bikeDetailsPtr)
	if err2==mongo.ErrNoDocuments {
		fmt.Println("No matching documents found")
		return nil,err2
	}

	return bikeDetailsPtr,nil
}
