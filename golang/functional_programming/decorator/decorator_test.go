package decorator_test

import (
	"fmt"

	"github.com/dushaoshuai/go-usage-examples/golang/functional_programming/decorator/decorator"
)

func Example_decorator() {
	h := decorator.Decorate(
		decorator.Go,
		decorator.WithPython,
		decorator.WithC,
	)
	fmt.Println(h())

	// Output:
	// Python C Go C Python
}
