package oncevery_test

import (
	"fmt"
	"sync"
	"time"

	"github.com/dushaoshuai/go-usage-examples/golang/sync/onceEvery"
)

var (
	onceEvery = oncevery.NewOnceEvery(3 * time.Second)
	t         time.Time
)

func ExampleOnceEvery() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 20; i++ {
			onceEvery.Do(func() {
				t = time.Now()
			})
			fmt.Println(t)
			time.Sleep(500 * time.Millisecond)
		}
	}()
	time.Sleep(700 * time.Millisecond)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 20; i++ {
			onceEvery.Do(func() {
				t = time.Now()
			})
			fmt.Println(t)
			time.Sleep(500 * time.Millisecond)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 20; i++ {
			onceEvery.Do(func() {
				t = time.Now()
			})
			fmt.Println(t)
			time.Sleep(2 * time.Second)
		}
	}()
	time.Sleep(time.Second)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 30; i++ {
			onceEvery.Do(func() {
				t = time.Now()
			})
			fmt.Println(t)
			time.Sleep(2 * time.Second)
		}
	}()
	wg.Wait()

	// Output:
}
