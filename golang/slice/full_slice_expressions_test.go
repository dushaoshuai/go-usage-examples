package slice_test

import "fmt"

func Example_full_slice_expressions() {
	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := a[:5:5] // 注意这里使用了 full slice expression，将切片 s 的容量限制为 5
	fmt.Println("a =>", a)
	fmt.Println("s =>", s)

	// 向切片 s 中添加一个新的元素，虽然底层数组 a 可以容纳下，
	// 但是 s 的容量被限制为 5，因此会给 s 分配一个新的底层数组，
	// 因此 a 的元素并没有发生变化，可以用这个特性来保护数据不被修改
	s = append(s, 999)
	fmt.Println("a =>", a)
	fmt.Println("s =>", s)
	// Output:
	// a => [0 1 2 3 4 5 6 7 8 9]
	// s => [0 1 2 3 4]
	// a => [0 1 2 3 4 5 6 7 8 9]
	// s => [0 1 2 3 4 999]
}
