package cor

import (
	"fmt"
)

type Req struct {
	n int
}

func NewReq(n int) *Req {
	return &Req{n: n}
}

type nextHandler struct {
	next Handler
}

func (n *nextHandler) SetNext(h Handler) { n.next = h }
func (n *nextHandler) do(r *Req) {
	if n.next != nil {
		n.next.Do(r)
	}
}

type Handler interface {
	Do(*Req)
	SetNext(Handler)
}

type Handler1 struct {
	nextHandler
}

func NewHandler1() Handler {
	return &Handler1{}
}

func (h *Handler1) Do(r *Req) {
	fmt.Printf("handler1, req.n = %d\n", r.n)
	r.n++
	h.nextHandler.do(r)
}

type Handler2 struct {
	nextHandler
}

func NewHandler2() Handler {
	return &Handler2{}
}

func (h *Handler2) Do(r *Req) {
	fmt.Printf("handler2, req.n = %d\n", r.n)
	r.n++
	h.nextHandler.do(r)
}

type Handler3 struct {
	nextHandler
}

func NewHandler3() Handler {
	return &Handler3{}
}

func (h *Handler3) Do(r *Req) {
	fmt.Printf("handler3, req.n = %d\n", r.n)
	if r.n >= 5 {
		fmt.Printf("handler3, req.n >= 5, stop")
		return
	}
	r.n++
	h.nextHandler.do(r)
}
