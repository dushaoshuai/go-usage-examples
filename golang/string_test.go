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

// 字符串的几种表示方法
func Example_string_representation() {
	// 一个汉字 3 字节
	a := "世界"
	b := "\xe4\xb8\x96\xe7\x95\x8c"
	c := "0xe40xb80x960xe70x950x8c" // 这样不行
	d := "\u4e16\u754c"
	e := "\U00004e16\U0000754c"
	fmt.Println(a, b, c, d, e)

	// Output:
	// 世界 世界 0xe40xb80x960xe70x950x8c 世界 世界
}
