package type_declarations_test

import (
	"fmt"
	"strings"
)

// Within the scope of stringBuilder, it serves as an alias for strings.Builder.
// stringBuilder and strings.Builder are identical types.
type stringBuilder = strings.Builder

func Example_alias_declarations() {
	var b stringBuilder
	for i := 3; i >= 1; i-- {
		fmt.Fprintf(&b, "%d...", i)
	}
	b.WriteString("ignition")
	fmt.Println(b.String())

	// Output:
	// 3...2...1...ignition
}
