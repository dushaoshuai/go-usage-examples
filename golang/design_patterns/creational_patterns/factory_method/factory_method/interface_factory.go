package interface_factory

type Command interface {
	Name() string
	Execute() error
}

type echo struct{}

func NewEcho() Command { return echo{} }

func (e echo) Name() string   { return "echo" }
func (e echo) Execute() error { return nil }

type cd struct{}

func NewCd() Command { return cd{} }

func (c cd) Name() string { return "cd" }

func (c cd) Execute() error { return nil }
