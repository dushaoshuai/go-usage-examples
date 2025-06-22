package main

import (
	"io"
	"log/slog"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8085")
	if err != nil {
		slog.Error(err.Error())
		return
	}
	defer func() {
		conn.Close()
		slog.Info("connection closed")
	}()
	slog.Info("connecting to server",
		slog.String("remote_addr", conn.RemoteAddr().String()),
	)

	buf := make([]byte, 12)

	for {
		err = conn.SetDeadline(time.Now().Add(time.Second))
		if err != nil {
			slog.Error(err.Error())
			return
		}

		_, err = conn.Write([]byte("hello world!"))
		if err != nil {
			slog.Error(err.Error())
			return
		}

		_, err = io.ReadFull(conn, buf)
		if err != nil {
			slog.Error(err.Error())
			return
		} else {
			slog.Info("read message", slog.String("msg", string(buf)))
		}
	}
}
