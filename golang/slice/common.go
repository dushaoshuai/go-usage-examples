package slice

import (
	"fmt"
)

func fmtSlice[T any](s []T) {
	fmt.Printf("slice: %v\nlen(slice): %d\ncap(slice): %d\n", s, len(s), cap(s))
}
