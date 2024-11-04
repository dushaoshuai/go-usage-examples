package bitwise

import (
	"fmt"
)

// bitwise And (&) treats its operands as just bit patterns with
// no concept of arithmetic carry or sign.
func ExampleName_bitwise_and() {
	var u81 uint8 = 12
	var u82 uint8 = 35
	var u8And uint8 = u81 & u82

	var positiveI8 int8 = 12
	var negativeI8 int8 = -35
	var i8And int8 = positiveI8 & negativeI8

	fmt.Printf("%s & %s = %s\n", integerBits(u81), integerBits(u82), integerBits(u8And))
	fmt.Printf("%s & %s = %s\n", integerBits(positiveI8), integerBits(negativeI8), integerBits(i8And))

	// Output:
	// 1100 & 100011 = 0
	// 1100 & 11011101 = 1100
}
