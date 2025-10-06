package link_X_test

import (
	"fmt"
)

// 测试包中的字符串变量不可被修改
var a = "a"

const b = "b" // 常量不可被修改

//go:generate go test -ldflags="-X 'github.com/dushaoshuai/go-usage-examples/golang/cmdgo/link-X.a=A' -X 'github.com/dushaoshuai/go-usage-examples/golang/cmdgo/link-X.b=B'" -test.run ^Example_link_testing_X$
func Example_link_testing_X() {
	fmt.Printf("a == %v\n", a)
	fmt.Printf("b == %v\n", b)

	// Output:
	// a == a
	// b == b
}
