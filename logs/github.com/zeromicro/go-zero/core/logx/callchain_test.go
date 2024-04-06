package logx_test

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/dushaoshuai/go-usage-examples/golang/runtime/funcs"
)

const (
	chainKey = "call_chain"
)

type contextKey struct{}

var callChainKey contextKey

func withCallChain(ctx context.Context) context.Context {
	chain, ok := ctx.Value(callChainKey).([]string)
	if !ok {
		chain = []string{funcs.GetCallerNameSkip(1)}
	} else {
		chain = append(chain, funcs.GetCallerNameSkip(1))
	}
	return context.WithValue(ctx, callChainKey, chain)
}

func callChainFromContext(ctx context.Context) logx.LogField {
	chain, ok := ctx.Value(callChainKey).([]string)
	if !ok {
		chain = []string{funcs.GetCallerNameSkip(1)}
	} else {
		chain = append(chain, funcs.GetCallerNameSkip(1))
	}
	return logx.Field(chainKey, chain)
}
