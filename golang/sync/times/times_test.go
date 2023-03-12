package times_test

import (
	"fmt"
	"sync"
	"time"

	"github.com/dushaoshuai/go-usage-examples/golang/sync/times"
)

var (
	Times = times.NewTimes(3 * time.Second)
	t     time.Time
)

func ExampleTimes() { // todo test
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 20; i++ {
			Times.Do(func() error {
				t = time.Now()
				return nil
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
			Times.Do(func() error {
				t = time.Now()
				return nil
			})
			fmt.Println(t)
			time.Sleep(500 * time.Millisecond)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 20; i++ {
			Times.Do(func() error {
				t = time.Now()
				return nil
			})
			fmt.Println(t)
			time.Sleep(1 * time.Second)
		}
	}()
	time.Sleep(time.Second)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 30; i++ {
			Times.Do(func() error {
				t = time.Now()
				return nil
			})
			fmt.Println(t)
			time.Sleep(2 * time.Second)
		}
	}()
	wg.Wait()

	// Output:
}
