package command_test

import (
	"github.com/dushaoshuai/go-usage-examples/golang/design_patterns/behavioral_patterns/command/command"
)

func ExampleCommand() {
	i1 := command.NewInvoker1(
		command.NewCommand1("Go."),
	)
	i2 := command.NewInvoker1(
		command.NewCommand2("Python."),
	)
	i3 := command.NewInvoker1(
		command.NewCommand3("C."),
	)
	i1.Invoke()
	i2.Invoke()
	i3.Invoke()

	command.Invoker2(command.NewCommand2("Go."))
	command.Invoker2(command.NewCommand1("Python."))
	command.Invoker2(command.NewCommand1("Java."))

	// Output:
	// Invoker1 calling Command.
	// Command1 executing. Go.
	// Invoker1 calling Command.
	// Command2 executing. Python.
	// Invoker1 calling Command.
	// Command3 executing. Command3 Hello. C.
	// Invoker2 calling Command.
	// Command2 executing. Go.
	// Invoker2 calling Command.
	// Command1 executing. Python.
	// Invoker2 calling Command.
	// Command1 executing. Java.
}
