package g_scoped_to_func_test

import (
	"sync"
	"testing"
	"time"
)

// https://github.com/golang/go/blob/dceee2e983f5dab65c3905ecf40e70e15cf41b7d/src/sync/mutex_test.go#L198
func Test_g_scoped_to_func_channel_cancellation(t *testing.T) {
	var mu sync.Mutex
	stop := make(chan bool)
	defer close(stop)
	go func() {
		for {
			mu.Lock()
			time.Sleep(100 * time.Microsecond)
			mu.Unlock()
			select {
			case <-stop:
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
