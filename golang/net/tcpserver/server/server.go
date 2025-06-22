package main

import (
	"errors"
	"io"
	"log/slog"
	"net"
	"os"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", ":8085")
	if err != nil {
		slog.Error(err.Error())
		return
	}
	defer ln.Close()
	slog.Info("tcp server started")

	for {
		conn, err := ln.Accept()
		if err != nil {
			slog.Error(err.Error())
			continue
		}

		slog.Info("accepted connection",
			slog.String("remote_addr", conn.RemoteAddr().String()),
		)
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func() {
		conn.Close()
		slog.Info("connection closed")
	}()

	buf := make([]byte, 12)

	for {
		err := conn.SetDeadline(time.Now().Add(time.Second))
		if err != nil {
			slog.Error(err.Error())
			return
		}

		_, err = io.ReadFull(conn, buf)
		if err != nil {
			// double check timeout errors
			if errors.Is(err, os.ErrDeadlineExceeded) {
				slog.Info(err.Error())
				continue
			}
			if te, ok := err.(interface{ Timeout() bool }); ok && te.Timeout() {
				slog.Info(err.Error())
				continue
			}

			if te, ok := err.(interface{ Temporary() bool }); ok && te.Temporary() {
				slog.Info(err.Error())
				continue
			} else {
				slog.Error(err.Error())
				return
			}
		} else {
			slog.Info("read message", slog.String("msg", string(buf)))
		}

		_, err = conn.Write(buf)
		if err != nil {
			// double check timeout errors
			if errors.Is(err, os.ErrDeadlineExceeded) {
				slog.Info(err.Error())
				continue
			}
			if te, ok := err.(interface{ Timeout() bool }); ok && te.Timeout() {
				slog.Info(err.Error())
				continue
			}

			if te, ok := err.(interface{ Temporary() bool }); ok && te.Temporary() {
				slog.Info(err.Error())
				continue
			} else {
				slog.Error(err.Error())
				return
			}
		}
	}
}
