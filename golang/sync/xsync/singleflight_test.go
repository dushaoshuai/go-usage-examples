package xsync_test

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/singleflight"
)

func ExampleSingleFlight() {
	var (
		g singleflight.Group

		dofn2 = make(chan struct{})
	)

	go func() {
		v1, err1, shared1 := g.Do("key1", func() (interface{}, error) {
			dofn2 <- struct{}{}

			// wait second Do to start
			time.Sleep(3 * time.Second)

			return 5, nil
		})
		fmt.Println(v1, err1, shared1)
	}()

	<-dofn2
	v2, err2, shared2 := g.Do("key1", func() (interface{}, error) {
		return 0, errors.New("some error")
	})
	fmt.Println(v2, err2, shared2)

	// Output:
	// 5 <nil> true
	// 5 <nil> true
}
