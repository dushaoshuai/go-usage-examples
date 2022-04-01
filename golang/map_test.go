package golang_test

import (
	"fmt"
	"math"
)

func Example_map_size_MaxInt64() {
	// 大小为 math.MaxInt64 的 map 也没有感觉慢
	m := make(map[string]string, math.MaxInt64)
	m["one"] = "1"
	fmt.Printf("%#v\n", m["one"])
	// Output:
	// "1"
}

func Example_len_map() {
	m := make(map[string]struct{}, math.MaxInt64)
	fmt.Println(len(m))
	// Output:
	// 0
}
