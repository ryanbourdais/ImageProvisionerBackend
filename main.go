package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Data struct {
	Number int
	String string
}

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+1.6.0"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("testDatabase").Collection("test")

	testData1 := Data{1, "A"}

	_, err = collection.InsertOne(context.TODO(), testData1)
	if err != nil {
		log.Fatal(err)
	}

	result, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	defer result.Close(ctx)
	for result.Next(ctx) {
		var results bson.M
		if err = result.Decode(&results); err != nil {
			log.Fatal(err)
		}
		fmt.Println(results)
	}

	fmt.Println(databases)

	fmt.Println("Hello, world")
}
