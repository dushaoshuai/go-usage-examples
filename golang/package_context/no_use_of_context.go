package main

import (
	"context"
	"fmt"
	"time"
)

// noUseContext 说明只传递 context.Context 是不行的，必须使用它.
func noUseContext() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	justAcceptContext(ctx)
}

func justAcceptContext(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("overslept")
	}
}
