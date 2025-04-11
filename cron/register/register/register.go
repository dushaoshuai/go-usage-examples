package register

import (
	"context"
	"slices"
	"sync"
	"sync/atomic"

	"github.com/robfig/cron/v3"
	"github.com/samber/lo"

	"github.com/dushaoshuai/go-usage-examples/cron/register/globalconfig"
)

type Job func(conf *globalconfig.Config) (spec string, job cron.Job, opts []cron.Option)

var (
	jobsMu     sync.Mutex
	atomicJobs atomic.Value // []Job
)

func Register(job Job) {
	jobsMu.Lock()
	defer jobsMu.Unlock()

	jobs, _ := atomicJobs.Load().([]Job)
	atomicJobs.Store(append(jobs, job))
}

func DefaultOptions() []cron.Option {
	return []cron.Option{
		cron.WithSeconds(),
		cron.WithLogger(CronLogger{}),
		cron.WithChain(
			cron.Recover(CronLogger{}),
			delayIfStillRunning(CronLogger{}),
		),
	}
}

func MustStartRegistered(conf *globalconfig.Config) (stop func()) {
	stop = lo.Must(StartRegistered(conf))
	return
}

func StartRegistered(conf *globalconfig.Config) (stop func(), err error) {
	var stops []func() context.Context

	c := cron.New(DefaultOptions()...)
	stops = append(stops, func() context.Context {
		return c.Stop()
	})

	jobs, _ := atomicJobs.Load().([]Job)
	for job := range slices.Values(jobs) {
		spec, j, opts := job(conf)
		if len(opts) == 0 {
			_, err = c.AddJob(spec, j)
			if err != nil {
				return nil, err
			}
		} else {
			cc := cron.New(opts...)
			_, err = cc.AddJob(spec, j)
			if err != nil {
				return nil, err
			}
			stops = append(stops, func() context.Context {
				return cc.Stop()
			})
			cc.Start()
		}
	}

	c.Start()
	return func() {
		ctxs := make([]context.Context, 0, len(stops))

		for f := range slices.Values(stops) {
			ctxs = append(ctxs, f())
		}

		for ctx := range slices.Values(ctxs) {
			<-ctx.Done()
		}
	}, nil
}
