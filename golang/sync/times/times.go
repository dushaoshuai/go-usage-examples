package times

import (
	"sync"
	"time"
)

// A Times will perform exactly one successful action at intervals.
// A Times may be used by multiple goroutines simultaneously.
type Times struct {
	last   time.Time // the last time one successful action was performed
	period time.Duration
	m      sync.Mutex
}

// NewTimes returns a Times that will perform exactly one successful action each time d duration elapses.
func NewTimes(d time.Duration) *Times {
	return &Times{
		period: d,
	}
}

// needDo reports whether an action should be performed.
func (t *Times) needDo() bool {
	return t.last.IsZero() || t.last.Add(t.period).Before(time.Now())
}

// Do calls the function f if and only if f has never been called successfully
// since the specified duration elapsed. In other words, within current interval,
// f will be invoked each time Do is called unless the previous call to f returns without error.
// After a successful call to f returns, next interval starts.
func (t *Times) Do(f func() error) error {
	if !t.needDo() {
		return nil
	}

	t.m.Lock()
	defer t.m.Unlock()
	if !t.needDo() {
		return nil
	}
	err := f()
	if err == nil {
		t.last = time.Now()
	}
	return err
}
