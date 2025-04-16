package factorymethod_test

import (
	"fmt"

	"github.com/dushaoshuai/go-usage-examples/golang/design_patterns/creational_patterns/factory_method"
)

func Example_factory_method() {
	var factory factorymethod.CommandFactory

	factory = factorymethod.NewEchoFactory()
	command := factory.NewCommand()
	fmt.Println(command.Name(), command.Execute())

	factory = factorymethod.NewCdFactory()
	command = factory.NewCommand()
	fmt.Println(command.Name(), command.Execute())

	factory = factorymethod.NewPwdFactory()
	command = factory.NewCommand()
	fmt.Println(command.Name(), command.Execute())

	// Output:
	// echo <nil>
	// cd <nil>
	// pwd <nil>
}

func Example_functional_factory_method() {
	var factory factorymethod.CmdFactory

	factory = factorymethod.NewEchoCommand
	command := factory()
	fmt.Println(command.Name(), command.Execute())

	factory = factorymethod.NewCdCommand
	command = factory()
	fmt.Println(command.Name(), command.Execute())

	factory = factorymethod.NewPwdCommand
	command = factory()
	fmt.Println(command.Name(), command.Execute())

	// Output:
	// echo <nil>
	// cd <nil>
	// pwd <nil>
}
