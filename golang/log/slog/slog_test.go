package slog_test

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/xhd2015/xgo/runtime/mock"
)

func Example_json_handler() {
	// xgo test -run Example_json_handler
	mock.Patch(time.Now, func() time.Time {
		return time.Date(2024, 8, 25, 16, 50, 7, 12234, time.Local)
	})

	slog.SetDefault(
		slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				AddSource:   true,
				Level:       slog.LevelDebug,
				ReplaceAttr: nil,
			}),
		),
	)

	slog.Debug("a debug level msg", "x1", 1, slog.Int("x2", 2))
	slog.DebugContext(context.Background(), "a debug level msg", "x1", 1, slog.String("x2", "x2 value"))

	slog.Warn("a warn level msg", "x1", 1, slog.Int("x2", 2))
	slog.WarnContext(context.Background(), "a warn level msg", "x1", 1, slog.Int("x2", 2))

	// Output:
	// {"time":"2024-08-25T16:50:07.000012234+08:00","level":"DEBUG","source":{"function":"github.com/dushaoshuai/go-usage-examples/golang/log/slog_test.Example_json_handler","file":"/tmp/go-usage-examples/golang/log/slog/slog_test.go","line":27},"msg":"a debug level msg","x1":1,"x2":2}
	// {"time":"2024-08-25T16:50:07.000012234+08:00","level":"DEBUG","source":{"function":"github.com/dushaoshuai/go-usage-examples/golang/log/slog_test.Example_json_handler","file":"/tmp/go-usage-examples/golang/log/slog/slog_test.go","line":28},"msg":"a debug level msg","x1":1,"x2":"x2 value"}
	// {"time":"2024-08-25T16:50:07.000012234+08:00","level":"WARN","source":{"function":"github.com/dushaoshuai/go-usage-examples/golang/log/slog_test.Example_json_handler","file":"/tmp/go-usage-examples/golang/log/slog/slog_test.go","line":30},"msg":"a warn level msg","x1":1,"x2":2}
	// {"time":"2024-08-25T16:50:07.000012234+08:00","level":"WARN","source":{"function":"github.com/dushaoshuai/go-usage-examples/golang/log/slog_test.Example_json_handler","file":"/tmp/go-usage-examples/golang/log/slog/slog_test.go","line":31},"msg":"a warn level msg","x1":1,"x2":2}
}
