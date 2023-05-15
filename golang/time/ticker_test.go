package time_test

import (
	"fmt"
	"time"
)

func ExampleTick() {
	// time.Ticker will not send the first tick immediately.
	// The first tick is delivered until the specified duration elapses.
	fmt.Println(time.Now())
	fmt.Println(<-time.Tick(10 * time.Second))
	// Output:
	// 2023-05-15 17:45:14.30176 +0800 CST m=+0.000603842
	// 2023-05-15 17:45:24.301992 +0800 CST m=+10.001274124
}
