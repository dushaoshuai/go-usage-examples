package command

import (
	"fmt"
)

type Command interface {
	Execute()
}

type Command1 struct {
	param string
}

func NewCommand1(param string) *Command1 {
	return &Command1{param: param}
}

func (c *Command1) Execute() {
	fmt.Println("Command1 executing.", c.param)
}

type Command2 func()

func NewCommand2(param string) Command2 {
	return func() {
		fmt.Println("Command2 executing.", param)
	}
}

func (c Command2) Execute() {
	c()
}

type Command3 struct {
	param string
	msg   string
}

func NewCommand3(param string) *Command3 {
	return &Command3{
		param: param,
		msg:   "Command3 Hello.",
	}
}

func (c *Command3) Execute() {
	fmt.Println("Command3 executing.", c.msg, c.param)
}
