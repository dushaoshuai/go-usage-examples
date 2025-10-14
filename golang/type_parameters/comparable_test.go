package type_parameters

// comparable is an interface that is implemented by all comparable types (booleans, numbers, strings, pointers, channels, arrays of comparable types, structs whose fields are all comparable types).
// The comparable interface may only be used as a type parameter constraint, not as the type of a variable.

import (
	"fmt"
	"sort"
)

func MapKeys[K comparable, V any](m map[K]V) []K {
	var s []K
	for k := range m {
		s = append(s, k)
	}
	return s
}

func Example_t_Comparable() {
	keys := MapKeys(map[string]int{
		"a": 10,
		"b": 11,
		"c": 12,
	})
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	fmt.Println(keys)
	// Output:
	// [a b c]
}
