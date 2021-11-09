package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	// connect to a MongoDB server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cli, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = cli.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	// ping the primary to verify that a MongoDB server has been found and connected
	if err = cli.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	} else {
		log.Print("Successfully connected and pinged.")
	}

	// instantiate a Database and a Collection from a client
	collection := cli.Database("test").Collection("details")

	// For methods that return a single item, a SingleResult instance is returned:
	ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var result bson.D
	err = collection.FindOne(ctx, bson.D{{"name", "xxxxxx"}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		fmt.Println("record does not exist")
	} else if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
