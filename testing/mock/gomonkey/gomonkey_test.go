package gomonkey_test

import (
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/dushaoshuai/goloop"
)

func TestApplyFunc(t *testing.T) {
	now := time.Date(2001, 3, 4, 5, 6, 7, 8, time.UTC)
	patches := gomonkey.ApplyFunc(time.Now, func() time.Time {
		return now
	})
	defer patches.Reset()

	if !time.Now().Equal(now) {
		t.Errorf("time.Now() should equal now, but it didn't")
	}
}

func TestApplyFuncReturn(t *testing.T) {
	now := time.Date(2001, 3, 4, 5, 6, 7, 8, time.UTC)
	patches := gomonkey.ApplyFuncReturn(time.Now, now)
	defer patches.Reset()

	if !time.Now().Equal(now) {
		t.Errorf("time.Now() should equal now, but it didn't")
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
		if !time.Now().Equal(now) {
			t.Errorf("time.Now() should equal now, but it didn't")
		}
	}
}
