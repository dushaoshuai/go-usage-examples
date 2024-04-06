package logx_test

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

func Example_use_ctx() {
	logx.AddGlobalFields( // global fields, added in all logs
		logx.Field("library", "logx"),
		logx.Field("purpose", "learn and testing"),
	)

	ctx := context.Background()
	ctx = logx.ContextWithFields(ctx, // fields propagated via context
		logx.Field("log_key_in_ctx", "log_val_in_ctx"),
		logx.Field("log_key2_in_ctx", "log_val2_in_ctx"),
	)

	logWithContext(ctx)

	// Output:
	// {"@timestamp":"2024-04-06T17:30:37.849+08:00","a":1,"b":2,"c":3,"caller":"logx/logx_test.go:31","content":"test logx","duration":"0.0ms","level":"info","library":"logx","log_key2_in_ctx":"log_val2_in_ctx","log_key_in_ctx":"log_val_in_ctx","purpose":"learn and testing"}
}

func logWithContext(ctx context.Context) {
	logger := logx.WithContext(ctx) // use fields in context
	defer func(start time.Time) {
		logger.WithDuration(time.Now().Sub(start)).
			Info("test logx") // default is json encoding
	}(time.Now())

	logger.WithFields(
		logx.Field("a", 1),
		logx.Field("b", 2),
		logx.Field("c", 3),
	)
}
