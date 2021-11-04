package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	var err error

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	CheckError(err)

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	CheckError(err)

	// get collection as ref
	collection := client.Database("go-web-dev-db").Collection("persons")

	// insert
	john := Person{"John", 24}
	jane := Person{"Jane", 27}
	ben := Person{"Ben", 16}

	_, err = collection.InsertOne(context.TODO(), john)
	CheckError(err)

	persons := []interface{}{jane, ben}
	_, err = collection.InsertMany(context.TODO(), persons)
	CheckError(err)

	// update
	filter := bson.D{primitive.E{Key: "name", Value: "John"}}

	update := bson.D{primitive.E{
		Key: "$set", Value: bson.D{
			primitive.E{Key: "age", Value: 26},
		}},
	}

	_, err = collection.UpdateOne(context.TODO(), filter, update)
	CheckError(err)

	// find
	var res Person
	f2 := bson.D{primitive.E{Key: "name", Value: "Ben"}}
	_ = collection.FindOne(context.TODO(), f2).Decode(&res)
	fmt.Println(res)

	// delete
	_, err = collection.DeleteMany(context.TODO(), bson.D{{}})
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
