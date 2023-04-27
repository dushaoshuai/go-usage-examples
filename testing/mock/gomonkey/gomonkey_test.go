package gomonkey_test

import (
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/dushaoshuai/goloop"
)

func TestApplyFunc(t *testing.T) {
	// go test -gcflags=all=-l -run TestApplyFunc
	now := time.Date(2001, 3, 4, 5, 6, 7, 8, time.UTC)
	patches := gomonkey.ApplyFunc(time.Now, func() time.Time {
		return now
	})
	defer patches.Reset()

	for range goloop.Repeat(10) {
		if n := time.Now(); !n.Equal(now) {
			t.Errorf("time.Now() should equal now, got %v", n)
		}
	}
}

func TestApplyFuncReturn(t *testing.T) {
	now := time.Date(2001, 3, 4, 5, 6, 7, 8, time.UTC)
	patches := gomonkey.ApplyFuncReturn(time.Now, now)
	defer patches.Reset()

	for range goloop.Repeat(10) {
		if n := time.Now(); !n.Equal(now) {
			t.Errorf("time.Now() should equal now, got %v", n)
		}
	}
}

func TestApplyFuncSeq(t *testing.T) {
	now := time.Date(2001, 3, 4, 5, 6, 7, 8, time.UTC)
	patches := gomonkey.ApplyFuncSeq(time.Now, []gomonkey.OutputCell{
		{Values: gomonkey.Params{now}, Times: 0}, // 0 次也会算一次
		{Values: gomonkey.Params{now}, Times: 2},
	})
	defer patches.Reset()

	for range goloop.Repeat(3) {
		if n := time.Now(); !n.Equal(now) {
			t.Errorf("time.Now() should equal now, got %v", n)
		}
	}
}

func TestApplyFuncVar(t *testing.T) {
	hello := "hello"
	patches := gomonkey.NewPatches()
	patches.ApplyFuncVar(&f1, func(s string, i int, b bool) string {
		return hello
	})
	patches.ApplyFuncVar(&f2, func(s string, i int, b bool) string {
		return hello
	})
	defer patches.Reset()

	if got := f1("a", 1, true); got != hello {
		t.Errorf("patched f1 should return %s, but got %s", hello, got)
	}
	if got := f2("a", 1, true); got != hello {
		t.Errorf("patched f2 should return %s, but got %s", hello, got)
	}
}

func TestApplyFuncVarReturn(t *testing.T) {
	hello := "hello"
	patches := gomonkey.NewPatches()
	patches.ApplyFuncVarReturn(&f1, hello)
	patches.ApplyFuncVarReturn(&f2, hello)
	defer patches.Reset()

	if got := f1("a", 1, true); got != hello {
		t.Errorf("patched f1 should return %s, but got %s", hello, got)
	}
	if got := f2("a", 1, true); got != hello {
		t.Errorf("patched f2 should return %s, but got %s", hello, got)
	}
}

func TestApplyFuncVarSeq(t *testing.T) {
	l1 := "C"
	l2 := "Go"
	patches := gomonkey.NewPatches()
	patches.ApplyFuncVarSeq(&f1, []gomonkey.OutputCell{
		{Values: gomonkey.Params{l1}, Times: 2},
		{Values: gomonkey.Params{l2}, Times: 3},
	})
	patches.ApplyFuncVarSeq(&f2, []gomonkey.OutputCell{
		{Values: gomonkey.Params{l1}, Times: 2},
		{Values: gomonkey.Params{l2}, Times: 3},
	})
	defer patches.Reset()

	for range goloop.Repeat(2) {
		if got := f1("a", 1, false); got != l1 {
			t.Errorf("patched f1 should return %s, got %s", l1, got)
		}
	}
	for range goloop.Repeat(3) {
		if got := f1("a", 1, false); got != l2 {
			t.Errorf("patched f1 should return %s, got %s", l2, got)
		}
	}
	for range goloop.Repeat(2) {
		if got := f2("a", 1, false); got != l1 {
			t.Errorf("patched f2 should return %s, got %s", l1, got)
		}
	}
	for range goloop.Repeat(3) {
		if got := f2("a", 1, false); got != l2 {
			t.Errorf("patched f2 should return %s, got %s", l2, got)
		}
	}

	defer func() {
		if err := recover(); err == nil {
			t.Errorf("expected panic didn't happen")
		}
	}()
	f2("a", 1, false)
}

func TestApplyGlobalVar(t *testing.T) {
	v1 := 150
	v2 := 160
	patches := gomonkey.NewPatches()
	patches.ApplyGlobalVar(&privateVar, v1)
	patches.ApplyGlobalVar(&PublicVar, v2)
	if privateVar != v1 {
		t.Fatalf("patched privateVar should equal %d, got %d", v1, privateVar)
	}
	if PublicVar != v2 {
		t.Fatalf("patched PublicVar should equal %d, got %d", v2, PublicVar)
	}

	patches.Reset()
	if privateVar != 10 {
		t.Fatalf("reseted privateVar should equal %d, got %d", 10, privateVar)
	}
	if PublicVar != 20 {
		t.Fatalf("reseted PublicVar should equal %d, got %d", 20, PublicVar)
	}
}
