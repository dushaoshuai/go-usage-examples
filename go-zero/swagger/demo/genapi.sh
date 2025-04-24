#!/usr/bin/env bash

GOCTL_EXPERIMENTAL=on goctl api swagger --api ./demo.api --dir ./internal/handler/swagger-ui-5.21.0/dist