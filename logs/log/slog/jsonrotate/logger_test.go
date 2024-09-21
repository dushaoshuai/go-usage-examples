package jsonrotate

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"testing"
)

func TestMain(m *testing.M) {
	closeUnderlyingWriter := setJsonRotateLogger()
	defer closeUnderlyingWriter()

	m.Run()
}

func Test_json_logger(t *testing.T) {
	ctx := context.Background()

	for i := 9; i >= 0; i-- {
		// time.Sleep(2 * time.Second)
		fmt.Println(i)
		slog.LogAttrs(ctx, slog.LevelInfo, "hello today",
			slog.Int("int", i),
			slog.String("str", strconv.Itoa(i)),
		)
	}
}

// func Test_rotate_logger(t *testing.T) {
// 	tomorrow := time.Now().AddDate(0, 0, 1)
// 	mock.Patch(time.Now, func() time.Time {
// 		return tomorrow
// 	})
//
// 	ctx := context.Background()
//
// 	for i := range 10 {
// 		slog.LogAttrs(ctx, slog.LevelWarn, "hello tomorrow",
// 			slog.Int("count", i),
// 		)
// 	}
// }
