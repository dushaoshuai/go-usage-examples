package conversion_test

import (
	"fmt"
	"unsafe"
)

// An explicit conversion is an expression of the form T(x)
// where T is a type and x is an expression that can be converted to type T.

// If the type starts with the operator * or <-,
// or if the type starts with the keyword func and has no result list,
// it must be parenthesized when necessary to avoid ambiguity.
func Example_conversions_must_be_parenthesized() {
	// (*T)(p)  p is converted to *T
	// see src/math/unsafe.go
	// print the IEEE 754 binary representation of a float64 value
	f := 6.5
	b := *(*uint64)(unsafe.Pointer(&f))
	fmt.Printf("the IEEE 754 binary representation of %f is %b\n", f, b)
	// get the float64 back from the IEEE 754 binary representation b
	f = *(*float64)(unsafe.Pointer(&b))
	fmt.Printf("the IEEE 754 binary representation %b is %f\n", b, f)

	// (<-chan int)(c)  c is converted to <-chan int
	c := make(chan int)
	rc := (<-chan int)(c)
	fmt.Printf("%T %T\n", c, rc)

	// (func())(x)      x is converted to func()
	// (func() int)(x)  x is converted to func() int
	var fu func()
	fn := (func())(fu)
	var fc func() int
	fc = (func() int)(fc)
	fmt.Printf("%T\n%T\n", fn, fc)

	// Output:
	// the IEEE 754 binary representation of 6.500000 is 100000000011010000000000000000000000000000000000000000000000000
	// the IEEE 754 binary representation 100000000011010000000000000000000000000000000000000000000000000 is 6.500000
	// chan int <-chan int
	// func()
	// func() int
}

// Finally, for historical reasons, an integer value may be converted to a string type.
// This form of conversion yields a string containing the (possibly multi-byte) UTF-8 representation of the Unicode code point with the given integer value.
// Values outside the range of valid Unicode code points are converted to "\uFFFD".
//
// string('a')          // "a"
// string(65)           // "A"
// string('\xf8')       // "\u00f8" == "ø" == "\xc3\xb8"
// string(-1)           // "\ufffd" == "\xef\xbf\xbd"
//
// type myString string
// myString('\u65e5')   // "\u65e5" == "日" == "\xe6\x97\xa5"
//
// Note: This form of conversion may eventually be removed from the language.
// The go vet tool flags certain integer-to-string conversions as potential errors.
// Library functions such as utf8.AppendRune or utf8.EncodeRune should be used instead.
func Example_integer_to_string() {
	// conversion from untyped int to string yields a string of one rune,
	// not a string of digits (did you mean fmt.Sprint(x)?)
	// fmt.Printf("string(-1) ==> %q\n", string(-1))
	// fmt.Printf("string(0xf8) ==> %q\n", string(0xf8))

	fmt.Printf("string('a') ==> %q\n", string('a'))

	// Output:
	// string('a') ==> "a"
}

// Converting a slice of bytes to a string type yields a string
// whose successive bytes are the elements of the slice.
func Example_convert_a_slice_of_bytes_to_a_string_type() {
	fmt.Printf(
		"string([]byte{34, 45, 67, 78, 100, 125, 255}) => %s\n",
		string([]byte{34, 45, 67, 78, 100, 125, 255}),
	)

	fmt.Printf(
		"string([]byte{'\\xf0', '\\x9f', '\\x8c', '\\x8d'}) => %s\n",
		string([]byte{'\xf0', '\x9f', '\x8c', '\x8d'}),
	)

	// Output:
	// string([]byte{34, 45, 67, 78, 100, 125, 255}) => "-CNd}�
	// string([]byte{'\xf0', '\x9f', '\x8c', '\x8d'}) => 🌍
}
