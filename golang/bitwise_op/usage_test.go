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

func Example_turning_bits_on() {
	var mask uint8 = 0b11010011
	var flags1 uint8 = 0b10101111
	var flags2 uint8 = 0b1111
	var flags3 uint8 = 0b11111000

	fmt.Printf("%s | %s = %s\n", integerBits(flags1), integerBits(mask), integerBits(flags1|mask))
	fmt.Printf("%s | %s = %s\n", integerBits(flags2), integerBits(mask), integerBits(flags2|mask))
	fmt.Printf("%s | %s = %s\n", integerBits(flags3), integerBits(mask), integerBits(flags3|mask))

	// Output:
	// 10101111 | 11010011 = 11111111
	// 1111 | 11010011 = 11011111
	// 11111000 | 11010011 = 11111011
}

// clearing bits
// bit clear
func Example_turning_bits_off() {
	var mask uint8 = 0b11010011
	var flags1 uint8 = 0b10101111
	var flags2 uint8 = 0b1111
	var flags3 uint8 = 0b11111000

	fmt.Printf("%s &^ %s = %s\n", integerBits(flags1), integerBits(mask), integerBits(flags1&^mask))
	fmt.Printf("%s &^ %s = %s\n", integerBits(flags2), integerBits(mask), integerBits(flags2&^mask))
	fmt.Printf("%s &^ %s = %s\n", integerBits(flags3), integerBits(mask), integerBits(flags3&^mask))

	// Output:
	// 10101111 &^ 11010011 = 101100
	// 1111 &^ 11010011 = 1100
	// 11111000 &^ 11010011 = 101000
}

func Example_toggling_bits() {
	var mask uint8 = 0b_1111_1111
	var flags1 uint8 = 0b_1010_1111
	var flags2 uint8 = 0b_1111
	var flags3 uint8 = 0b_1111_1000

	fmt.Printf("%s ^ %s = %s\n", integerBits(flags1), integerBits(mask), integerBits(flags1^mask))
	fmt.Printf("%s ^ %s = %s\n", integerBits(flags2), integerBits(mask), integerBits(flags2^mask))
	fmt.Printf("%s ^ %s = %s\n", integerBits(flags3), integerBits(mask), integerBits(flags3^mask))

	// Output:
	// 10101111 ^ 11111111 = 1010000
	// 1111 ^ 11111111 = 11110000
	// 11111000 ^ 11111111 = 111
}

func Example_checking_the_value_of_a_bit() {
	var flags uint8 = 0b_1010_1111

	fmt.Println(integerBits(flags))

	for i := range 8 {
		if flags&(1<<i) != 0 {
			fmt.Printf("%d is at position %d\n", 1, i)
		} else {
			fmt.Printf("%d is at position %d\n", 0, i)
		}
	}

	// Output:
	// 10101111
	// 1 is at position 0
	// 1 is at position 1
	// 1 is at position 2
	// 1 is at position 3
	// 0 is at position 4
	// 1 is at position 5
	// 0 is at position 6
	// 1 is at position 7
}
