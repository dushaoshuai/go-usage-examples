package HandlersChain_test

import (
	"fmt"

	. "github.com/dushaoshuai/go-usage-examples/gin/HandlersChain"
)

func dataHandler(ctx *Context) {
	fmt.Println("data", ctx.Data)
	ctx.Data++
}

func nextHandler(ctx *Context) {
	fmt.Println("next", ctx.Data)
	ctx.Next()
	fmt.Println("next", ctx.Data)

	ctx.Data++
}

// todo 补充文字说明
func ExampleContext_Next() {
	handlers := HandlersChain{
		dataHandler,
		dataHandler,
		dataHandler,
		nextHandler,
		nextHandler,
		dataHandler,
		nextHandler,
		nextHandler,
		dataHandler,
	}

	ctx := NewContext(0, handlers...)
	ctx.Next()

	// Output:
	// data 0
	// data 1
	// data 2
	// next 3
	// next 3
	// data 3
	// next 4
	// next 4
	// data 4
	// next 5
	// next 6
	// next 7
	// next 8
}
