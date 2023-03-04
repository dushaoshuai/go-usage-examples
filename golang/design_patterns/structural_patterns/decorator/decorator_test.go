package decorator_test

import "github.com/dushaoshuai/go-usage-examples/golang/design_patterns/structural_patterns/decorator/decorator"

func Example_decorator() {
	var a decorator.Component
	a = decorator.NewConcreteComponentA()
	a = decorator.DecorateByA(a)
	a = decorator.DecorateByB(a)
	a = decorator.DecorateByB(a)
	a = decorator.DecorateByA(a)
	a = decorator.DecorateByA(a)
	a.OperationA()
	a.OperationB()

	var b decorator.Component
	b = decorator.NewConcreteComponentB()
	b = decorator.DecorateByA(b)
	b.OperationB()

	// Output:
}
