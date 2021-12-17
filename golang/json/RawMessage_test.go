package json_test

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// This example uses RawMessage to use a precomputed JSON during marshal.
// https://pkg.go.dev/encoding/json@go1.17.5#RawMessage
func ExampleRawMessage_marshal() {
	h := json.RawMessage(`{"precomputed": true}`)

	c := struct {
		Header *json.RawMessage `json:"header"`
		Body   string           `json:"body"`
	}{
		Header: &h,
		Body:   "Hello Gophers!",
	}

	b, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	_, err = os.Stdout.Write(b)
	if err != nil {
		fmt.Println("error:", err)
	}
	// Output:
	// {
	// 	"header": {
	// 	"precomputed": true
	// },
	// 	"body": "Hello Gophers!"
	// }
}

// This example uses RawMessage to delay parsing part of a JSON message.
// https://pkg.go.dev/encoding/json@go1.17.5#example-RawMessage-Unmarshal
func ExampleRawMessage_unmarshal() {
	type Color struct {
		Space string
		Point json.RawMessage // delay parsing until we know the color space
	}
	type RGB struct {
		R uint8
		G uint8
		B uint8
	}
	type YCbCr struct {
		Y  uint8
		Cb int8
		Cr int8
	}

	var j = []byte(`[
{"Space": "YCbCr", "Point": {"Y": 255, "Cb": 0, "Cr": -10}},
{"Space": "RGB", "Point": {"R": 98, "G": 218, "B": 255}}
]`)
	var colors []Color
	err := json.Unmarshal(j, &colors)
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range colors {
		var dst interface{}
		switch c.Space {
		case "RGB":
			dst = new(RGB)
		case "YCbCr":
			dst = new(YCbCr)
		}
		err = json.Unmarshal(c.Point, dst)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(c.Space, dst)
	}
	// Output:
	// YCbCr &{255 0 -10}
	// RGB &{98 218 255}
}
