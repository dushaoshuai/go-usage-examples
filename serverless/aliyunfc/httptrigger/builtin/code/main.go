package main

import (
	"context"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
)

func HandleHttpRequest(_ context.Context, w http.ResponseWriter, req *http.Request) error {
	resp := []string{
		time.Now().Format(time.DateTime),
		"Request Method: " + req.Method,
	}
	w.Write([]byte(strings.Join(resp, "\n")))
	w.Write([]byte{'\n'})

	body, err := io.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "text/plain")
		w.Write([]byte(err.Error()))
		return nil
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/plain")
	w.Write(body)
	return nil
}

func main() {
	fc.StartHttp(HandleHttpRequest)
}
