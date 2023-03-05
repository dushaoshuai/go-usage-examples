package decorator

func Decorate[T any](origin T, decorators ...func(T) T) T {
	for i := range decorators {
		d := decorators[len(decorators)-1-i]
		origin = d(origin)
	}
	return origin
}

type Languager func() string

var Go Languager = func() string { // 注意这里 Go 的类型必须是 Languager, 而不是 func() string,
	return "Go" // 否则范型机制通不过. 或者可以在调用 Decorate 时显式地将 Go 转换为 Languager 类型: Languager(Go)
}

func WithPython(l Languager) Languager {
	return func() string {
		return "Python " + l() + " Python"
	}
}

func WithC(l Languager) Languager {
	return func() string {
		return "C " + l() + " C"
	}
}
