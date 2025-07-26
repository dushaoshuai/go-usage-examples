package slice_test

func Example_full_slice_expressions() {
	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmtSlice("a", a[:])

	s := a[:5:5] // 注意这里使用了 full slice expression，将切片 s 的容量限制为 5
	printlnln("s := a[:5:5]")
	fmtSlice("a", a[:])
	fmtSlice("s", s)

	// 向切片 s 中添加一个新的元素，虽然底层数组 a 可以容纳下，
	// 但是 s 的容量被限制为 5，因此会给 s 分配一个新的底层数组，
	// 因此 a 的元素并没有发生变化，可以用这个特性来保护数据不被修改
	s = append(s, 999)
	printlnln("s = append(s, 999)")
	fmtSlice("a", a[:])
	fmtSlice("s", s)

	// Output:
	// a: [0 1 2 3 4 5 6 7 8 9]
	// len(a): 10
	// cap(a): 10
	//
	// s := a[:5:5]
	//
	// a: [0 1 2 3 4 5 6 7 8 9]
	// len(a): 10
	// cap(a): 10
	//
	// s: [0 1 2 3 4]
	// len(s): 5
	// cap(s): 5
	//
	// s = append(s, 999)
	//
	// a: [0 1 2 3 4 5 6 7 8 9]
	// len(a): 10
	// cap(a): 10
	//
	// s: [0 1 2 3 4 999]
	// len(s): 6
	// cap(s): 10
}
