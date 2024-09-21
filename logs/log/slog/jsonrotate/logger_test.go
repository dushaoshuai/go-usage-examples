package jsonrotate

import (
	"context"
	"log/slog"
	"testing"
	"time"
)

func Test_json_rotate_logger(t *testing.T) {
	ctx := context.Background()

	closeUnderlyingWriter := setJsonRotateLogger()
	defer closeUnderlyingWriter()

	for tim := range time.Tick(time.Second) {
		slog.LogAttrs(ctx, slog.LevelInfo, "hello",
			slog.Int("count", 3),
			slog.Time("tim", tim),
		)
	}
}
