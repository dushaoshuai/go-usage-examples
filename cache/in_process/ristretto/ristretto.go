package ristretto

import (
	"github.com/dgraph-io/ristretto/v2"
)

const (
	numCounters int64 = 1e7
	maxCost     int64 = 1e6
)

func newCache() (*ristretto.Cache[string, any], error) {
	return ristretto.NewCache(&ristretto.Config[string, any]{
		NumCounters:        numCounters,
		MaxCost:            maxCost,
		BufferItems:        64,
		Metrics:            true,
		IgnoreInternalCost: true,
	})
}
