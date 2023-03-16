package xsync_test

import (
	"math/rand"
	"sync"
	"testing"

	"github.com/dushaoshuai/go-usage-examples/golang/sync/xsync"
)

func TestMap(t *testing.T) {
	var xmap xsync.Map[int, int]
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10_000_000; i++ {
			xmap.Delete(rand.Int())
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10_000_000; i++ {
			xmap.Add(rand.Int(), rand.Int())
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10_000_000; i++ {
			xmap.Load(rand.Int())
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10_000_000; i++ {
			xmap.Load2(rand.Int())
		}
	}()

	wg.Wait()
}
