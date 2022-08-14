package syntax_test

import (
	"fmt"

	"google.golang.org/protobuf/types/known/typepb"

	"api-examples/protobuf/well_known_types/syntax/protobuf"
)

func ExampleSyntax() {
	foo := &protobuf.Foo{
		TestSyntax: typepb.Syntax_SYNTAX_PROTO3,
	}
	fmt.Printf("%#v\n", foo)

	// Output:
	// &protobuf.Foo{state:impl.MessageState{NoUnkeyedLiterals:pragma.NoUnkeyedLiterals{}, DoNotCompare:pragma.DoNotCompare{}, DoNotCopy:pragma.DoNotCopy{}, atomicMessageInfo:(*impl.MessageInfo)(nil)}, sizeCache:0, unknownFields:[]uint8(nil), TestSyntax:1}
}
