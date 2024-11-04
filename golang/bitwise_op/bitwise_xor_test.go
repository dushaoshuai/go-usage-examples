package bitwise_op

import (
	"fmt"
)

// bitwise XOR (^) treats its operands as just bit patterns with
// no concept of arithmetic carry or sign.
func ExampleName_bitwise_xor() {
	var u81 uint8 = 12
	var u82 uint8 = 35
	var u8Xor uint8 = u81 ^ u82

	var positiveI8 int8 = 12
	var negativeI8 int8 = -35
	var i8Xor int8 = positiveI8 ^ negativeI8

	fmt.Printf("%s ^ %s = %s\n", integerBits(u81), integerBits(u82), integerBits(u8Xor))
	fmt.Printf("%s ^ %s = %s\n", integerBits(positiveI8), integerBits(negativeI8), integerBits(i8Xor))

	// Output:
	// 1100 ^ 100011 = 101111
	// 1100 ^ 11011101 = 11010001
}
