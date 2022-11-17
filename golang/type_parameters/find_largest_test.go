package type_parameters

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// 参考 slices.Sort(),
// https://pkg.go.dev/golang.org/x/exp@v0.0.0-20220927162542-c76eaa363f9d/slices#Sort
func findSmallest[T constraints.Ordered](x []T) T {
	var smallest T
	if len(x) == 0 {
		return smallest
	}

	smallest = x[0]
	for i, d := range x {
		if i == 0 {
			continue
		}
		if d < smallest {
			smallest = d
		}
	}
	return smallest
}

func Example_findSmallest() {
	fmt.Println(findSmallest([]uint8{1, 56, 255, 67, 90, 34, 23}))
	fmt.Println(findSmallest([]string{"one"}))
	fmt.Println(findSmallest([]int{}))
	// Output:
	// 1
	// one
	// 0
}

// 参考 slices.SortFunc(),
// https://pkg.go.dev/golang.org/x/exp@v0.0.0-20220927162542-c76eaa363f9d/slices#SortFunc
func findSmallestFunc[T any](x []T, less func(T, T) bool) T {
	var smallest T
	if len(x) == 0 {
		return smallest
	}

	smallest = x[0]
	for i, d := range x {
		if i == 0 {
			continue
		}
		if less(d, smallest) {
			smallest = d
		}
	}
	return smallest
}

type orderedStruct struct {
	val int
}

func Example_findSmallestFunc() {
	fmt.Println(
		findSmallestFunc(
			[]int{-5, 0, 3, 24, -56},
			func(x, y int) bool { return x < y },
		),
	)
	fmt.Println(
		findSmallestFunc(
			[]orderedStruct{
				{0},
				{-5},
				{3},
				{24},
				{-56},
			},
			func(x, y orderedStruct) bool { return x.val < y.val },
		).val,
	)
	fmt.Println(
		findSmallestFunc(
			[]int{},
			func(x, y int) bool { return x < y },
		),
	)
	// Output:
	// -56
	// -56
	// 0
}
