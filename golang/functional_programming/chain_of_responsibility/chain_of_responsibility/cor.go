package chain_of_responsibility

import (
	"context"
)

// Handler defines the handler invoked by Interceptor to complete the normal execution.
type Handler func(ctx context.Context, req any) (resp any, err error)

// Interceptor provides a hook to intercept the execution.
// It is the responsibility of the interceptor to invoke handler to complete the execution.
type Interceptor func(ctx context.Context, req any, handler Handler) (resp any, err error)

// chainInterceptors chains interceptors into one. The first interceptor will be the outer most,
// while the last interceptor will be the inner most wrapper around handler.
func chainInterceptors(interceptors []Interceptor) Interceptor {
	return func(ctx context.Context, req any, handler Handler) (resp any, err error) {
		return interceptors[0](ctx, req, getChainedHandler(interceptors, 0, handler))
	}
}

func getChainedHandler(interceptors []Interceptor, curr int, finalHandler Handler) Handler {
	if curr == len(interceptors)-1 {
		return finalHandler
	}
	return func(ctx context.Context, req any) (resp any, err error) {
		return interceptors[curr+1](ctx, req, getChainedHandler(interceptors, curr+1, finalHandler))
	}
}

// 若干 `Interceptor` 和 `Handler` 串成了一条责任链。请求经过 `Interceptor` 的拦截处理，最后由 `Handler` 完成真正的响应。
// 每个 `Interceptor` 都可以决定是否调用下一个 `Interceptor`，最后一个 `Interceptor` 可以决定是否调用 `Handler`。
//
// `chainInterceptors` 将多个 `Interceptor` 组合成一个 `Interceptor`，这是函数式编程-装饰模式的写法。
// 可以理解为前一个 `Interceptor` 修饰（包裹）后一个 `Interceptor`，变成一个 `Interceptor`。
// 所有的 `Interceptor` 层层包裹，第一个 `Interceptor` 在最外面，最后一个在最里面，这就是洋葱模型。
// 请求到达时，先由外向内进入，再由内向外返回。
