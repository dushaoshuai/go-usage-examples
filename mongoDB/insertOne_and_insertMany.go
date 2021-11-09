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
	cli, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = cli.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	// ping to verify that the database has been connected
	// Using Ping reduces application resilience because applications starting up will error if the
	// server is temporarily unavailable or is failing over (e.g. during autoscaling due to a load spike).
	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel()
	err = cli.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print("a MongoDB server has been found and connected to")
	}

	// To insert a document into a collection, first retrieve
	// a Database and the Collection instance form the client:
	coll := cli.Database("test").Collection("inventory")
	// The Collection instance can then  be used to insert documents:
	result, err := coll.InsertOne(
		context.Background(),
		bson.D{
			{"item", "canvas"},
			{"qty", 100},
			{"tags", bson.A{"cotton"}},
			{"size", bson.D{
				{"h", 28},
				{"w", 35.5},
				{"uom", "cm"},
			}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	id := result.InsertedID
	fmt.Printf("the inserted id is %v\n", id)

	// insert multiple documents
	resultMany, err := coll.InsertMany(
		context.Background(),
		[]interface{}{
			bson.D{
				{"item", "journal"},
				{"qty", int32(25)},
				{"tags", bson.A{"blank", "red"}},
				{"size", bson.D{
					{"h", 14},
					{"w", 21},
					{"uom", "cm"},
				}},
			},
			bson.D{
				{"item", "mat"},
				{"qty", int32(25)},
				{"tags", bson.A{"gray"}},
				{"size", bson.D{
					{"h", 27.9},
					{"w", 35.5},
					{"uon", "cm"},
				}},
			},
			bson.D{
				{"item", "mousepad"},
				{"qty", 25},
				{"tags", bson.A{"gel", "blue"}},
				{"size", bson.D{
					{"h", 19},
					{"w", 22.85},
					{"uom", "cm"},
				}},
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	id = resultMany.InsertedIDs
	fmt.Println(id)
}
