package bitwise

import (
	"fmt"
)

// bitwise or (|) treats its operands as just bit patterns with
// no concept of arithmetic carry or sign.
func ExampleName_bitwise_or() {
	var u81 uint8 = 12
	var u82 uint8 = 35
	var u8Or uint8 = u81 | u82

	var positiveI8 int8 = 12
	var negativeI8 int8 = -35
	var i8Or int8 = positiveI8 | negativeI8

	fmt.Printf("%s | %s = %s\n", integerBits(u81), integerBits(u82), integerBits(u8Or))
	fmt.Printf("%s | %s = %s\n", integerBits(positiveI8), integerBits(negativeI8), integerBits(i8Or))

	// Output:
	// 1100 | 100011 = 101111
	// 1100 | 11011101 = 11011101
}
