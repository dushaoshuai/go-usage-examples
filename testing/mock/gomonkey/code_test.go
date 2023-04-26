package gomonkey_test

import "fmt"

type funcType func(string, int, bool) string

var (
	f1 = func(s string, i int, b bool) string {
		return fmt.Sprintf("%v:%v:%v", s, i, b)
	}
	f2 funcType = func(s string, i int, b bool) string {
		return fmt.Sprintf("%v:%v:%v", s, i, b)
	}
)

var (
	privateVar = 10
	PublicVar  = 20
)
