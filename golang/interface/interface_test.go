package interface_test

type bar interface {
	Add(x int) int
}

type age struct {
	a int
}

func (a *age) Add(y int) int { // 实现方法时，参数名不必相同
	a.a += y
	return a.a
}

func ExampleInterface() {

	// Output:
}
