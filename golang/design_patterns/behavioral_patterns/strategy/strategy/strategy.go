package strategy

type Strategy interface {
	Execute(int, int) int
}

type Add struct{}

func (s *Add) Execute(a, b int) int {
	return a + b
}

type Subtract struct{}

func (s *Subtract) Execute(a, b int) int {
	return a - b
}

type Multiply struct{}

func (s *Multiply) Execute(a, b int) int {
	return a * b
}

type Context struct {
	strategy Strategy
}

func NewContext(strategy Strategy) *Context {
	return &Context{strategy: strategy}
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy(a, b int) int {
	return c.strategy.Execute(a, b)
}
