package reflect_test

import (
	"fmt"
	"reflect"
)

// The zero Value represents no value. Its IsValid method returns false,
// its Kind method returns Invalid, its String method returns "<invalid Value>",
// and all other methods panic. Most functions and methods never return an invalid value.
// If one does, its documentation states the conditions explicitly.
func Example_zero_Value() {
	v := reflect.Value{}
	fmt.Println(v.IsValid(), v.Kind(), v.String())
	// Output:
	// false invalid <invalid Value>
}
