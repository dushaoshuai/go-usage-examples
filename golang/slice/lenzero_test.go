package slice_test

import (
	"fmt"
)

func Example_zero_length() {
	zs := make([]int, 0)
	fmt.Println(zs, len(zs), cap(zs))

	zs = append(zs, 18)
	fmt.Println(zs, len(zs), cap(zs))

	// Output:
	// [] 0 0
	// [18] 1 1
}
