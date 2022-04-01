package functional_programming

import (
	"fmt"
	"time"
)

func Example_closure() {
	closure := returnClosure()
	fmt.Println(closure())
	time.Sleep(time.Second * 10)
	fmt.Println(closure())
	// Output:
	// 2022-03-30 21:57:53.565756 +0800 CST m=+0.000682543
	// 2022-03-30 21:57:53.565756 +0800 CST m=+0.000682543
}

func returnClosure() func() time.Time {
	now := time.Now()
	return func() time.Time {
		return now
	}
}
