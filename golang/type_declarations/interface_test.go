package type_declarations_test

import "fmt"

type echoer interface {
	echo()
}

// cmder 继承了 echoer 的 method set
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
