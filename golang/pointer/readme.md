# go 的指针

## 指针的三种表现形式

* `*T`
* [uintprt](https://pkg.go.dev/builtin#uintptr)
* [unsafe.Pointer](https://cs.opensource.google/go/go/+/go1.18:src/unsafe/unsafe.go;l=184)

`unsafe.Pointer` 是连接 `*T` 和 `uintprt` 的纽带，可以借助 `unsafe.Pointer` 在 `T` 和 `uintptr` 之间转换，
`*T` 就是 C 语言中的指针，但不支持指针运算。`uintptr` 是一个整数，可以进行运算，所以可以进行这样的转换：
`*T` --> `unsafe.Pointer` --> `uintptr` --> arithmetic --> `unsafe.Pointer` --> `*T`。
具体可以参看 `unsafe.Pointer` 的文档。
