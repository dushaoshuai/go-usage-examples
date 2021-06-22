// https://docs.mongodb.com/drivers/go/
// https://www.mongodb.com/blog/search/golang%20quickstart
// https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo?utm_source=godoc
package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cli, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Print(err)
		return
	}
	defer func() {
		if err = cli.Disconnect(ctx); err != nil {
			log.Print(err)
			return
		}
	}()

	// ping
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = cli.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Print(err)
		return
	} else {
		log.Print("a MongoDB server has been found and connected to !!!")
	}
}
