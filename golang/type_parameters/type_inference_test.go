package type_parameters

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Function argument type inference.
// Function argument type inference only works for type parameters that are used in the function parameters,
// not for type parameters used only in function results or only in the function body.

// 这个函数适用函数参数类型推断
func min[T any](x, y T, less func(T, T) bool) T {
	if less(x, y) {
		return x
	}
	return y
}

func Example_generic_min() {
	fmt.Println(
		min[int](
			2,
			4,
			func(x, y int) bool {
				return x < y
			},
		),
	)
	fmt.Println(
		min( // 这里不用明确地传入 type argument
			2,
			4,
			func(x, y int) bool {
				return x < y
			},
		),
	)

	// Output:
	// 2
	// 2
}

// 这个函数不适用函数参数类型推断，
// 因为 type parameters used only in function results or only in the function body
func newT[T any]() *T {
	return new(T)
}

func Example_new_type() {
	// 必须要明确地传入 type argument
	fmt.Printf("%T\n", newT[int]())
	fmt.Printf("%T\n", newT[complex128])

	// Output:
	// *int
	// func() *complex128
}

// Constraint type inference
// It is used when one type parameter has a constraint defined in terms of another type parameter.
// When the type argument of one of those type parameters is known,
// the constraint is used to infer the type argument of the other.

func scaleInPlace[E constraints.Integer](s []E, factor E) {
	for i := range s {
		s[i] *= factor
	}
}

type ints []int

func Example_scaleInPlace() {
	s := ints{45, 7, 23, 89, 10} // function argument type inference
	scaleInPlace(s, 5)
	fmt.Println(s)
	// Output:
	// [225 35 115 445 50]
}

func scale[T ~[]E, E constraints.Integer](s T, factor E) T {
	r := make(T, len(s))
	for i := range s {
		r[i] = s[i] * factor
	}
	return r
}

func Example_constraint_type_inference() {
	fmt.Println(scale(ints{45, 7, 23, 89, 10}, 5)) // function argument type inference and constraint type inference
	// Output:
	// [225 35 115 445 50]
}
