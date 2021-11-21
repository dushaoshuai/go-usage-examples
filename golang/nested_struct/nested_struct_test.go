package nested_struct_test

import "fmt"

type point struct {
	x int
	y int
}

type colorPoint struct {
	point
	color uint64
}

func ExampleAssignment() {
	cp := colorPoint{
		point: point{
			x: 12,
			y: 34,
		},
		color: 56,
	}
	fmt.Println(cp)
	// Output:
	// {{12 34} 56}
}
