package type_declarations_test

import "fmt"

// A defined type may have methods associated with it.
// It does not inherit any methods bound to the given type,
// but the method set of an interface type or of elements of a composite type remains unchanged.

type echoer interface {
	echo()
}

// cmder 和 echoer 的 method set 相同
type cmder echoer

type echo func()

func (e echo) echo() { e() }

func Example_interface_interface() {
	var f echo = func() {
		fmt.Println("echo ech ec e ...")
	}
	var c cmder = f // cmder 有一个方法 echo()
	c.echo()

	// Output:
	// echo ech ec e ...
}
