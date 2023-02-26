package type_declarations_test

import "fmt"

// A type definition creates a new, distinct type.
// The new type is called a defined type.
// It is different from any other type, including the type it is created from.
func Example_type_definitions() {
	type myFloat64 float64

	type (
		myString   string
		stringList []myString
	)

	var f1 float64 = 23
	var f2 myFloat64 = f1 // cannot use f1 (variable of type float64) as myFloat64 value in variable declaration

	// Output:
}

// A defined type may have methods associated with it.
type os struct{}

func (os) halt()      {}
func (os) poweroff()  {}
func (os) reboot()    {}
func (os) suspend()   {}
func (os) hibernate() {}

func Example_new_defined_type_with_methods() {
	var linux os
	linux.reboot()
	linux.poweroff()

	// Output:
}

// A defined type does not inherit any methods bound to the given type,
// but the method set of an interface type or of elements of a composite type remains unchanged.
type table string

func (t table) tableName() string {
	return string(t)
}

type model table

func Example_no_inherit() {
	var t table = "user"
	fmt.Println(t.tableName())

	var m model = "error"
	fmt.Println(m.tableName()) // m.tableName undefined (type model has no field or method tableName)

	// Output:
}

// generic type
// A generic type may also have methods associated with it.

type list[T any] struct {
	next  *list[T]
	value T
}

func (l *list[T]) Len() int { return 0 }
