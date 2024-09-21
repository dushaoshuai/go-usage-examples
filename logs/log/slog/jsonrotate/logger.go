package jsonrotate

import (
	"log/slog"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	logLevelVar = new(slog.LevelVar) // Info by default
)

func setJsonRotateLogger() {
	logFile := "/tmp/jsonrotatelogger/jsonratote.log"

	dailyRotateRule := logx.DefaultRotateRule(logFile, "-", 10, true)
	logWriter, err := logx.NewLogger(logFile, dailyRotateRule, true)
	if err != nil {
		panic(err)
	}

	jsonHandler := slog.NewJSONHandler(logWriter, &slog.HandlerOptions{
		AddSource:   true,
		Level:       logLevelVar,
		ReplaceAttr: nil,
	})
	slog.SetDefault(slog.New(jsonHandler))
}
