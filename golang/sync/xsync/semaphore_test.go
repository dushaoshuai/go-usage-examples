package xsync_test

import (
	"context"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"

	"golang.org/x/sync/semaphore"
)

func TestSemaphore(t *testing.T) {
	var (
		maxWorkers = runtime.GOMAXPROCS(8)
		sem        = semaphore.NewWeighted(int64(maxWorkers))

		resource int32
		ctx      = context.Background()
	)

	go func() {
		for {
			if r := atomic.LoadInt32(&resource); r > int32(maxWorkers) {
				t.Errorf("resource > maxWorkers, %v > %v", r, maxWorkers)
			}
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < 1<<20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			_ = sem.Acquire(ctx, 1)
			defer sem.Release(1)

			defer atomic.AddInt32(&resource, -1)
			atomic.AddInt32(&resource, 1)
		}()
	}

	wg.Wait()
}
