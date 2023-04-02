package defer_test

import (
	"fmt"
)

func Example_defer_parameters_evaluation() {
	x := 10

	defer func(v int) {
		fmt.Println(v) // 10
	}(x)

	defer func() {
		fmt.Println(x) // 11
	}()

	x++

	// Output:
	// 11
	// 10
}
