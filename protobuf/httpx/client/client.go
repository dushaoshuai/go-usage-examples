package main

import (
	"io"
	"log/slog"
	"net/http"
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/dushaoshuai/go-usage-examples/protobuf/httpx"
	protox "github.com/dushaoshuai/go-usage-examples/protobuf/httpx/proto"
)

func main() {
	for range time.Tick(3 * time.Second) {
		resp, err := http.Get("http://127.0.0.1" + httpx.Port)
		if err != nil {
			slog.Error(err.Error())
			continue
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			slog.Error(err.Error())
			continue
		}
		defer resp.Body.Close()

		var book protox.AddressBook
		err = proto.Unmarshal(body, &book)
		if err != nil {
			slog.Error(err.Error())
			continue
		}

		slog.Info("succeed",
			slog.String("book", book.String()),
		)
	}
}
