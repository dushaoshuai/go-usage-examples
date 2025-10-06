// Package link_X
//
// $ go help build
// usage: go build [-o output] [build flags] [packages]
//
// The build flags are shared by the build, clean, get, install, list, run,
// and test commands:
//
//        -ldflags '[pattern=]arg list'
//                arguments to pass on each go tool link invocation.
//
// $ go tool link
// usage: link [options] main.o
//  -X definition
//        add string value definition of the form importpath.name=value
//
// -ldflags="-X ..." 是在 Go 链接阶段 动态修改二进制中全局字符串变量的值的机制。
// 可参考 https://github.com/mattn/go-sqlite3/blob/8bf7a8a844faf952aa0245b4c0ad0a47e84f4efd/sqlite3.go#L248 .

package link_X

var a = "a" // 字符串变量可在链接时修改

const b = "b" // 常量不可被修改
