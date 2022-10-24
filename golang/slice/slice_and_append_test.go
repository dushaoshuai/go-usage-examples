package slice

func Example_slice_and_append() {
	a := []int{0, 1, 2, 3, 4}
	fmtSlice(a)

	// 这种方法可以用来从切片中间删除元素，
	// 并且保持元素的相对位置，
	// 就是效率不怎么高吧
	a = append(a[:2], a[3:]...)
	fmtSlice(a)

	// Output:
	// slice: [0 1 2 3 4]
	// len(slice): 5
	// cap(slice): 5
	// slice: [0 1 3 4]
	// len(slice): 4
	// cap(slice): 5
}
