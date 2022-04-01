package fmt_test

import "fmt"

type foo struct {
	A int
	B string
	C float64
	D interface{}
}

func Example_plus_flag() {
	f := foo{}
	// %v: the value in a default format
	fmt.Printf("%v\n", f)
	// when printing structs, the plus flag (%+v) adds field names
	fmt.Printf("%+v\n", f)
	// %#v: a Go-syntax representation of the value
	fmt.Printf("%#v\n", f)
	// Output:
	// {0  0 <nil>}
	// {A:0 B: C:0 D:<nil>}
	// fmt_test.foo{A:0, B:"", C:0, D:interface {}(nil)}
}
