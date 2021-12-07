package embed_test

import (
	"embed"
	_ "embed"
	"fmt"
	"log"
)

//go:embed hello.txt
var a string

func Example_embed_to_string() {
	fmt.Println(a)
	// Output:
	// Hello World!
	// Hello world!
	// hello world!
}

//go:embed hello.txt
var b []byte

func Example_embed_to_bytes() {
	fmt.Println(string(b))
	// Output:
	// Hello World!
	// Hello world!
	// hello world!
}

//go:embed hello.txt
var f embed.FS

func Example_embed_to_FS() {
	data, err := f.ReadFile("hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
	// Output:
	// Hello World!
	// Hello world!
	// hello world!
}
