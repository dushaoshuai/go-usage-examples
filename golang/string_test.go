package golang_test

import (
	"fmt"
	"reflect"
	"unicode/utf8"
)

// for range 会进行 utf8 解析
func Example_range_string() {
	s := "Hello, gophers! Hello, 世界!"
	var i int
	for range s {
		i++
	}
	fmt.Println(i, len(s), utf8.RuneCountInString(s))
	// Output:
	// 26 30 26
}

// len()，索引字符串，substring
// 这些对字符串的基本操作都是以 byte 为单位进行的
func Example_len_index_substring_string() {
	s := "Hello, gophers! Hello, 世界!"
	fmt.Println(len(s), utf8.RuneCountInString(s))
	fmt.Println(reflect.TypeOf(s[6]))
	s = s[:28]
	fmt.Println(s)
	// Output:
	// 30 26
	// uint8
	// Hello, gophers! Hello, 世��
}
