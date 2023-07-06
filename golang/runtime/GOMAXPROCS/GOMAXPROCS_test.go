package main_test

import (
	"fmt"
	"runtime"
)

func ExampleGOMAXPROCS() {
	// defaults to the number of logical CPUs
	fmt.Println(runtime.GOMAXPROCS(0))
	fmt.Println(runtime.NumCPU())

	// Output:
	// 16
	// 16
}
