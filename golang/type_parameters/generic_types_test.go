package type_parameters_test

// vector is a name for a slice of any element type.
type vector[T any] []T

// v is a vector of int values.
var v vector[int]

// Push adds a value to the end of a vector.
func (v *vector[T]) Push(x T) {
	*v = append(*v, x)
}
