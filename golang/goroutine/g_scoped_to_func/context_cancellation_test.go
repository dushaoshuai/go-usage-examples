package g_scoped_to_func_test

import (
	"context"
	"sync"
	"testing"
	"time"
)

func Test_g_scoped_to_func_context_cancellation(t *testing.T) {
	var mu sync.Mutex

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // ensures the goroutine is asked to stop when this function returns

	go func() {
		for {
			mu.Lock()
			time.Sleep(100 * time.Microsecond)
			mu.Unlock()
			select {
			case <-ctx.Done():
				return
			default:
			}
		}
	}()

	done := make(chan bool, 1)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(100 * time.Microsecond)
			mu.Lock()
			mu.Unlock()
		}
		done <- true
	}()

	select {
	case <-done:
	case <-time.After(10 * time.Second):
		t.Fatalf("can't acquire Mutex in 10 seconds")
	}
}
