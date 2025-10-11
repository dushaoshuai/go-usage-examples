package interface_test

import "fmt"

type implementEcho struct{}

func (ie implementEcho) echo() {
	fmt.Println("echo ech ec e ...")
}

func Example_interface_literal() {
	var ie implementEcho
	var ifce interface{} = ie

	if echo, ok := ifce.(interface {
		echo()
	}); ok {
		echo.echo()
	}

	// Output:
	// echo ech ec e ...
}
