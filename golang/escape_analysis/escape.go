package escape_analysis

// $ go build -gcflags '-m -m -l' escape.go
// -m    print optimization decisions
// -l    disable inlining

func main() {
	returnPointer()

	var name fooImple = "Gopher"
	// name escapes to heap
	acceptFooI(name)

	// 对象过大
	slic1 := make([]int, 1024*1024)
	_ = slic1
	// 无法判断对象大小
	l := 100
	slic2 := make([]int, l)
	_ = slic2

	// closure
	incrm := returnClosure()
	incrm()
}

// 返回局部变量指针
func returnPointer() *int {
	a := 10
	return &a
}

type fooI interface {
	fooM()
}

type fooImple string

func (fooImple) fooM() {}

func acceptFooI(f fooI) {
	f.fooM()
}

func returnClosure() func() int {
	m := 10
	return func() int {
		m++
		return m
	}
}

// $ go build -gcflags '-m -l' escape.go
// # command-line-arguments
// ./escape.go:29:2: moved to heap: a
// ./escape.go:41:17: leaking param: f
// ./escape.go:46:2: moved to heap: m
// ./escape.go:47:9: func literal escapes to heap
// ./escape.go:12:13: name escapes to heap
// ./escape.go:15:15: make([]int, 1024 * 1024) escapes to heap
// ./escape.go:19:15: make([]int, l) escapes to heap
