package unsafe_test

import "fmt"

func Example_unsafe() {
	a := interface{}(5)
	fmt.Printf("%v %[1]T\n", a)
	// Output:
}
