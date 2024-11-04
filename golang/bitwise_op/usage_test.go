package bitwise_op

import "fmt"

func Example_mask() {
	var mask uint8 = 0b11010011
	var flags1 uint8 = 0b10101111
	var flags2 uint8 = 0b1111
	var flags3 uint8 = 0b11111000

	fmt.Printf("%s & %s = %s\n", integerBits(flags1), integerBits(mask), integerBits(flags1&mask))
	fmt.Printf("%s & %s = %s\n", integerBits(flags2), integerBits(mask), integerBits(flags2&mask))
	fmt.Printf("%s & %s = %s\n", integerBits(flags3), integerBits(mask), integerBits(flags3&mask))

	// Output:
	// 10101111 & 11010011 = 10000011
	// 1111 & 11010011 = 11
	// 11111000 & 11010011 = 11010000
}
