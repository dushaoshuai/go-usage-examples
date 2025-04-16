package simplefactory

const (
	Echo = "echo"
	Cd   = "cd"
	Pwd  = "pwd"
)

type CommandFactory struct{}

func (CommandFactory) GetCommand(typ string) Command {
	switch typ {
	case Echo:
		return echo{}
	case Cd:
		return cd{}
	case Pwd:
		return pwd{}
	default:
		return nil
	}
}
