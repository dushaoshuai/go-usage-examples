package simplefactory

type Command interface {
	Name() string
	Execute() error
}

type echo struct{}

func (e echo) Name() string   { return "echo" }
func (e echo) Execute() error { return nil }

type cd struct{}

func (c cd) Name() string   { return "cd" }
func (c cd) Execute() error { return nil }

type pwd struct{}

func (p pwd) Name() string   { return "pwd" }
func (p pwd) Execute() error { return nil }
