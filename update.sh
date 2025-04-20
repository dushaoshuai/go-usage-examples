#!/usr/bin/env bash

set -e

go mod tidy -x # go help mod

go get -x go@latest # go help get
# go mod edit -go=1.23.4 # go help mod edit

go get -x toolchain@latest # go help get
# go mod edit -toolchain=version # go help mod edit

# The new tool meta-pattern refers to all tools in the current module.
# This can be used to upgrade them all with go get tool or to install them into your GOBIN directory with go install tool.
go get -x tool

go get -u -x ./...
go mod tidy -x
