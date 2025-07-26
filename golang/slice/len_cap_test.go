package slice_test

func Example_slice_literal_len_cap() {
	a := []int{}
	printlnln("a := []int{}")
	fmtSlice("a", a)

	b := []int{1}
	printlnln("b := []int{1}")
	fmtSlice("b", b)

	c := []int{1, 2}
	printlnln("c := []int{1, 2}")
	fmtSlice("c", c)

	d := []int{1, 2, 3}
	printlnln("d := []int{1, 2, 3}")
	fmtSlice("d", d)

	// Output:
	// a := []int{}
	//
	// a: []
	// len(a): 0
	// cap(a): 0
	//
	// b := []int{1}
	//
	// b: [1]
	// len(b): 1
	// cap(b): 1
	//
	// c := []int{1, 2}
	//
	// c: [1 2]
	// len(c): 2
	// cap(c): 2
	//
	// d := []int{1, 2, 3}
	//
	// d: [1 2 3]
	// len(d): 3
	// cap(d): 3
}

func Example_make_slice_len_cap() {
	a := make([]int, 0)
	printlnln("a := make([]int, 0)")
	fmtSlice("a", a)

	b := make([]int, 1)
	printlnln("b := make([]int, 1)")
	fmtSlice("b", b)

	c := make([]int, 2)
	printlnln("c := make([]int, 2)")
	fmtSlice("c", c)

	d := make([]int, 3)
	printlnln("d := make([]int, 3)")
	fmtSlice("d", d)

	e := make([]int, 0, 4)
	printlnln("e := make([]int, 0, 4)")
	fmtSlice("e", e)

	f := make([]int, 1, 4)
	printlnln("f := make([]int, 1, 4)")
	fmtSlice("f", f)

	g := make([]int, 2, 2)
	printlnln("g := make([]int, 2, 2)")
	fmtSlice("g", g)

	h := make([]int, 3, 4)
	printlnln("h := make([]int, 3, 4)")
	fmtSlice("h", h)

	// Output:
	//
	// a := make([]int, 0)
	//
	// a: []
	// len(a): 0
	// cap(a): 0
	//
	// b := make([]int, 1)
	//
	// b: [0]
	// len(b): 1
	// cap(b): 1
	//
	// c := make([]int, 2)
	//
	// c: [0 0]
	// len(c): 2
	// cap(c): 2
	//
	// d := make([]int, 3)
	//
	// d: [0 0 0]
	// len(d): 3
	// cap(d): 3
	//
	// e := make([]int, 0, 4)
	//
	// e: []
	// len(e): 0
	// cap(e): 4
	//
	// f := make([]int, 1, 4)
	//
	// f: [0]
	// len(f): 1
	// cap(f): 4
	//
	// g := make([]int, 2, 2)
	//
	// g: [0 0]
	// len(g): 2
	// cap(g): 2
	//
	// h := make([]int, 3, 4)
	//
	// h: [0 0 0]
	// len(h): 3
	// cap(h): 4
}

func Example_slice_slice_len_cap() {
	a := make([]int, 7)
	printlnln("a := make([]int, 7)")
	fmtSlice("a", a)

	b := a[:4] // [simple slice expressions](https://go.dev/ref/spec#:~:text=on%20the%20capacity.-,Simple%20slice%20expressions,-The%20primary%20expression)
	printlnln("b := a[:4]")
	fmtSlice("b", b)

	c := a[:4:4] // [Full slice expressions](https://go.dev/ref/spec#:~:text=0%5D%20%20%20%20//%20s3%20%3D%3D%20nil-,Full%20slice%20expressions,-The%20primary%20expression)
	printlnln("c := a[:4:4]")
	fmtSlice("c", c)

	// Output:
	//
	// a := make([]int, 7)
	//
	// a: [0 0 0 0 0 0 0]
	// len(a): 7
	// cap(a): 7
	//
	// b := a[:4]
	//
	// b: [0 0 0 0]
	// len(b): 4
	// cap(b): 7
	//
	// c := a[:4:4]
	//
	// c: [0 0 0 0]
	// len(c): 4
	// cap(c): 4
}
