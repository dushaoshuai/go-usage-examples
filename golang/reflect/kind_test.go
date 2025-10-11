package reflect_test

import (
	"fmt"
	"reflect"
)

func ExampleKind() {
	for _, v := range []interface{}{"hi", 42, func() {}} {
		switch v := reflect.ValueOf(v); v.Kind() {
		case reflect.String:
			fmt.Println(v.String())
		case reflect.Int, reflect.Int8, reflect.Int16,
			reflect.Int32, reflect.Int64:
			fmt.Println(v.Int())
		default:
			fmt.Println("unhandled kind:", v.Kind())
		}
	}
	// Output:
	// hi
	// 42
	// unhandled kind: func
}

func ExampleKind_string() {
	var strVal = "hello"
	var myStrVal myString = "hello"
	var bytesVal = []byte("hello")
	var myBytesVal = myBytes("hello")
	var runesVal = []rune("hello")
	var myRunesVal = myRunes("hello")
	var intVal = 65

	f := func(v any) {
		if rv := reflect.ValueOf(v); rv.Kind() == reflect.String {
			fmt.Printf("type: %v, value: %v\n", rv.Type().Name(), rv.Interface())
		}
	}

	f(strVal)
	f(myStrVal)
	f(bytesVal)
	f(myBytesVal)
	f(runesVal)
	f(myRunesVal)
	f(intVal)

	// Output:
	// type: string, value: hello
	// type: myString, value: hello
}
