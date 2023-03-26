package decorator

type Component interface {
	OperationA()
	OperationB()
}

type baseComponent struct{}

func NewBaseComponent() Component    { return &baseComponent{} }
func (c *baseComponent) OperationA() {}
func (c *baseComponent) OperationB() {}

type DecoratorA struct {
	Component
}

func DecorateByA(c Component) Component {
	da := &DecoratorA{}
	da.Component = c
	return da
}

func (d *DecoratorA) addedBehavior() {}

func (d *DecoratorA) OperationA() {
	d.addedBehavior()
	defer d.addedBehavior()
	d.Component.OperationA()
}

type DecoratorB struct {
	Component
}

func DecorateByB(c Component) Component {
	db := &DecoratorB{}
	db.Component = c
	return db
}

func (d *DecoratorB) addedBehavior() {}

func (d *DecoratorB) OperationB() {
	d.Component.OperationB()
	d.addedBehavior()
}
