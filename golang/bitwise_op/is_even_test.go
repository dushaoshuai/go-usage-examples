package bitwise_op

import (
	"fmt"
)

// 判断一个数是不是偶数，正负都适用.
func isEven(n int) bool {
	return n&1 == 0
}

func Example_is_even() {
	for _, num := range []int{
		-100, -99, -98, -97, -96, -95, -94, -90,
		100, 99, 98, 97, 96, 95, 94, 90,
		-3, -2, -1, 0, 1, 2, 3,
	} {
		if isEven(num) {
			fmt.Printf("%d is even\n", num)
		} else {
			fmt.Printf("%d is odd\n", num)
		}
	}

	// Output:
	// -100 is even
	// -99 is odd
	// -98 is even
	// -97 is odd
	// -96 is even
	// -95 is odd
	// -94 is even
	// -90 is even
	// 100 is even
	// 99 is odd
	// 98 is even
	// 97 is odd
	// 96 is even
	// 95 is odd
	// 94 is even
	// 90 is even
	// -3 is odd
	// -2 is even
	// -1 is odd
	// 0 is even
	// 1 is odd
	// 2 is even
	// 3 is odd
}
