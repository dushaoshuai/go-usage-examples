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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cli, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017").)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = cli.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	// ping to verify that a mongoDB serer has been connected
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err = cli.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	} else {
		log.Print("a MongoDB server has been connected !!!")
	}

	// instantiate a Database and a Collection form the client:
	collection := cli.Database("test").Collection("details")

	// several query methods return a cursor, which can be used like this:
	ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var result bson.D
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}
	if err = cursor.Err(); err != nil {
		log.Fatal(err)
	}
}
