package type_parameters_test

import "fmt"

type coo struct{}

func (c *coo) String() string {
	return "My type is coo !!!"
}

func foo[T fmt.Stringer](x T) {
	fmt.Println(x.String())
}

func Example_interface_as_constraints() {
	foo(&coo{})
	// Output:
	// My type is coo !!!
}
