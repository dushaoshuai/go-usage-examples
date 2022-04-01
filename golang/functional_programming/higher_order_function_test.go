package functional_programming_test

import (
	"fmt"
)

// Package twice
// https://en.wikipedia.org/wiki/Higher-order_function#Go

// twice (higher-order function) takes a function, and applies the function to some value twice.
// If twice has to be applied several times for the same f it preferably should return a function rather than a value.
// This is in line with the "don't repeat yourself" principle.
func twice(f func(int) int) func(int) int {
	return func(x int) int {
		return f(f(x))
	}
}

func ExampleCurrying() {
	plusSix := twice(func(x int) int {
		return x + 3
	})
	fmt.Println(plusSix(7))
	// Output:
	// 13
}
