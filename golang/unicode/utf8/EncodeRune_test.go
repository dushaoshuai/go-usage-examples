package utf8_test

import (
	"fmt"
	"unicode/utf8"
)

func ExampleEncodeRune() {
	r := '世'
	buf := make([]byte, 3)

	n := utf8.EncodeRune(buf, r)

	fmt.Println(buf)
	fmt.Println(n)
	// Output:
	// [228 184 150]
	// 3
}

func ExampleEncodeRune_outOfRange() {
	runes := []rune{
		// Less than 0, out of range.
		-1,
		// Greater than utf8.MaxRune, out of range.
		0x110000,
		// The Unicode replacement character.
		utf8.RuneError,
	}
	for _, c := range runes {
		buf := make([]byte, 3)
		size := utf8.EncodeRune(buf, c)
		fmt.Printf("%d %[1]s %d\n", buf, size)
	}
	// Output:
	// [239 191 189] � 3
	// [239 191 189] � 3
	// [239 191 189] � 3
}
