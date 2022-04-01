package reflect_test

import (
	"fmt"
	"reflect"
)

func ExampleValue_Interface() {
	for _, v := range []reflect.Value{
		reflect.ValueOf(3),
		reflect.ValueOf(4.4),
		reflect.ValueOf("hi"),
		reflect.ValueOf(struct{}{}),
		reflect.ValueOf([]string{}),
		reflect.ValueOf(func() {}),
	} {
		printLn(v)
	}
	// Output:
	// 3 3
	// 4.4 4.4
	// hi hi
	// {} {}
	// [] []
	// 0x10e83a0 0x10e83a0
}

func printLn(v reflect.Value) {
	fmt.Println(v, v.Interface())
}
