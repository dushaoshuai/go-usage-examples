package interface_test

import (
	"fmt"
)

func Example_assert_core_type() {
	type t string
	var vt t = "value"

	var va any
	va = vt

	{
		v, ok := va.(t)
		fmt.Println(v, ok)
	}
	{
		v, ok := va.(string)
		fmt.Println(v, ok)
	}

	// Output:
	// value true
	//  false
}
