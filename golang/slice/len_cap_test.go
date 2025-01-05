package slice_test

func Example_slice_literal_len_cap() {
	a := []int{}
	b := []int{1}
	c := []int{1, 2}
	d := []int{1, 2, 3}
	fmtSlice("a", a)
	fmtSlice("b", b)
	fmtSlice("c", c)
	fmtSlice("d", d)

	// Output:
	// a: []
	// len(a): 0
	// cap(a): 0
	//
	// b: [1]
	// len(b): 1
	// cap(b): 1
	//
	// c: [1 2]
	// len(c): 2
	// cap(c): 2
	//
	// d: [1 2 3]
	// len(d): 3
	// cap(d): 3
}

func Example_make_slice_len_cap() {
	a := make([]int, 0)
	b := make([]int, 1)
	c := make([]int, 2)
	d := make([]int, 3)
	fmtSlice("a", a)
	fmtSlice("b", b)
	fmtSlice("c", c)
	fmtSlice("d", d)

	e := make([]int, 0, 4)
	f := make([]int, 1, 4)
	g := make([]int, 2, 2)
	h := make([]int, 3, 4)
	fmtSlice("e", e)
	fmtSlice("f", f)
	fmtSlice("g", g)
	fmtSlice("h", h)

	// Output:
	// a: []
	// len(a): 0
	// cap(a): 0
	//
	// b: [0]
	// len(b): 1
	// cap(b): 1
	//
	// c: [0 0]
	// len(c): 2
	// cap(c): 2
	//
	// d: [0 0 0]
	// len(d): 3
	// cap(d): 3
	//
	// e: []
	// len(e): 0
	// cap(e): 4
	//
	// f: [0]
	// len(f): 1
	// cap(f): 4
	//
	// g: [0 0]
	// len(g): 2
	// cap(g): 2
	//
	// h: [0 0 0]
	// len(h): 3
	// cap(h): 4
}

func Example_slice_slice_len_cap() {
	a := make([]int, 7)
	b := a[:4]   // [simple slice expressions](https://go.dev/ref/spec#:~:text=on%20the%20capacity.-,Simple%20slice%20expressions,-The%20primary%20expression)
	c := a[:4:4] // [Full slice expressions](https://go.dev/ref/spec#:~:text=0%5D%20%20%20%20//%20s3%20%3D%3D%20nil-,Full%20slice%20expressions,-The%20primary%20expression)

	fmtSlice("a", a)
	fmtSlice("b", b)
	fmtSlice("c", c)

	// Output:
	// a: [0 0 0 0 0 0 0]
	// len(a): 7
	// cap(a): 7
	//
	// b: [0 0 0 0]
	// len(b): 4
	// cap(b): 7
	//
	// c: [0 0 0 0]
	// len(c): 4
	// cap(c): 4
}
