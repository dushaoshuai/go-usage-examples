package utf8_test

import (
	"fmt"
	"unicode/utf8"
)

func ExampleDecodeRune_v1() {
	b := []byte("Hello,世界")

	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		fmt.Printf("%c %v\n", r, size)
		b = b[size:]
	}
	// Output:
	// H 1
	// e 1
	// l 1
	// l 1
	// o 1
	// , 1
	// 世 3
	// 界 3
}

func ExampleDecodeRune_v2() {
	b := []byte("Hello,世界")

	for i := 0; i < len(b); {
		r, size := utf8.DecodeRune(b[i:])
		fmt.Printf("%c %v\n", r, size)
		i += size
	}
	// Output:
	// H 1
	// e 1
	// l 1
	// l 1
	// o 1
	// , 1
	// 世 3
	// 界 3
}
