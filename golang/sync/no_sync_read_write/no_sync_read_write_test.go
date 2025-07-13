package no_sync_read_write

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/dushaoshuai/go-usage-examples/golang/sync/xsync/errgroup"
)

// https://go.dev/ref/mem#restrictions
// Implementation Restrictions for Programs Containing Data Races
// A read r of a memory location x holding a value that is not larger than
// a machine word must observe some write w such that r does not happen before w
// and there is no write w' such that w happens before w' and w' happens before r.
// That is, each read must observe a value written by a preceding or concurrent write.

func Test_no_sync_read_write(t *testing.T) {
	var v int
	gs := 1 << 14

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	eg, ctx := errgroup.WithContext(ctx)

	for i := range gs {
		eg.Go(func() error {
			for {
				select {
				case <-ctx.Done():
					return nil
				default:
					v = i
				}
			}
		})
	}

	for range gs {
		eg.Go(func() error {
			for {
				select {
				case <-ctx.Done():
					return nil
				default:
					vv := v
					require.GreaterOrEqual(t, vv, 0)
					require.LessOrEqual(t, vv, gs-1)
				}
			}
		})
	}

	_ = eg.Wait()
}
