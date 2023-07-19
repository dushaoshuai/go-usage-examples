package redis_test

import (
	"testing"

	"github.com/redis/go-redis/v9"
)

// Redis transactions:
//
// https://redis.io/docs/interact/transactions/
// https://redis.uptrace.dev/guide/go-redis-pipelines.html#transactions
//
// Implement INCR:
//
// WATCH mykey
// val = GET mykey
// val = val + 1
// MULTI
// SET mykey $val
// EXEC

func Test_transaction(t *testing.T) {
	ctx, cancel := defaultCtx()
	defer cancel()

	key := "key"
	db2SetVal := 100

	db1 := mustNewDB(ctx, 14)
	db2 := mustNewDB(ctx, 14)
	defer func() {
		db1.FlushDB(ctx) // ignore errors
	}()

	// this transaction should fail
	txf := func(tx *redis.Tx) error {
		n, err := tx.Get(ctx, key).Int()
		if err != nil && err != redis.Nil {
			return err
		}
		n++

		// another client changes the key
		db2.Set(ctx, key, db2SetVal, 0) // ignore errors to make things clear

		_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			pipe.Set(ctx, key, n, 0)
			return nil
		})
		return err
	}

	// this transaction should succeed
	txf2 := func(tx *redis.Tx) error {
		n, err := tx.Get(ctx, key).Int()
		if err != nil && err != redis.Nil {
			return err
		}
		n++
		_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			pipe.Set(ctx, key, n, 0)
			return nil
		})
		return err
	}

	err := db1.Watch(ctx, txf, key)
	if err != redis.TxFailedErr { // ignore other errors here
		t.Error("transaction should have failed")
	}
	err = db1.Watch(ctx, txf2, key)
	if err != nil {
		t.Error(err)
	}

	n, err := db1.Get(ctx, key).Int()
	if err != nil {
		t.Error(err)
	}
	if n != db2SetVal+1 {
		t.Errorf("n: got %d, want %d", n, db2SetVal+1)
	}
}
