package comparison_test

import (
	"errors"
	"fmt"
)

func Example_interface_comparisons() {
	// Interface values are comparable. Two interface values are equal if they have
	// identical dynamic types and equal dynamic values or if both have value nil.
	var ia interface{}
	var ib error
	printInfo("ia", ia)
	printInfo("ib", ib)
	fmt.Printf("ia == ib ?=> %v\n", ia == ib)
	fmt.Println()
	err := errors.New("example error")
	ia = err
	ib = err
	printInfo("ia", ia)
	printInfo("ib", ib)
	fmt.Printf("ia == ib ?=> %v\n", ia == ib)
	fmt.Println()

	// A value x of non-interface type X and a value t of interface type T
	// are comparable when values of type X are comparable and X implements T.
	// They are equal if t's dynamic type is identical to X and t's dynamic value is equal to x.
	var a int
	ia = a
	printInfo("a", a)
	printInfo("ia", ia)
	fmt.Printf("a == ia ?=> %v\n", a == ia)

	// Output:
	// ia <nil> <nil>
	// ib <nil> <nil>
	// ia == ib ?=> true
	//
	// ia *errors.errorString example error
	// ib *errors.errorString example error
	// ia == ib ?=> true
	//
	// a int 0
	// ia int 0
	// a == ia ?=> true
}

func printInfo(name string, variable any) {
	fmt.Printf("%s %T %[2]v\n", name, variable)
}
