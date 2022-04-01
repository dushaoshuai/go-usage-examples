package golang_test

import "fmt"

func Example_shift() {
	fmt.Println(1<<16, 64<<10)
	// Output:
	// 65536 65536
}
