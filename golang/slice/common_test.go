package slice_test

import (
	"fmt"
)

func fmtSlice[T any](name string, s []T) {
	fmt.Printf("%s: %v\nlen(%s): %d\ncap(%s): %d\n\n", name, s, name, len(s), name, cap(s))
}
