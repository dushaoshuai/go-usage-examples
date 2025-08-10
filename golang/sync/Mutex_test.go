package sync_test

import (
	"sync"
	"testing"
)

func TestMutex_normal_Unlock(t *testing.T) {
	var mu sync.Mutex
	defer func() {
		if !mu.TryLock() {
			t.Fatal("Mutex.TryLock() failed")
		}
		mu.Unlock()
	}()

	mu.Lock()
	defer mu.Unlock()
}

// It is allowed for one goroutine to lock a Mutex and then
// arrange for another goroutine to unlock it.
func TestMutex_another_goroutine_unlock(t *testing.T) {
	ch := make(chan struct{})
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Go(func() {
		mu.Lock()
		ch <- struct{}{}
	})

	wg.Go(func() {
		<-ch
		mu.Unlock()
	})

	wg.Wait()

	// Ensure the mutex is indeed unlocked now.
	mu.Lock()
	mu.Unlock()
}
