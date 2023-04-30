package overflow_test

import "fmt"

func Example_signed_integer_overflow() {
	var vInt8 int8 = 126
	fmt.Printf("vInt8 : %d\n", vInt8)
	fmt.Printf("vInt8 + 1 : %d\n", vInt8+1)
	fmt.Printf("vInt8 + 2 : %d\n", vInt8+2)

	vInt8 = -127
	fmt.Printf("vInt8 : %d\n", vInt8)
	fmt.Printf("vInt8 - 1 : %d\n", vInt8-1)
	fmt.Printf("vInt8 - 2 : %d\n", vInt8-2)

	fmt.Printf("vInt8 : %d\n", vInt8)
	vInt8 -= 2
	fmt.Printf("vInt8 -= 2: %d\n", vInt8)

	// Output:
	// vInt8 : 126
	// vInt8 + 1 : 127
	// vInt8 + 2 : -128
	// vInt8 : -127
	// vInt8 - 1 : -128
	// vInt8 - 2 : 127
	// vInt8 : -127
	// vInt8 -= 2: 127
}

func Example_unsigned_integer_overflow() {
	var uint8V uint8 = 254
	fmt.Printf("uint8V : %d\n", uint8V)
	fmt.Printf("uint8V + 1 : %d\n", uint8V+1)
	fmt.Printf("uint8V + 2 : %d\n", uint8V+2)

	uint8V = 1
	fmt.Printf("uint8V : %d\n", uint8V)
	fmt.Printf("uint8V - 1 : %d\n", uint8V-1)
	fmt.Printf("uint8V - 2 : %d\n", uint8V-2)

	// Output:
	// uint8V : 254
	// uint8V + 1 : 255
	// uint8V + 2 : 0
	// uint8V : 1
	// uint8V - 1 : 0
	// uint8V - 2 : 255
}
