package golang_test

func Example_new_a_map() {
	m := new(map[string]int)
	(*m)["a"] = 'a'

	// $ go test -run Example_new_a_map
	// --- FAIL: Example_new_a_map (0.00s)
	// panic: assignment to entry in nil map [recovered]
	//        panic: assignment to entry in nil map
}
