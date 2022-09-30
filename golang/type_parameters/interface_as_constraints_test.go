package type_parameters_test

import "fmt"

type foo struct{}

func (f *foo) String() string {
	return "My type is foo !!!"
}

// 接口作为约束
func bar[T fmt.Stringer](x T) {
	fmt.Println(x.String())
}

func Example_interface_as_constraints() {
	bar(&foo{})
	// Output:
	// My type is foo !!!
}

func baz[T interface { // 字面量也可作约束
	a()
	b()
	c()
}](t T) {
	t.a()
	t.b()
	t.c()
}

type qux struct{}

func (qux) a() { fmt.Println("a") }
func (qux) b() { fmt.Println("b") }
func (qux) c() { fmt.Println("c") }

func Example_literal_interface_as_constraints() {
	baz(qux{})
	// Output:
	// a
	// b
	// c
}
