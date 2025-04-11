package register

import (
	"log/slog"
)

// CronLogger implements github.com/robfig/cron/v3.Logger.
type CronLogger struct{}

func (c CronLogger) Info(msg string, keysAndValues ...any) {
	slog.Info(msg, keysAndValues...)
}

func (c CronLogger) Error(err error, msg string, keysAndValues ...any) {
	slog.Error(msg, append(keysAndValues, slog.Any("error", err))...)
}
