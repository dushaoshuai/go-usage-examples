package functional_options

// 此文件中的代码是从 google.golang.org/grpc
// https://github.com/grpc/grpc-go/blob/v1.53.0/server.go 中抽离出来的,
// 主要是为了学习 函数式选项 模式. 只保留了一些我认为重要的代码.

import (
	"math"
	"time"
)

const (
	defaultServerMaxReceiveMessageSize = 1024 * 1024 * 4
	defaultServerMaxSendMessageSize    = math.MaxInt32
	defaultWriteBufSize                = 32 * 1024
	defaultReadBufSize                 = 32 * 1024
)

type Server struct {
	opts serverOptions
	// ...
}

func NewServer(opt ...ServerOption) *Server {
	opts := defaultServerOptions
	for _, o := range opt {
		o.apply(&opts)
	}
	s := &Server{
		opts: opts,
	}
	// ...
	return s
}

type serverOptions struct {
	maxReceiveMessageSize int
	maxSendMessageSize    int
	writeBufferSize       int
	readBufferSize        int
	connectionTimeout     time.Duration
	initialWindowSize     int32
	initialConnWindowSize int32
}

var defaultServerOptions = serverOptions{
	maxReceiveMessageSize: defaultServerMaxReceiveMessageSize,
	maxSendMessageSize:    defaultServerMaxSendMessageSize,
	connectionTimeout:     120 * time.Second,
	writeBufferSize:       defaultWriteBufSize,
	readBufferSize:        defaultReadBufSize,
}

type ServerOption interface {
	apply(*serverOptions)
}

// funcServerOption wraps a function that modifies serverOptions into an
// implementation of the ServerOption interface.
type funcServerOption struct {
	f func(options *serverOptions)
}

func (fso *funcServerOption) apply(so *serverOptions) {
	fso.f(so)
}

func newFuncServerOption(f func(*serverOptions)) ServerOption {
	return &funcServerOption{
		f: f,
	}
}

func WriteBufferSize(s int) ServerOption {
	return newFuncServerOption(func(o *serverOptions) {
		o.writeBufferSize = s
	})
}

func ReadBufferSize(s int) ServerOption {
	return newFuncServerOption(func(o *serverOptions) {
		o.readBufferSize = s
	})
}

func InitialWindowSize(s int32) ServerOption {
	return newFuncServerOption(func(o *serverOptions) {
		o.initialWindowSize = s
	})
}

func InitialConnWindowSize(s int32) ServerOption {
	return newFuncServerOption(func(o *serverOptions) {
		o.initialConnWindowSize = s
	})
}

// ...
