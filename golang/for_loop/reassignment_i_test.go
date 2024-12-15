package for_loop_test

import "fmt"

func Example_loop_reassignment_i() {
	for i := 0; i < 10; i++ {
		i = 9
		fmt.Println(i)
	}

	fmt.Println()
	for i := 0; i < 10; i++ {
		i := i
		i = -1
		fmt.Println(i)
	}

	// Output:
	// 9
	//
	// -1
	// -1
	// -1
	// -1
	// -1
	// -1
	// -1
	// -1
	// -1
	// -1
}
