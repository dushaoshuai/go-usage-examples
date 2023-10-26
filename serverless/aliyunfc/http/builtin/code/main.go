package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
)

func HandleHttpRequest(_ context.Context, w http.ResponseWriter, req *http.Request) error {
	t := time.Now().Format(time.DateTime)
	t += "\n"

	body, err := io.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "text/plain")
		w.Write([]byte(t + err.Error()))
		return nil
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte(t + fmt.Sprintf("Hi，%s!\n", body)))
	return nil
}

func main() {
	fc.StartHttp(HandleHttpRequest)
}
