package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("runtime.GOMAXPROCS():", runtime.GOMAXPROCS(0))
	for {
		go func() {
			fmt.Println(0)
			fmt.Println("runtime.NumGoroutine():", runtime.NumGoroutine())
		}()
		fmt.Println(1)
		fmt.Println("runtime.NumGoroutine():", runtime.NumGoroutine())
	}
}

// $ GOMAXPROCS=1 go run main.go
// runtime.GOMAXPROCS(): 1
// 1
// runtime.NumGoroutine(): 2
// 1
// runtime.NumGoroutine(): 3
// 1
// runtime.NumGoroutine(): 4
// 1
// runtime.NumGoroutine(): 5
// ...
// ...
// 1
// runtime.NumGoroutine(): 3066
// 1
// runtime.NumGoroutine(): 3067
// 1
// runtime.NumGoroutine(): 3068
// 1
// runtime.NumGoroutine(): 3069
// 0
// runtime.NumGoroutine(): 3070
// 0
// runtime.NumGoroutine(): 3069
// 0
// runtime.NumGoroutine(): 3068
// 0
// runtime.NumGoroutine(): 3067
// ...
// ...
// 0
// runtime.NumGoroutine(): 5
// 0
// runtime.NumGoroutine(): 4
// 0
// runtime.NumGoroutine(): 3
// 0
// runtime.NumGoroutine(): 2
// 1
// runtime.NumGoroutine(): 1
// 1
// runtime.NumGoroutine(): 2
// 1
// runtime.NumGoroutine(): 3
// 1
// runtime.NumGoroutine(): 4
// ...

// 在上面的这段输出结果中，
// 因为 Go 程序只能使用一个逻辑 CPU，
// 同一时刻最多只有一个线程在执行 goroutine,
// 同一时刻最多只有一个 goroutine 在运行。
// Go 调度器调度这些 goroutine,
// 先是打印 1 的 main goroutine,
// 再是打印 0 的其他 goroutine,
// 再是打印 1 的 main goroutine。
// main goroutine 一直存在，
// 打印 0 的其他 goroutine 打印完 0 就终止了，
// 因此可以观察到，
// 打印 1 时 goroutine 数量增多，
// 打印 0 时 goroutine 数量减少。
