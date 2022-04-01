package fmt_test

import "fmt"

type X string

// Sprintf format %s with arg x causes recursive String method call
// func (x X) String() string {
// 	return fmt.Sprintf("<%s>", x)
// }

func (x X) String() string {
	return fmt.Sprintf("<%s>", string(x))
}

func ExampleStringer_recursion() {
	fmt.Println(X("Hi, Gopher!"))
	// Output:
	// <Hi, Gopher!>
}
