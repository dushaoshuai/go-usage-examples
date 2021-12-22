package interface_test

import "testing"

type person interface {
	speak()
}

type chinese struct {
	name string
	age  int
}

func (c *chinese) speak() {}

func BenchmarkDirectCall_pointer(b *testing.B) {
	c := &chinese{}
	for i := 0; i < b.N; i++ {
		c.speak()
	}
}

func BenchmarkDynamicDispatch_pointer(b *testing.B) {
	c := person(&chinese{})
	for i := 0; i < b.N; i++ {
		c.speak()
	}
}

type japanese struct {
	name string
	age  int
}

func (j japanese) speak() {}

func BenchmarkDirectCall_struct(b *testing.B) {
	j := japanese{}
	for i := 0; i < b.N; i++ {
		j.speak()
	}
}

func BenchmarkBenchmarkDynamicDispatch_struct(b *testing.B) {
	j := person(japanese{})
	for i := 0; i < b.N; i++ {
		j.speak()
	}
}

// https://draveness.me/golang/docs/part2-foundation/ch04-basic/golang-interface/#%E5%9F%BA%E5%87%86%E6%B5%8B%E8%AF%95

// goos: darwin
// goarch: amd64
// pkg: api-examples/golang/interface
// cpu: Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz
// BenchmarkDirectCall_pointer
// BenchmarkDirectCall_pointer-8                	1000000000	         0.4822 ns/op
// BenchmarkDynamicDispatch_pointer
// BenchmarkDynamicDispatch_pointer-8           	415445734	         2.835 ns/op
// BenchmarkDirectCall_struct
// BenchmarkDirectCall_struct-8                 	1000000000	         0.4986 ns/op
// BenchmarkBenchmarkDynamicDispatch_struct
// BenchmarkBenchmarkDynamicDispatch_struct-8   	400639581	         3.046 ns/op
