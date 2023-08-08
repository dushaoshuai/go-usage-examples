package redis_test

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/redis/go-redis/v9"
)

func defaultCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}

func mustNewDB(ctx context.Context, DB int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     net.JoinHostPort("127.0.0.1", "6379"),
		Password: "", // no password set
		DB:       DB,
	})

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	log.Println(pong)

	return rdb
}

func mustNew(ctx context.Context) *redis.Client {
	return mustNewDB(ctx, 0)
}

func Example_quick_start() {
	ctx, cancel := defaultCtx()
	defer cancel()

	rdb := mustNew(ctx)

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}

	// Output:
	// key value
	// key2 does not exist
}
