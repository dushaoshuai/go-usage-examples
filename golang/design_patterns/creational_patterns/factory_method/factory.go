package factorymethod

type CommandFactory interface {
	NewCommand() Command
}

type echoFactory struct{}

func NewEchoFactory() CommandFactory { return echoFactory{} }

func (e echoFactory) NewCommand() Command { return echo{} }

type cdFactory struct{}

func NewCdFactory() CommandFactory { return cdFactory{} }

func (c cdFactory) NewCommand() Command { return cd{} }

type pwdFactory struct{}

func NewPwdFactory() CommandFactory { return pwdFactory{} }

func (p pwdFactory) NewCommand() Command { return pwd{} }
