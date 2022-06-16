package json_test

import (
	"encoding/json"
	"fmt"
)

type desc []string

func Example_marshal_array() {
	var a desc = []string{"a", "b", "c"}
	b, _ := json.Marshal(a)
	fmt.Printf("%s\n", b)

	// Output:
	// ["a","b","c"]
}
