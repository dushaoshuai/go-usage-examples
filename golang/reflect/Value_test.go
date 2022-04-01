package reflect_test

import (
	"fmt"
	"reflect"
)

func ExampleValue() {
	s := reflect.ValueOf([]int{})

	x1 := reflect.ValueOf(1)
	x2 := reflect.ValueOf(2)
	x3 := reflect.ValueOf(3)

	s = reflect.Append(s, x1, x2, x3)

	switch s.Kind() {
	case reflect.Slice:
		fmt.Printf("cap: %v\n", s.Cap())
		fmt.Printf("len: %v\n", s.Len())
		l := s.Len()
		for i := 0; i < l; i++ {
			fmt.Printf("%v ", s.Index(i).Int())
		}
		fmt.Println()
	}
	// Output:
	// cap: 3
	// len: 3
	// 1 2 3
}
