package jsonrotate

import (
	"context"
	"log/slog"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/xhd2015/xgo/runtime/mock"
)

func TestMain(m *testing.M) {
	logger, err := NewJSONRotateLogger("/tmp/JSONRotateLogger/x.log")
	if err != nil {
		os.Exit(1)
	}
	defer logger.Close()

	slog.SetDefault(logger.Slogger())

	m.Run()
}

func Test_json_logger(t *testing.T) {
	ctx := context.Background()

	for i := range 10 {
		slog.LogAttrs(ctx, slog.LevelInfo, "hello today",
			slog.Int("int", i),
			slog.String("str", strconv.Itoa(i)),
		)
	}
}

func Test_rotate_logger(t *testing.T) {
	// The actual write operation and rotation happen in a separate goroutine,
	// so the patch here can't affect them (https://github.com/xhd2015/xgo/tree/master/runtime/mock#:~:text=the%20first%20parameter-,Scope,-Based%20on%20the),
	// making unit testing quite tricky.
	// As a result, this unit test is not successful.
	tomorrow := time.Now().AddDate(0, 0, 1)
	mock.Patch(time.Now, func() time.Time {
		return tomorrow
	})

	ctx := context.Background()

	for i := range 10 {
		slog.LogAttrs(ctx, slog.LevelWarn, "hello tomorrow",
			slog.Int("int", i),
			slog.String("str", strconv.Itoa(i)),
		)
	}
}
