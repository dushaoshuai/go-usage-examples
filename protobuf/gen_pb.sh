#!/usr/bin/env bash

SRC_DIR=./protobuf
DST_DIR=$SRC_DIR

protoc --proto_path=$SRC_DIR \
  --go_out=$DST_DIR --go_opt=paths=source_relative \
  --go-grpc_out=$DST_DIR --go-grpc_opt=paths=source_relative \
  $SRC_DIR/*.proto
