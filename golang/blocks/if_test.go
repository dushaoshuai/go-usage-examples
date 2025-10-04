package blocks_test

import (
	"fmt"
)

// https://go.dev/ref/spec#If_statements
// https://go.dev/ref/spec#Blocks
// https://go.dev/ref/spec#Declarations_and_scope
// An identifier declared in a block may be redeclared in an inner block.
// While the identifier of the inner declaration is in scope, it denotes the entity declared by the inner declaration.
func Example_if() {
	f1 := func() int { return 10 }
	f2 := func() int { return 20 }
	y := 5
	z := 5

	if x := f1(); x < y {
		fmt.Println(y)
	} else if x := f2(); x > z {
		fmt.Println(x)
	} else {
		fmt.Println(z)
	}

	// Output:
	// 20
}
