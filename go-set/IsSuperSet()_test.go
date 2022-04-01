package go_set_test

import (
	"fmt"

	"github.com/scylladb/go-set"
)

var (
	s1 = set.NewByte8Set([8]byte{0}, [8]byte{1})
	s2 = set.NewByte8Set([8]byte{2, 3, 1, 0, 8, 9}, [8]byte{5, 3}, [8]byte{0}, [8]byte{1})

	s3 = set.NewStringSet("a", "b", "c", "d", "e")
)

func ExampleIsSuperSet() {
	fmt.Printf("s1.IsEqual(s2) == %t\n", s1.IsEqual(s2))
	fmt.Printf("s1.IsSubset(s2) == %t\n", s1.IsSubset(s2))
	fmt.Printf("s1.IsSuperset(s2) == %t\n", s1.IsSuperset(s2))
	// Output:
	// s1.IsEqual(s2) == false
	// s1.IsSubset(s2) == false
	// s1.IsSuperset(s2) == true
}

func ExampleString() {
	fmt.Println(s2.String())
	fmt.Println(s3.String())
	// Output:
	// [[5 3 0 0 0 0 0 0], [0 0 0 0 0 0 0 0], [1 0 0 0 0 0 0 0], [2 3 1 0 8 9 0 0]]
	// [a, b, c, d, e]
}
