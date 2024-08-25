package array_test

import (
	"fmt"
)

func Example_zero_length() {
	var za [0]int
	fmt.Println(len(za), za)

	var zb [0]func()
	fmt.Println(len(zb), zb)

	var zc [0]int
	fmt.Println(za == zc)

	// invalid operation: zb == zd ([0]func() cannot be compared)
	// var zd [0]func()
	// fmt.Println(zb == zd)

	// Output:
	// 0 []
	// 0 []
	// true
}

// slog.Value 使用
// _ [0]func()
// 防止直接比较
type slogValue struct {
	_   [0]func() // disallow ==
	num uint64
	any any
}

type comparableSlogValue struct {
	num uint64
	any any
}

func Example_comp_SlogValue() {
	// invalid operation: slogValue{} == slogValue{} (struct containing [0]func() cannot be compared)
	// fmt.Println(slogValue{} == slogValue{})

	fmt.Println(comparableSlogValue{} == comparableSlogValue{})

	// Output:
	// true
}
