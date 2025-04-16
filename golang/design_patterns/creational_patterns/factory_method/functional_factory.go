package factorymethod

type CmdFactory func() Command

func NewEchoCommand() Command { return echo{} }

func NewCdCommand() Command { return cd{} }

func NewPwdCommand() Command { return pwd{} }
