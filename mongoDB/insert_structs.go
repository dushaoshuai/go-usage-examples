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

type person struct {
	Name    string   `bson:"name, omitempty"`
	Gender  string   `bson:"gender, omitempty"`
	Hobbies []string `bson:"hobbies, omitempty"`
	Age     int64    `bson:"age, omiempty"`
}

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

	coll := cli.Database("test").Collection("inventory")

	people := []person{
		person{"LiHua", "male", []string{"computer", "reading"}, 56},
		person{"LiMing", "male", []string{"reading", "basketball"}, 34},
		person{"ZhangHong", "female", []string{"basketball", "baseball"}, 25},
	}
	data := make([]interface{}, len(people))
	for i := range people {
		data[i] = &people[i]
	}
	res, err := coll.InsertMany(context.Background(), data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("the inserted id is %v\n", res.InsertedIDs)

	people = []person{}
	cursor, err := coll.Find(
		context.Background(),
		bson.D{
			{"gender", bson.M{
				"$exists": true,
			}},
		})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(context.Background(), &people); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", people)
}
