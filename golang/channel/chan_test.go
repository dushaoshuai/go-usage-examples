package channel_test

import "fmt"

func Example_ok() {
	ch := make(chan int, 1)
	ch <- 100
	close(ch)

	x, ok := <-ch
	fmt.Println(x, ok)

	// Output:
	// 100 true
}
