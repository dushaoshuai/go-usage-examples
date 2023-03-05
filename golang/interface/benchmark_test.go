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
// 接口速度还是比较慢的啊

// goos: darwin
// goarch: amd64
// pkg: github.com/dushaoshuai/go-usage-examples/golang/interface
// cpu: Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz
// BenchmarkDirectCall_pointer
// BenchmarkDirectCall_pointer-8                	1000000000	         0.3399 ns/op
// BenchmarkDynamicDispatch_pointer
// BenchmarkDynamicDispatch_pointer-8           	576121688	         1.949 ns/op
// BenchmarkDirectCall_struct
// BenchmarkDirectCall_struct-8                 	1000000000	         0.3276 ns/op
// BenchmarkBenchmarkDynamicDispatch_struct
// BenchmarkBenchmarkDynamicDispatch_struct-8   	561326664	         2.104 ns/op
// PASS
