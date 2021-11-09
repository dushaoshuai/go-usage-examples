package goto_statement

import "fmt"

// Goto demonstrates that a labeled statement has its own syntactic block.
// The variable "a" was declared several times.
func Goto() {
	n := 0

declarationA:
	a := 1
	fmt.Println(a)
	n++
	if n < 10 {
		goto declarationA
	}
}
