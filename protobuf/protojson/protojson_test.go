package protojson_test

import (
	"fmt"
	"reflect"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"

	protobuf "google.golang.org/protobuf/proto"

	"api-examples/protobuf/protojson/proto"
)

func Example_protojson_Marshal_and_Unmarshal() {
	bars := []protobuf.Message{
		&proto.Bar{
			TestOneof: &proto.Bar_Name{
				Name: "Go",
			},
		},
		&proto.Bar{
			TestOneof: &proto.Bar_Empty{
				Empty: nil,
			},
		},
		&proto.Bar{
			TestOneof: &proto.Bar_Null{
				Null: structpb.NullValue_NULL_VALUE,
			},
		},
	}

	var (
		marshaler protojson.MarshalOptions
		encodings = make([][]byte, 3)

		unmarshaler protojson.UnmarshalOptions
		decodings   = make([]protobuf.Message, 3)
	)

	// test marshal/encoding
	fmt.Println("encoding ...")
	for i, bar := range bars {
		encoding, err := marshaler.Marshal(bar)
		if err != nil {
			fmt.Println(i, err)
			continue
		}
		fmt.Println(i, encoding, string(encoding))
		encodings[i] = encoding
	}

	// test unmarshal/decoding
	for i := range decodings {
		decodings[i] = &proto.Bar{}
	}
	fmt.Println("decoding ...")
	for i, encodedBar := range encodings {
		err := unmarshaler.Unmarshal(encodedBar, decodings[i])
		if err != nil {
			fmt.Println(i, err)
			continue
		}

		if !reflect.DeepEqual(bars[i], decodings[i]) {
			fmt.Println(i, "decoding and original bar is not deep equal")
		}

		fmt.Printf("%d %#v\n", i, decodings[i].(*proto.Bar).GetTestOneof())

		bar := decodings[i].(*proto.Bar)
		switch bar.GetTestOneof().(type) {
		case *proto.Bar_Name:
			fmt.Println("name is", bar.GetName())
		case *proto.Bar_Empty:
			fmt.Println("empty", bar.GetEmpty())
		case *proto.Bar_Null:
			fmt.Println("null", bar.GetNull())
		}
	}

	// Output:
	// encoding ...
	// 0 [123 34 110 97 109 101 34 58 34 71 111 34 125] {"name":"Go"}
	// 1 [123 34 101 109 112 116 121 34 58 123 125 125] {"empty":{}}
	// 2 [123 34 110 117 108 108 34 58 110 117 108 108 125] {"null":null}
	// decoding ...
	// 0 &proto.Bar_Name{Name:"Go"}
	// name is Go
	// 1 decoding and original bar is not deep equal
	// 1 &proto.Bar_Empty{Empty:(*emptypb.Empty)(0xc0000a2de0)}
	// empty
	// 2 &proto.Bar_Null{Null:0}
	// null NULL_VALUE
}
