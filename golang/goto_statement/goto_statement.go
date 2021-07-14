package goto_statement

import "fmt"

// This function demonstrates that a labeled statement has it's
// own syntactic block. The variable "a" was declarated several times.
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
