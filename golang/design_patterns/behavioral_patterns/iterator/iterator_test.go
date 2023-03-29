package iterator_test

import (
	"fmt"

	"github.com/dushaoshuai/go-usage-examples/golang/design_patterns/behavioral_patterns/iterator/iterator"
)

func Example_iterator() {
	visit := func(e *iterator.Elem) {
		fmt.Println(e.Name)
	}

	c := iterator.NewCollection(
		&iterator.Elem{Name: "Go"},
		&iterator.Elem{Name: "C"},
		&iterator.Elem{Name: "Java"},
	)

	f := c.ForwardIterator()
	for f.HasNext() {
		visit(f.GetNext())
	}

	r := c.ReverseIterator()
	for r.HasNext() {
		visit(r.GetNext())
	}

	// Output:
	// Go
	// C
	// Java
	// Java
	// C
	// Go
}
