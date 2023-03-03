package decorator_test

import "api-examples/golang/design_patterns/structural_patterns/decorator/decorator"

func Example_decorator() {
	var a decorator.Component
	a = decorator.NewConcreteComponentA()
	a = decorator.DecorateByA(a)
	a = decorator.DecorateByB(a)
	a.OperationA()
	a.OperationB()

	var b decorator.Component
	b = decorator.NewConcreteComponentB()
	b = decorator.DecorateByA(b)
	b.OperationB()

	// Output:
}
