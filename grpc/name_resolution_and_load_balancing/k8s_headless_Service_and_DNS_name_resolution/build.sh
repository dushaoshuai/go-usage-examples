#!/usr/bin/env bash

GOPROXY="https://goproxy.cn,direct"
GOFLAGS=" -mod=readonly -trimpath"
CGO_ENABLED=0
GO111MODULE="on"

go build -o client ./client/*.go
go build -o server ./server/*.go
