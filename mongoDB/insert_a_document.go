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
	cli, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = cli.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	// ping to verify that the database has been connected
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = cli.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Print(err)
		return
	} else {
		log.Print("a MongoDB server has been found and connected to")
	}

	// To insert a document into a collection, first retrieve
	// a Database and then Collection instance from the Client:
	collection := cli.Database("test").Collection("details")
	// The Collection instance can then be used to insert documents:
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, bson.D{{"name", "pi"}, {"value", 3.14159}})
	if err != nil {
		log.Print(err)
		return
	}
	id := res.InsertedID
	fmt.Printf("the inserted id is %v\n", id)

	person := struct {
		Name    string
		Gender  string
		Hobbies []string
		Age     int64
	}{"Lihua", "male", []string{"computer", "books"}, 56}
	res, err = collection.InsertOne(ctx, person)
	if err != nil {
		log.Print(err)
		return
	}
	fmt.Printf("the inserted id is %v\n", res.InsertedID)
}
