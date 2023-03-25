package command

import (
	"fmt"
)

type Invoker1 struct {
	command Command
}

func NewInvoker1(c Command) *Invoker1 {
	return &Invoker1{command: c}
}

func (i *Invoker1) Invoke() {
	fmt.Println("Invoker1 calling Command.")
	i.command.Execute()
}

func Invoker2(c Command) {
	fmt.Println("Invoker2 calling Command.")
	c.Execute()
}
