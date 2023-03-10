package context_test

import (
	"context"
	"fmt"
	"time"
)

func Example_ctx() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := acceptCtx(ctx); err != nil {
		fmt.Println(err)
	}

	// Output:
	// context deadline exceeded
}

func acceptCtx(ctx context.Context) error {
	select {
	case <-time.After(13 * time.Second):
		fmt.Println("13 seconds passed")
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
