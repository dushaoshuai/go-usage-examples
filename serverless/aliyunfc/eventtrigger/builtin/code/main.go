package main

import (
	"context"
	"github.com/aliyun/fc-runtime-go-sdk/fc"
)

type Req struct {
	Name       string `json:"name"`
	Age        int8   `json:"age"`
	Department string `json:"department"`
}

type Resp = Req

func HandleRequest(_ context.Context, event Req) (Resp, error) {
	return event, nil
}

func main() {
	fc.Start(HandleRequest)
}
