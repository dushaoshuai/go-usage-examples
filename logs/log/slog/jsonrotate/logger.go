package jsonrotate

import (
	"log/slog"

	"github.com/zeromicro/go-zero/core/logx"
)

type copyDataRotateLogger struct {
	w *logx.RotateLogger
}

func newCopyDataRotateLogger(logFile string) (copyDataRotateLogger, error) {
	dailyRotateRule := logx.DefaultRotateRule(logFile, "-", 10, true)
	rotateLogger, err := logx.NewLogger(logFile, dailyRotateRule, true)
	if err != nil {
		return copyDataRotateLogger{}, err
	}

	return copyDataRotateLogger{
		w: rotateLogger,
	}, nil
}

func (l copyDataRotateLogger) Close() error {
	return l.w.Close()
}

func (l copyDataRotateLogger) Write(data []byte) (int, error) {
	dataCopy := make([]byte, len(data))
	copy(dataCopy, data)
	return l.w.Write(dataCopy)
}

type JSONRotateLogger struct {
	w copyDataRotateLogger

	SlogLevelVar *slog.LevelVar
	slogger      *slog.Logger
}

func NewJSONRotateLogger(logFile string) (*JSONRotateLogger, error) {
	logWriter, err := newCopyDataRotateLogger(logFile)
	if err != nil {
		return nil, err
	}

	logLevelVar := new(slog.LevelVar) // Info by default

	jsonHandler := slog.NewJSONHandler(logWriter, &slog.HandlerOptions{
		AddSource:   true,
		Level:       logLevelVar,
		ReplaceAttr: nil,
	})

	return &JSONRotateLogger{
		SlogLevelVar: logLevelVar,
		w:            logWriter,
		slogger:      slog.New(jsonHandler),
	}, nil
}

func (l *JSONRotateLogger) Close() error {
	return l.w.Close()
}

func (l *JSONRotateLogger) Slogger() *slog.Logger {
	return l.slogger
}
