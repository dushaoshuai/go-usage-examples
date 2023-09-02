#!/usr/bin/env bash

GOPROXY="https://goproxy.cn,direct"
GOFLAGS=" -mod=readonly -trimpath"
CGO_ENABLED=0
GO111MODULE="on"
GOMAXPROCS=4
GOOS=linux

client_bin="client_bin"
server_bin="server_bin"

go build -o ${client_bin} ./client
go build -o ${server_bin} ./server

commit_id=$(git rev-parse HEAD)
image_tag=git.${commit_id}

function docker_build_and_push() {
  image=registry.cn-zhangjiakou.aliyuncs.com/shaouai/${1}:${image_tag}
  docker build -t "${image}" --build-arg=BIN="${2}" -f ./Dockerfile .
  docker push "${image}"
}

docker_build_and_push grpc-client-side-lb-client ${client_bin}
docker_build_and_push grpc-client-side-lb-server ${server_bin}

rm ./${client_bin} ./${server_bin}
