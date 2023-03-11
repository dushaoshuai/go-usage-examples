package times

import (
	"sync"
	"time"
)

type Times struct {
	lastDo time.Time
	every  time.Duration
	m      sync.Mutex
}

func NewTimes(every time.Duration) *Times {
	return &Times{
		every: every,
	}
}

func (t *Times) needDo() bool {
	return t.lastDo.IsZero() || t.lastDo.Add(t.every).Before(time.Now())
}

func (t *Times) Do(f func()) {
	if t.needDo() {
		t.doSlow(f)
	}
}

func (t *Times) doSlow(f func()) {
	t.m.Lock()
	defer t.m.Unlock()
	if t.needDo() {
		defer func() {
			t.lastDo = time.Now()
		}()
		f()
	}
}
