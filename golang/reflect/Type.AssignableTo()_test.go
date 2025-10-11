package reflect_test

import (
	"fmt"
	"reflect"
)

type myString string

var (
	stringReflectType = reflect.TypeOf("")
)

// https://go.dev/ref/spec#Assignability
// V and T have identical underlying types but are not type parameters and at least one of V or T is not a named type.
func ExampleType_AssignableTo() {
	var vStr string
	var vMyStr myString = "hello"
	// vStr = vMyStr // Cannot use 'vMyStr' (type myString) as the type string

	myStringRt := reflect.TypeOf(vMyStr)
	if myStringRt.AssignableTo(stringReflectType) {
		fmt.Println("AssignableTo")
		reflect.ValueOf(&vStr).Elem().Set(reflect.ValueOf(vMyStr))
	}

	fmt.Println(vStr)

	// Output:
}
