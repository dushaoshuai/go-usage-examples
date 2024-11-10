package bitwise_op

import (
	"fmt"
)

func Example_right_shift_and_divisor_is_a_power_of_2() {
	var u8 uint8 = 77
	var i81 int8 = 77
	var i82 int8 = -77

	fmt.Printf("%d >> 1 = %d\n", u8, u8>>1)
	fmt.Printf("%d / 2 = %d\n", u8, u8/2)

	fmt.Printf("%d >> 1 = %d\n", i81, i81>>1)
	fmt.Printf("%d / 2 = %d\n", i81, i81/2)

	fmt.Printf("%d >> 1 = %d\n", i82, i82>>1)
	fmt.Printf("%d / 2 = %d\n", i82, i82/2)

	// Output:
	// 77 >> 1 = 38
	// 77 / 2 = 38
	// 77 >> 1 = 38
	// 77 / 2 = 38
	// -77 >> 1 = -39
	// -77 / 2 = -38
}
