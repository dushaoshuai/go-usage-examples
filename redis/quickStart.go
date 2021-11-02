package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		logrus.Fatalln(err)
	}
	fmt.Println(pong)

	err = rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		logrus.Fatal(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Println("key:", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exists")
	} else if err != nil {
		logrus.Fatal(err)
	} else {
		fmt.Println("key2", val2)
	}

	//	Output:
	//	PONG
	//	key: value
	//	key2 does not exists
}
