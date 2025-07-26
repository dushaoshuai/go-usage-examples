package slice_test

func Example_zero_length() {
	zs := make([]int, 0)
	printlnln("zs := make([]int, 0)")
	fmtSlice("zs", zs)

	zs = append(zs, 18)
	printlnln("zs = append(zs, 18)")
	fmtSlice("zs", zs)

	// Output:
	// zs := make([]int, 0)
	//
	// zs: []
	// len(zs): 0
	// cap(zs): 0
	//
	// zs = append(zs, 18)
	//
	// zs: [18]
	// len(zs): 1
	// cap(zs): 1
}
