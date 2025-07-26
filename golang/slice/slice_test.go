package slice_test

import "fmt"

func Example_slice_slice() {
	s1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	printlnln("s1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}")
	fmtSlice("s1", s1)

	s2 := s1[3:5]
	printlnln("s2 := s1[3:5]")
	fmtSlice("s2", s2)

	s2[0] = 22
	s2[1] = 33
	fmt.Println("s2[0] = 22")
	printlnln("s2[1] = 33")
	fmtSlice("s1", s1)
	fmtSlice("s2", s2)

	s2[0] = 3
	s2[1] = 4
	s2 = append(s2, 44, 55, 66, 77)
	fmt.Println("s2[0] = 3")
	fmt.Println("s2[1] = 4")
	printlnln("s2 = append(s2, 44, 55, 66, 77)")
	fmtSlice("s1", s1)
	fmtSlice("s2", s2)

	s2 = append(s2, 88, 99)
	printlnln("s2 = append(s2, 88, 99)")
	fmtSlice("s1", s1)
	fmtSlice("s2", s2)

	s2 = []int{15, 16}
	printlnln("s2 = []int{15, 16}")
	fmtSlice("s1", s1)
	fmtSlice("s2", s2)

	// Output:
	// s1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//
	// s1: [0 1 2 3 4 5 6 7 8 9]
	// len(s1): 10
	// cap(s1): 10
	//
	// s2 := s1[3:5]
	//
	// s2: [3 4]
	// len(s2): 2
	// cap(s2): 7
	//
	// s2[0] = 22
	// s2[1] = 33
	//
	// s1: [0 1 2 22 33 5 6 7 8 9]
	// len(s1): 10
	// cap(s1): 10
	//
	// s2: [22 33]
	// len(s2): 2
	// cap(s2): 7
	//
	// s2[0] = 3
	// s2[1] = 4
	// s2 = append(s2, 44, 55, 66, 77)
	//
	// s1: [0 1 2 3 4 44 55 66 77 9]
	// len(s1): 10
	// cap(s1): 10
	//
	// s2: [3 4 44 55 66 77]
	// len(s2): 6
	// cap(s2): 7
	//
	// s2 = append(s2, 88, 99)
	//
	// s1: [0 1 2 3 4 44 55 66 77 9]
	// len(s1): 10
	// cap(s1): 10
	//
	// s2: [3 4 44 55 66 77 88 99]
	// len(s2): 8
	// cap(s2): 14
	//
	// s2 = []int{15, 16}
	//
	// s1: [0 1 2 3 4 44 55 66 77 9]
	// len(s1): 10
	// cap(s1): 10
	//
	// s2: [15 16]
	// len(s2): 2
	// cap(s2): 2
}
