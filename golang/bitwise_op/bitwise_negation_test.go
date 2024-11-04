package bitwise_op

import (
	"fmt"
)

// bitwise negation/complement (^) treats its operands as just bit patterns with
// no concept of arithmetic carry or sign.
func ExampleName_bitwise_negation() {
	var u8 uint8 = 12
	var positiveI8 int8 = 12
	var negativeI8 int8 = -12

	fmt.Printf("uint8 = %s, ^uint8 = %s\n", integerBits(u8), integerBits(^u8))
	fmt.Printf("positiveI8 = %s, ^positiveI8 = %s\n", integerBits(positiveI8), integerBits(^positiveI8))
	fmt.Printf("negativeI8 = %s, ^negativeI8 = %s\n", integerBits(negativeI8), integerBits(^negativeI8))

	// Output:
	// uint8 = 1100, ^uint8 = 11110011
	// positiveI8 = 1100, ^positiveI8 = 11110011
	// negativeI8 = 11110100, ^negativeI8 = 1011
}
