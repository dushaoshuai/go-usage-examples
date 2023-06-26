package golang_test

import "fmt"

// https://go.dev/ref/spec#Address_operators
func Example_address_op() {
	// a variable
	a := 10
	pa := &a
	fmt.Println(*pa)

	// a pointer indirection
	b := &*pa
	fmt.Println(*b)

	// a slice indexing operation
	sli := []int{0, 1, 2}
	pSliElem := &sli[1]
	fmt.Println(*pSliElem)

	// a field selector of an addressable struct operand
	struc := struct {
		a int
		b string
		c float64
	}{c: 6.6}
	pStructField := &struc.c
	fmt.Println(*pStructField)

	// an array indexing operation of an addressable array
	array := [3]int{0, 1, 2}
	pArrayElem := &array[1]
	fmt.Println(*pArrayElem)

	// a (possibly parenthesized) composite literal
	pStruct := &struct {
		a int
	}{a: 10}
	fmt.Println(*pStruct)
	pArray := &[3]int{0, 1, 2}
	fmt.Println(*pArray)
	pArray = &[...]int{0, 1, 2}
	fmt.Println(*pArray)
	pSlice := &[]int{0, 1, 2}
	fmt.Println(*pSlice)
	pMap := &map[int]string{10: "a"}
	fmt.Println(*pMap)

	// Output:
	// 10
	// 10
	// 1
	// 6.6
	// 1
	// {10}
	// [0 1 2]
	// [0 1 2]
	// [0 1 2]
	// map[10:a]
}
