package any_test

import (
	"fmt"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/emptypb"

	"api-examples/protobuf/well_known_types/any/protobuf"
)

func ExampleAny() {
	wellKnownAny, err := anypb.New(new(emptypb.Empty))
	if err != nil {
		panic(err)
	}

	foo := &protobuf.Foo{
		TestAny: wellKnownAny,
	}
	encodedFoo, err := protojson.Marshal(foo)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", encodedFoo)

	// Output:
	// {"testAny":{"@type":"type.googleapis.com/google.protobuf.Empty", "value":{}}}
}
