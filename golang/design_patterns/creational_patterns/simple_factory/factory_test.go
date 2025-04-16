package simplefactory_test

import (
	"fmt"

	"github.com/dushaoshuai/go-usage-examples/golang/design_patterns/creational_patterns/simple_factory"
)

func Example_simple_factory() {
	var factory simplefactory.CommandFactory

	command := factory.GetCommand(simplefactory.Echo)
	fmt.Println(command.Name(), command.Execute())

	// Output:
	// echo <nil>
}
