package jsonrotate

import (
	"context"
	"log/slog"
	"testing"
	"time"
)

func Test_json_logger(t *testing.T) {
	ctx := context.Background()

	for tim := range time.Tick(time.Second) {
		slog.LogAttrs(ctx, slog.LevelInfo, "hello",
			slog.Int("count", 3),
			slog.Time("tim", tim),
		)
	}
}

func TestMain(m *testing.M) {
	setJsonRotateLogger()

	m.Run()
}
