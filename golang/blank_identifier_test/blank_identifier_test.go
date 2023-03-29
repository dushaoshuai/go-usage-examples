package blank_identifier_test // The PackageName must not be the blank identifier.

import (
	"fmt"
	// in import declaration
	// import a package solely for its side-effects (initialization)
	_ "image/png"
)

// a constant declaration
const _ = 10

// type declaration and type parameter declaration
type _[T comparable, _ any] struct {
	a T
}

// a variable declaration
var _ = 10

// a function declaration
func _() string {
	a := "Go"

_: // a labeled statement
	a += "o"

	return a
}

func Example_ignore_right_hand_side_values_in_an_assignment() {
	f := func() (int, int) { return 6, 7 }

	_, x, _ := 3, 4, 7
	_, x = f() // evaluate f() but ignore first result value
	x, _ = f() // evaluate f() but ignore second result value

	_ = x // evaluate x but ignore it
	// error: No new variables on the left side of ':='
	// _ := 10
}

func Example_in_for_statements_with_range_clause() {
	s := []int{0, 1, 2, 3}

	for range s {
		fmt.Println("One iteration.")
	}
	for _ = range s { // redundant '_' expression
		fmt.Println("One iteration.")
	}

	for i := range s {
		fmt.Println(i, "iteration.")
	}
	for i, _ := range s { // redundant '_' expression
		fmt.Println(i, "iteration.")
	}
	for i, intV := range s {
		fmt.Println(i, intV)
	}
	for _, intV := range s {
		fmt.Println(intV)
	}
}

// interface checks
type Iface interface {
	a()
	b()
}

type ImpleIface struct{}

func (i *ImpleIface) a() {}
func (i *ImpleIface) b() {}

// interface checks
// guarantee *ImpleIface satisfies Iface
var _ Iface = (*ImpleIface)(nil)
