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
	// The underlying *logx.RotateLogger instance writes logs asynchronously,
	// and the *slog.commonHandler reuses buffer instances to improve efficiency.
	// To ensure that the *logx.RotateLogger does not encounter a mutated buffer
	// (which could happen due to the asynchronous nature and buffer reuse),
	// we create a copy of the input data bytes before passing them along.
	// This step guarantees that each log write operation is handled with an independent buffer copy.
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
