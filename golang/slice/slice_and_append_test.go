package slice_test

func Example_slice_and_append() {
	a := []int{0, 1, 2, 3, 4}
	fmtSlice("a", a)

	// 这种方法可以用来从切片中间删除元素，
	// 并且保持元素的相对位置，
	// 就是效率不怎么高吧
	a = append(a[:2], a[3:]...)
	fmtSlice("a", a)

	// Output:
	// a: [0 1 2 3 4]
	// len(a): 5
	// cap(a): 5
	//
	// a: [0 1 3 4]
	// len(a): 4
	// cap(a): 5
}
