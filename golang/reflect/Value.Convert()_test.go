package reflect_test

import (
	"fmt"
	"reflect"
)

type myBytes []byte
type myRunes []rune

// https://go.dev/ref/spec#Conversions_to_and_from_a_string_type
func ExampleValue_Convert() {
	var myStrVal myString = "hello"
	var bytesVal = []byte("hello")
	var myBytesVal = myBytes("hello")
	var runesVal = []rune("hello")
	var myRunesVal = myRunes("hello")
	var intVal = 65

	f := func(v any) {
		rv := reflect.ValueOf(v)
		if rv.CanConvert(stringReflectType) {
			fmt.Println(rv.Convert(stringReflectType).Interface())
		}
	}

	f(myStrVal)
	f(bytesVal)
	f(myBytesVal)
	f(runesVal)
	f(myRunesVal)
	f(intVal)

	// Output:
	// hello
	// hello
	// hello
	// hello
	// hello
	// A
}
