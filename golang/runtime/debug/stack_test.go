package debug_test

import (
	"os"
	"runtime/debug"
)

func ExampleDebugStack() {
	stack := debug.Stack()
	os.Stdout.Write(stack)

	// Output:
	// goroutine 1 [running]:
	// runtime/debug.Stack()
	//	/home/shaouai/.local/go/src/runtime/debug/stack.go:24 +0x65
	// github.com/dushaoshuai/go-usage-examples/golang/runtime/debug_test.ExampleDebugStack()
	//	/home/shaouai/dev/github.com/dushaoshuai/go-usage-examples/golang/runtime/debug/stack_test.go:9 +0x19
	// testing.runExample({{0x528e0a, 0x11}, 0x5317c8, {0x0, 0x0}, 0x0})
	//	/home/shaouai/.local/go/src/testing/run_example.go:63 +0x2d0
	// testing.runExamples(0xc00006ddd8, {0x604760?, 0x1, 0x0?})
	//	/home/shaouai/.local/go/src/testing/example.go:44 +0x17d
	// testing.(*M).Run(0xc0000a40a0)
	//	/home/shaouai/.local/go/src/testing/testing.go:1908 +0x6ef
	// main.main()
	//	_testmain.go:47 +0x1aa
}
