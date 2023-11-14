// Package errgroup provides synchronization, error propagation, and Context
// cancelation for groups of goroutines working on subtasks of a common task.
package errgroup

import (
	"context"
	"sync"
)

type token struct{}

// A Group is a collection of goroutines working on subtasks
// that are part of the same overall task.
//
// A zero Group is valid, has no limit on the number of active goroutines,
// and does not cancel on error.
type Group struct {
	wg   sync.WaitGroup
	sema chan token

	cancel  func()
	errOnce sync.Once
	err     error
}

// WithContext returns a new Group and an associated Context derived from ctx.
//
// The derived Context is canceled the first time a function passwd to Go returns
// a non-nil error or the first time Wait returns, whichever occurs first.
func WithContext(ctx context.Context) (*Group, context.Context) {
	childCtx, cancel := context.WithCancel(ctx)
	return &Group{
		cancel: cancel,
	}, childCtx
}

func (g *Group) done() {
	if g.sema != nil {
		<-g.sema
	}
	g.wg.Done()
}

// Go calls the given function in a new goroutine. It blocks until the new
// goroutine can be added without the number of active goroutines in the group
// exceeding the configured limit.
//
// The first call to return a non-nil error cancels the group's context, if the
// group was created by calling WithContext. The error will be returned by Wait.
func (g *Group) Go(f func() error) {
	g.wg.Add(1)

	if g.sema != nil {
		g.sema <- token{}
	}

	go func() {
		defer g.done()

		err := f()
		if err != nil {
			g.errOnce.Do(func() {
				g.err = err
				if g.cancel != nil {
					g.cancel()
				}
			})
		}
	}()
}

// TryGo calls the given function in a new goroutine only if the number of active
// goroutines in the group is currently below the configured limit.
//
// The return value reports whether the goroutine was started.
func (g *Group) TryGo(f func() error) bool {
	if g.sema != nil {
		select {
		case g.sema <- token{}:
			// Note: this allows barging iff channels in general allow barging.
		default:
			return false
		}
	}

	g.wg.Add(1)
	go func() {
		defer g.done()

		err := f()
		if err != nil {
			g.errOnce.Do(func() {
				g.err = err
				if g.cancel != nil {
					g.cancel()
				}
			})
		}
	}()
	return true
}

// SetLimit limits the number of active goroutines in this group to at most n.
// A negative value indicates no limit.
//
// Any subsequent call to the Go method will block until it can add an active
// goroutine without exceeding the configured limit.
//
// The limit must not be modified while any goroutines in the group are active.
func (g *Group) SetLimit(n int) {
	if n < 0 {
		g.sema = nil
		return
	}
	if len(g.sema) > 0 {
		panic("errgroup: modify limit while there are active goroutines")
	}
	g.sema = make(chan token, n)
}

// Wait blocks until all function calls from the Go method have returned,
// the returns the first non-nil error (if any) form them.
func (g *Group) Wait() error {
	g.wg.Wait()
	if g.cancel != nil {
		g.cancel()
	}
	return g.err
}
