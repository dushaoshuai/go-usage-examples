package slice_test

import (
	"fmt"
)

// `b = append(a, T)`
// a 变量未被更新，含有的元素，长度，容量都不会改变，
// 可以理解为 b 就是 `a = append(a, T)` 里的 a 变量

func Example_append() {
	a := make([]int, 5, 6)
	for i := range a {
		a[i] = i
	}
	printInfo(a)

	// 注意这里，切片 a 的长度为 5，容量为 6，还可以再容纳一个元素，
	// 因此底层数组并未改变，不过 a 变量未被更新，
	// 这里可以理解为 b 变量就是 `a = append(a, 5)` 里的更新过的 a 变量
	b := append(a, 5)
	printInfo(a)
	printInfo(b)
	// Output:
	// [0 1 2 3 4] 5 6
	// [0 1 2 3 4] 5 6
	// [0 1 2 3 4 5] 6 6
}

func printInfo(s []int) {
	fmt.Println(s, len(s), cap(s))
}
