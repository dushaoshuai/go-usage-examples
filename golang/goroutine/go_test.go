package goroutine_test

import (
	"fmt"
	"sync"
)

func Example_go_statement_parameters_evaluation() {
	var (
		c = make(chan struct{})
		x = 10

		wg sync.WaitGroup
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-c
		fmt.Println(x) // 11
	}()

	wg.Add(1)
	go func(v int) { // The go statement itself completes immediately. The function value and parameters are evaluated as usual in the calling goroutine.
		defer wg.Done()
		<-c
		fmt.Println(v) // 10
	}(x)

	x++
	close(c)
	wg.Wait()

	// Unordered Output:
	// 10
	// 11
}
