package redis_test

import "fmt"

func ExampleStatusCmd() {
	ctx, cancel := defaultCtx()
	defer cancel()

	rdb := mustNew(ctx)

	// 似乎返回的结果不用做进一步解析的 Redis 命令，
	// 会用这个 StatusCmd ?
	statusCmd := rdb.MSet(ctx,
		"ExampleStatusCmd_key1", "val1",
		"ExampleStatusCmd_key2", "val2",
		"ExampleStatusCmd_key3", "val3",
	)
	result, err := statusCmd.Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	// Output:
	// OK
}
