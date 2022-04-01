package type_parameters_test

import (
	"fmt"
)

// printSlice prints the elements of any slice.
// printSlice has a type parameter T and has a single (non-type)
// parameter s which is a slice of that type parameter.
func printSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

// Call printSlice with a []int.
// printSlice has a type parameter T, and we want to pass a []int,
// so we pass a type argument of int by writing printSlice[int].
// The function print[int] expects a []int as an argument.
func Example_print_slice() {
	printSlice[int]([]int{1, 2, 3})
	printSlice([]int{1, 2, 3})
	// Output:
	// 1
	// 2
	// 3
	// 1
	// 2
	// 3
}

func test_any[T any](x T, ok bool) T {
	// The operations permitted for any type are:
	// * declare variables of those types
	var a T
	// * assign other values of the same type to those variables
	a = x
	// * pass those variables to functions or return them from functions
	fmt.Println(a, x)
	if ok {
		return a
	}
	// * take the address of those variables
	_ = &a
	// * convert or assign values of those types to the type interface{}
	fmt.Println(interface{}(x))
	var b interface{}
	b = a
	fmt.Println(b)
	// * convert a value of type T to type T (permitted but useless)
	a = T(x)
	fmt.Println(a)
	// * use a type assertion to convert an interface value to the type
	c := interface{}(6)
	a, ok = c.(T)
	if ok {
		fmt.Println(a)
	}
	// * use the type as a case in a type switch
	switch c.(type) {
	case T:
		fmt.Println("type switch to T")
	case int:
		fmt.Println("type switch to int")
	}
	// * define and use composite types that use those types, such as a slice of that type
	d := make([]T, 10)
	for i := range d {
		d[i] = any(i).(T)
	}
	printSlice(d)
	// * pass the type to some predeclared functions such as new
	e := new(T)
	*e = any(5).(T)
	fmt.Println(*e)

	return a
}

func Example_operations_permitted_for_any_type() {
	fmt.Println(test_any(6, false))
	// Output:
	// 6 6
	// 6
	// 6
	// 6
	// 6
	// type switch to T
	// 0
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
	// 5
	// 6
}
