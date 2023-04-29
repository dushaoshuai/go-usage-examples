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

type fooErr string

func (e fooErr) Error() string {
	return string(e)
}

func (e fooErr) String(prefix string) string {
	return prefix + "/error: " + string(e)
}

func (e fooErr) ok() bool {
	return string(e) == ""
}
