package fmt_test

import (
	"fmt"
)

const poundFlagFormat = "%#v\n"

// %#v: a Go-syntax representation of the value
func Example_sharp_flag() {
	fmt.Printf(poundFlagFormat, 3)
	fmt.Printf(poundFlagFormat, 3.4)
	fmt.Printf(poundFlagFormat, []int{1, 2, 3, 4})
	fmt.Printf(poundFlagFormat, interface{}(func() int { return 5 }))
	fmt.Printf(poundFlagFormat, struct {
		A int
		B string
		C float64
	}{})
	// Output:
	// 3
	// 3.4
	// []int{1, 2, 3, 4}
	// (func() int)(0x10e3fa0)
	// struct { A int; B string; C float64 }{A:0, B:"", C:0}
}
