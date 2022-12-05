package strconv_test

import (
	"fmt"
	"strconv"
)

func ExampleQuote() {
	s := `"The next planet was inhabited by a tippler.	☺&"`
	fmt.Println(strconv.Quote(s))
	fmt.Println(strconv.QuoteToASCII(s))

	// Output:
	// "\"The next planet was inhabited by a tippler.\t☺&\""
	// "\"The next planet was inhabited by a tippler.\t\u263a&\""
}

func ExampleQuoteRune() {
	r := '☺'
	fmt.Println(strconv.QuoteRune(r))
	fmt.Println(strconv.QuoteRuneToASCII(r))

	// Output:
	// '☺'
	// '\u263a'
}

func ExampleUnquote() {
	// Unquote unquotes single-quoted, double-quoted, ok backquoted Go string literal.
	// Output:
}
