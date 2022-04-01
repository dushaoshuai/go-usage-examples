// https://go.dev/blog/laws-of-reflection
package reflect_test

import (
	"fmt"
	"reflect"
)

// The first law of reflection:
// Reflection goes from interface value to reflection object.

func ExampleTypeOf() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x), reflect.TypeOf(x).String())
	// Output:
	// type: float64 float64
}

func ExampleValueOf() {
	var x float64 = 3.4
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
