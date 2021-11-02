package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

var (
	myZset = "2345"

	sponsorID = "71923741071"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	result := rdb.ZAdd(ctx, myZset,
		&redis.Z{
			Score:  40,
			Member: "12345566_0997979",
		}, &redis.Z{
			Score:  68,
			Member: "43854_13491834",
		}, &redis.Z{
			Score:  23,
			Member: "456564564_980890",
		})
	fmt.Println(result.Result())
	fmt.Println(result.String())
	printMembers(ctx, rdb, myZset)

	floatCmd := rdb.ZAddArgsIncr(ctx, myZset, redis.ZAddArgs{
		Members: []redis.Z{
			{90, "456564564_980890"},
		},
	})
	fmt.Println(floatCmd.Result())
	printMembers(ctx, rdb, myZset)

	intCmd := rdb.ZAdd(ctx, sponsorID,
		&redis.Z{
			Score:  0,
			Member: 28340289_17340143010,
		}, &redis.Z{
			Score:  0,
			Member: 28340289_10732804170,
		}, &redis.Z{
			Score:  0,
			Member: 28340289_70983714730,
		}, &redis.Z{
			Score:  0,
			Member: 28340289_70983728347,
		}, &redis.Z{
			Score:  0,
			Member: 28340289_25758715349,
		}, &redis.Z{
			Score:  0,
			Member: 28340289_87234091730,
		})
	fmt.Println(intCmd.Result())
	countMembersByLex(ctx, rdb, sponsorID, "28340289")
}

func printMembers(ctx context.Context, rdb *redis.Client, key string) {
	members, err := rdb.ZRange(ctx, key, 0, -1).Result()
	switch {
	case err == redis.Nil:
		logrus.Println("key \"myZset\" does not exist")
	case err != nil:
		logrus.Fatal(err)
	default:
		fmt.Println(members)
	}
}

func countMembersByLex(ctx context.Context, rdb *redis.Client, key string, prefix string) {
	intCmd := rdb.ZLexCount(ctx, key, prefix+"_00000000000", prefix+"_00000000000")
	fmt.Println(intCmd.Val())
}
