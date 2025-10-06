package link_X

import (
	"fmt"
)

//go:generate go test -ldflags="-X 'github.com/dushaoshuai/go-usage-examples/golang/cmdgo/link-X.a=A' -X 'github.com/dushaoshuai/go-usage-examples/golang/cmdgo/link-X.b=B'" -test.run ^Example_link_X$
func Example_link_X() {
	fmt.Printf("a == %v\n", a)
	fmt.Printf("b == %v\n", b)

	// Output:
	// a == A
	// b == b
}
