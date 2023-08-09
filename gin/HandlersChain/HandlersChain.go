package HandlersChain

import "math"

type dummyContext struct {
	handlers HandlersChain
	index    int8

	Data int
}

type Context dummyContext

func NewContext(data int, middleware ...HandlerFunc) *Context {
	return &Context{
		handlers: middleware,
		index:    -1,
		Data:     data,
	}
}

// HandlerFunc defines the handler used by gin middleware as return value.
type HandlerFunc func(ctx *Context)

// HandlersChain defines a HandlerFunc slice.
type HandlersChain []HandlerFunc

// abortIndex represents a typical value used in abort functions.
const abortIndex int8 = math.MaxInt8 >> 1

/************************************/
/*********** FLOW CONTROL ***********/
/************************************/

// Next should be used only inside middleware.
// It executes the pending handlers in the chain inside the calling handler.
// See example in GitHub.
func (c *Context) Next() {
	c.index++
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
}

// IsAborted returns true if the current context was aborted.
func (c *Context) IsAborted() bool {
	return c.index >= abortIndex
}

// Abort prevents pending handlers from being called. Note that this will not stop the current handler.
// Let's say you have an authorization middleware that validates that the current request is authorized.
// If the authorization fails (ex: the password does not match), call Abort to ensure the remaining handlers
// for this request are not called.
func (c *Context) Abort() {
	c.index = abortIndex
}
