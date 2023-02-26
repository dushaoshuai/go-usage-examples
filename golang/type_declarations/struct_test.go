package type_declarations_test

import "fmt"

// A defined type may have methods associated with it.
// It does not inherit any methods bound to the given type,
// but the method set of an interface type or of elements of a composite type remains unchanged.

type aoo struct{}

func (a aoo) name() {
	fmt.Println("aoo")
}

type bar struct {
	aoo
}

type foo bar

func Example_struct_struct() {
	var x foo
	x.name()
	x.aoo.name()

	// Output:
	// aoo
	// aoo
}
