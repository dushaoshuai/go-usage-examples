package json_test

import (
	"encoding/json"
	"fmt"
)

type trival struct {
	A int
	B string
	C float64
}

func Example_unmarshal_json_null() {
	var d1 trival
	err := json.Unmarshal([]byte("null"), &d1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", d1)

	d2 := &trival{}
	err = json.Unmarshal([]byte("null"), d2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", d2)

	// Output:
	// {A:0 B: C:0}
}
