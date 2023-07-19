package redis_test

import (
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func ExampleSetArgs() {
	ctx, cancel := defaultCtx()
	defer cancel()

	rdb := mustNew(ctx)

	key := "thereIsNeverSuchKey"
	val := "valueOfNoSuchKey"
	defer func() {
		result, err := rdb.Del(ctx, key).Result()
		if err != nil {
			panic(err)
		}
		if result != 1 {
			panic("Redis DEL failed")
		}
	}()

	result, err := rdb.SetArgs(ctx, key, val, redis.SetArgs{
		Mode:     "NX",             // Only set the key if it does not already exist.
		TTL:      time.Second * 10, // EX or PX
		ExpireAt: time.Time{},      // EXAT or PXAT
		Get:      false,            // Return the old string stored at key, or nil if key did not exist. An error is returned and SET aborted if the value stored at key is not a string.
		KeepTTL:  false,            // Retain the time to live associated with the key.
	}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	// Output:
	// OK
}
