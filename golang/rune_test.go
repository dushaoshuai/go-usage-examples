package golang_test

import (
	"fmt"
	"reflect"
)

func Example_rune_type() {
	fmt.Println(reflect.TypeOf('-').Kind(), reflect.TypeOf('-'))
	// Output:
	// int32 int32
}

func Example_rune_in_string() {
	fmt.Println("\u0000\u0045")
	// Output:
}
