package golang_test

import "fmt"

func Example_base() {
	a := 0x_65e5
	fmt.Printf("%%b: %b\n", a)
	fmt.Printf("%%d: %d\n", a)
	fmt.Printf("%%o: %o\n", a)
	fmt.Printf("%%O: %O\n", a)
	fmt.Printf("%%x: %x\n", a)
	fmt.Printf("%%X: %X\n", a)
	fmt.Printf("%%U: %U\n", a)

	// Output:
	// %b: 110010111100101
	// %d: 26085
	// %o: 62745
	// %O: 0o62745
	// %x: 65e5
	// %X: 65E5
	// %U: U+65E5
}
