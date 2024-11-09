package bitwise_op

import (
	"fmt"
)

func ExampleName_bitwise_shift_unsigned() {
	var u81 uint8 = 12
	var u82 uint8 = 35

	for i := range 8 {
		fmt.Printf("%s >> %d = %s\n", integerBits(u81), i, integerBits(u81>>i))
	}
	for i := range 8 {
		fmt.Printf("%s << %d = %s\n", integerBits(u82), i, integerBits(u82<<i))
	}

	// Output:
	// 1100 >> 0 = 1100
	// 1100 >> 1 = 110
	// 1100 >> 2 = 11
	// 1100 >> 3 = 1
	// 1100 >> 4 = 0
	// 1100 >> 5 = 0
	// 1100 >> 6 = 0
	// 1100 >> 7 = 0
	// 100011 << 0 = 100011
	// 100011 << 1 = 1000110
	// 100011 << 2 = 10001100
	// 100011 << 3 = 11000
	// 100011 << 4 = 110000
	// 100011 << 5 = 1100000
	// 100011 << 6 = 11000000
	// 100011 << 7 = 10000000
}
