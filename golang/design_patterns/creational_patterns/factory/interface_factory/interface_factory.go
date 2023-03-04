package interface_factory

type Command interface {
	Name() string
	Execute() error
}

// 对外只暴露接口，不暴露具体的实现

type echo struct{}

func (e echo) Name() string {
	return "echo"
}

func (e echo) Execute() error {
	return nil
}

func NewEcho() Command {
	return echo{}
}

type cd struct{}

func (c cd) Name() string {
	return "cd"
}

func (c cd) Execute() error {
	return nil
}

func NewCd() Command {
	return cd{}
}
