package sync_test

import (
	"sync"
	"testing"
)

// It is allowed for one goroutine to lock a Mutex and then
// arrange for another goroutine to unlock it.
func TestMutex_another_goroutine_unlock(t *testing.T) {
	ch := make(chan struct{})
	var mu sync.Mutex

	go func() {
		mu.Lock()
		ch <- struct{}{}
	}()

	go func() {
		<-ch
		mu.Unlock()
	}()
}
