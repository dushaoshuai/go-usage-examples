package register

import (
	"reflect"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

// delayIfStillRunning serializes jobs, delaying subsequent runs until the
// previous one is complete. Jobs running after a delay of more than a minute
// have the delay logged at Info.
func delayIfStillRunning(logger cron.Logger) cron.JobWrapper {
	return func(j cron.Job) cron.Job {
		var mu sync.Mutex
		return cron.FuncJob(func() {
			start := time.Now()
			mu.Lock()
			defer mu.Unlock()

			if dur := time.Since(start); dur > time.Minute {
				var jobName string
				if namer, ok := j.(JobNamer); ok {
					jobName = namer.JobName()
				} else {
					jobName = reflect.TypeOf(j).Name()
				}

				logger.Info("delay", "job", jobName, "duration", dur.String())
			}

			j.Run()
		})
	}
}

type JobNamer interface {
	JobName() string
}
