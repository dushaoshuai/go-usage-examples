// https://go.dev/blog/laws-of-reflection
package reflect_test

import (
	"fmt"
	"reflect"
)

// The first law of reflection:
// Reflection goes from interface value to reflection object.

func ExampleTypeOf() {
	var x = 3.4
	fmt.Println("type:", reflect.TypeOf(x), reflect.TypeOf(x).String())
	// Output:
	// type: float64 float64
}

func ExampleValueOf() {
	var x = 3.4
	fmt.Println("value:", reflect.ValueOf(x), reflect.ValueOf(x).String())

	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
	// Output:
	// value: 3.4 <float64 Value>
	// type: float64
	// kind is float64: true
	// value: 3.4
}

// 函数 myFunc 的类型是 non-defined type
func myFunc() string { return "hello" }

func Example_func_name() {
	fmt.Println("type:", reflect.TypeOf(myFunc))
	fmt.Println("name:", reflect.TypeOf(myFunc).Name())
	// Output:
	// type: func() string
	// name:
}
