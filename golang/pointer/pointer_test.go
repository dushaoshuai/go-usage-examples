package pointer_test

import (
	"fmt"
	"testing"
)

func TestNestedPointerDereference(t *testing.T) {
	var a *****int

	b := 10
	c := &b
	d := &c
	e := &d
	f := &e
	a = &f

	fmt.Println(a)
	fmt.Println(*****a)
}
