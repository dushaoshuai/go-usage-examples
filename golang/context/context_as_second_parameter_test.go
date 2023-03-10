package context_test

import (
	"context"
	"fmt"
	"time"
)

// "The Context should be the first parameter" 是一个建议, 而不是强制规定.
// https://stackoverflow.com/questions/50522658/why-recommend-use-ctx-as-the-first-parameter.

func Example_context_as_second_parameter() {
	f := func(x int, ctx context.Context, y int) (int, error) {
		select {
		case <-time.After(13 * time.Second):
			return x + y, nil
		case <-ctx.Done():
			return 0, ctx.Err()
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println(f(3, ctx, 4))

	// Output:
	// 0 context deadline exceeded
}
