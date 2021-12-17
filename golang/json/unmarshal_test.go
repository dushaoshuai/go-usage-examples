package json_test

import (
	"encoding/json"
	"fmt"
	"log"
)

type foo struct {
	F1 int
	F2 string
}

func (f *foo) unmarshal1(j []byte) {
	err := json.Unmarshal(j, f) // 主要是为了对比两种用法
	if err != nil {
		log.Fatal(err)
	}
}

func (f *foo) unmarshal2(j []byte) {
	err := json.Unmarshal(j, &f) // 主要是为了对比这两种用法，没想到这种还真可以
	if err != nil {
		log.Fatal(err)
	}
}

func Example_unmarshal_and_pointer() {
	j := []byte(`{"F1": 2, "F2": "gopher"}`)
	f := &foo{}

	f.unmarshal1(j)
	fmt.Println(f)
	f.unmarshal2(j)
	fmt.Println(f)
	// Output:
	// &{2 gopher}
	// &{2 gopher}
}
