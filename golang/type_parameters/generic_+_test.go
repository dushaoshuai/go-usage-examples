package type_parameters_test

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type addable interface {
	constraints.Ordered
}

func add[T addable](a, b T) T {
	return a + b
}

func Example_add_operator() {
	fmt.Println(add[int](3, 2))
	fmt.Println(add(4.5, 6.9))           // type inference
	fmt.Println(add("hello ", "world!")) // type inference
	// Output:
	// 5
	// 11.4
	// hello world!
}
