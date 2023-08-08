package redis_test

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

// Redis pipelining 为什么快？
// https://redis.io/docs/manual/pipelining/
//
// 1. 节省了很多 RTT(round-trip time)
// 2. 节省了很多系统调用（read(2), write(2) -- socket I/O）

func Example_pipelining() {
	ctx, cancel := defaultCtx()
	defer cancel()

	rdb := mustNewDB(ctx, 15)
	defer func() {
		err := rdb.FlushDB(ctx).Err()
		if err != nil {
			panic(err)
		}
	}()

	// 有两种拿到结果的方法：
	//
	// 1. Pipelined() 的返回值：results
	// 2. 直接获得的每个命令的结果：getRes1，setRes，getRes2 ...
	var getRes1 *redis.StringCmd
	var setRes *redis.StatusCmd
	var getRes2 *redis.StringCmd
	var lpushRes *redis.IntCmd
	var lrangeRes *redis.StringSliceCmd
	results, err := rdb.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		getRes1 = pipe.Get(ctx, "nosuchkey")
		setRes = pipe.Set(ctx, "key", "value", 0)
		getRes2 = pipe.Get(ctx, "key")
		lpushRes = pipe.LPush(ctx, "list", 3, 2, 1)
		lrangeRes = pipe.LRange(ctx, "list", 0, -1)
		return nil
	})
	if err != nil {
		// Note: 这个 error 是 redis.Nil
		// 因为 pipeline 里的第一个命令 GET 返回的是 redis.Nil
		// 见 (*redis.Pipeline).Exec()
		fmt.Println(err)
		fmt.Println()
	}
	// check results are getRes1, setRes, getRes2 ...
	r1 := results[0].(*redis.StringCmd)
	fmt.Println(getRes1.Result())
	fmt.Println(r1.Result())
	fmt.Println()
	r2 := results[1].(*redis.StatusCmd)
	fmt.Println(setRes.Result())
	fmt.Println(r2.Result())
	fmt.Println()
	r3 := results[2].(*redis.StringCmd)
	fmt.Println(getRes2.Result())
	fmt.Println(r3.Result())
	fmt.Println()
	r4 := results[3].(*redis.IntCmd)
	fmt.Println(lpushRes.Result())
	fmt.Println(r4.Result())
	fmt.Println()
	r5 := results[4].(*redis.StringSliceCmd)
	fmt.Println(lrangeRes.Result())
	fmt.Println(r5.Result())
	fmt.Println()

	// Output:
	// redis: nil
	//
	//  redis: nil
	//  redis: nil
	//
	// OK <nil>
	// OK <nil>
	//
	// value <nil>
	// value <nil>
	//
	// 3 <nil>
	// 3 <nil>
	//
	// [1 2 3] <nil>
	// [1 2 3] <nil>
}
