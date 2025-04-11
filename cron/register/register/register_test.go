package register_test

import (
	"log/slog"
	"testing"
	"time"

	"github.com/robfig/cron/v3"

	"github.com/dushaoshuai/go-usage-examples/cron/register/globalconfig"
	"github.com/dushaoshuai/go-usage-examples/cron/register/register"
)

func TestRegister(t *testing.T) {
	register.Register(func(conf *globalconfig.Config) (spec string, job cron.Job, opts []cron.Option) {
		return "* * * * * *", cron.FuncJob(func() {
			slog.Info("without opts")
		}), nil
	})
	register.Register(func(conf *globalconfig.Config) (spec string, job cron.Job, opts []cron.Option) {
		return "*/5 * * * * *", cron.FuncJob(func() {
				slog.Info("with custom opts")
			}), []cron.Option{
				cron.WithSeconds(),
			}
	})

	stop := register.MustStartRegistered(nil)
	time.Sleep(20 * time.Second)
	stop()
}
