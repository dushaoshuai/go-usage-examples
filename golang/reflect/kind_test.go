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
