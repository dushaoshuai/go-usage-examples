#!/usr/bin/env bash

set -e

go mod tidy -x # go help mod

go get -x go@latest # go help get
# go mod edit -go=1.23.4 # go help mod edit

go get -x toolchain@latest # go help get
# go mod edit -toolchain=version # go help mod edit

go get -u -x ./...
go mod tidy -x
