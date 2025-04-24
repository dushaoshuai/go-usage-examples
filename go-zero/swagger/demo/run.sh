#!/usr/bin/env bash

# goctl -v
# goctl version 1.8.3 linux/amd64

goctl api go --api ./swagger.api --dir . --style go_zero
GOCTL_EXPERIMENTAL=on goctl api swagger --api ./swagger.api --dir ./internal/handler/swagger-ui-5.21.0/dist
go run .