package for_loop_test

import (
	"fmt"
)

// https://go.dev/ref/spec#RangeClause:~:text=//%20fibo%20generates%20the%20Fibonacci%20sequence
func Example_fibo() {
	// fibo generates the Fibonacci sequence
	fibo := func(yield func(x int) bool) {
		f0, f1 := 0, 1
		for yield(f0) {
			f0, f1 = f1, f0+f1
		}
	}

	// print the Fibonacci numbers below 1000:
	for x := range fibo {
		if x >= 1000 {
			break
		}
		fmt.Println(x)
	}

	// Output:
	// 0
	// 1
	// 1
	// 2
	// 3
	// 5
	// 8
	// 13
	// 21
	// 34
	// 55
	// 89
	// 144
	// 233
	// 377
	// 610
	// 987
}
