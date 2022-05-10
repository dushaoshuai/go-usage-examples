package type_declarations_test

import "fmt"

type aoo struct{}

func (a aoo) name() {
	fmt.Println("aoo")
}

type bar struct {
	aoo
}

// foo 继承了 bar 的嵌入结构体 aoo 的方法
type foo bar

func Example_struct_struct() {
	var x foo
	x.name()
	x.aoo.name()

	// Output:
	// aoo
	// aoo
}
