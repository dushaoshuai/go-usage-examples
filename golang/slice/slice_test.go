package slice_test

import "fmt"

func Example_slice_slice() {
	s1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := s1[3:5]
	fmt.Println(s1, s2)
	fmt.Println(len(s2), cap(s2))
	s2[0] = 22
	s2[1] = 33
	fmt.Println(s1, s2)
	s2[0] = 3
	s2[1] = 4
	s2 = append(s2, 44, 55, 66, 77)
	fmt.Println(s1, s2)
	s2 = append(s2, 88, 99)
	fmt.Println(s1, s2)
	s2 = []int{15, 16}
	fmt.Println(s1, s2)
	// Output:
	// [0 1 2 3 4 5 6 7 8 9] [3 4]
	// 2 7
	// [0 1 2 22 33 5 6 7 8 9] [22 33]
	// [0 1 2 3 4 44 55 66 77 9] [3 4 44 55 66 77]
	// [0 1 2 3 4 44 55 66 77 9] [3 4 44 55 66 77 88 99]
	// [0 1 2 3 4 44 55 66 77 9] [15 16]
}
