package oncevery

import (
	"sync"
	"time"
)

type OnceEvery struct {
	t     time.Time // todo rename
	every time.Duration
	m     sync.Mutex
}

func NewOnceEvery(every time.Duration) *OnceEvery {
	return &OnceEvery{
		every: every,
	}
}

func (o *OnceEvery) needDo() bool {
	return o.t.IsZero() ||
		o.t.Add(o.every).Before(time.Now())
}

func (o *OnceEvery) Do(f func()) {
	if o.needDo() {
		o.doSlow(f)
	}
}

func (o *OnceEvery) doSlow(f func()) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.needDo() {
		defer func() {
			o.t = time.Now()
		}()
		f()
	}
}
