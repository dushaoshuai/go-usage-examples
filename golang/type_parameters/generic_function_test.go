package type_parameters_test

func Example_generic_function() {
	printStringSlice := printSlice[string]
	printStringSlice([]string{"a", "b", "c", "d"})
	// Output:
	// a
	// b
	// c
	// d
}
