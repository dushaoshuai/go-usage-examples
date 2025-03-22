package client

import (
	"time"

	"github.com/dgraph-io/ristretto/v2"
	"github.com/samber/lo"
)

var (
	c *ristretto.Cache[string, any]
)

func init() {
	c = lo.Must(newCache())
}

const (
	defaultCost = 1

	maxCost     int64 = 1e6
	numCounters       = maxCost * 10
)

func newCache() (*ristretto.Cache[string, any], error) {
	return ristretto.NewCache(&ristretto.Config[string, any]{
		NumCounters:        numCounters,
		MaxCost:            maxCost,
		BufferItems:        64,
		IgnoreInternalCost: true,
	})
}

func Set(key string, value any) bool {
	return c.Set(key, value, defaultCost)
}

func SetWithTTL(key string, value any, ttl time.Duration) bool {
	return c.SetWithTTL(key, value, defaultCost, ttl)
}

func Get[T any](key string) (T, bool) {
	get, ok := c.Get(key)
	if !ok {
		var zeroT T
		return zeroT, false
	}
	return get.(T), true
}

func GetFunc[T any](key string, fn func() (T, error)) (T, error) {
	get, ok := c.Get(key)
	if ok {
		return get.(T), nil
	}

	val, err := fn()
	if err != nil {
		var zeroT T
		return zeroT, err
	}

	c.Set(key, val, defaultCost)
	return val, nil
}

func GetFuncWithTTL[T any](key string, ttl time.Duration, fn func() (T, error)) (T, error) {
	get, ok := c.Get(key)
	if ok {
		return get.(T), nil
	}

	val, err := fn()
	if err != nil {
		var zeroT T
		return zeroT, err
	}

	c.SetWithTTL(key, val, defaultCost, ttl)
	return val, nil
}
