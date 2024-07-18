package time_test

import (
	"fmt"
	"time"

	"github.com/jinzhu/now"
)

func ExampleDate() {
	t1 := time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(t1)

	// The month, day, hour, min, sec, and nsec values may be outside their usual ranges and will be normalized during the conversion.
	t2 := time.Date(2024, 0, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(t2)

	// The month, day, hour, min, sec, and nsec values may be outside their usual ranges and will be normalized during the conversion.
	t3 := time.Date(2024, -1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(t3)

	endOfLastMonthOfT1 := now.With(t1).Add(-time.Nanosecond)
	beginningOfLastMonthOfT1 := now.With(endOfLastMonthOfT1).BeginningOfMonth()
	fmt.Println(beginningOfLastMonthOfT1.Equal(t2))

	// Output:
	// 2024-01-01 00:00:00 +0000 UTC
	// 2023-12-01 00:00:00 +0000 UTC
	// 2023-11-01 00:00:00 +0000 UTC
	// true
}
