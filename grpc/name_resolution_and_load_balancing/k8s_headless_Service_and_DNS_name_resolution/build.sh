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

client_image=registry.cn-zhangjiakou.aliyuncs.com/shaouai/grpc-dns_discovery-k8s_headless_svc-lb-client:${image_tag}
docker build -t "${client_image}" --build-arg=BIN=${client_bin} -f ./Dockerfile .
docker push "${client_image}"

server_image=registry.cn-zhangjiakou.aliyuncs.com/shaouai/grpc-dns_discovery-k8s_headless_svc-lb-server:${image_tag}
docker build -t "${server_image}" --build-arg=BIN=${server_bin} -f ./Dockerfile .
docker push "${server_image}"

rm ./${client_bin} ./${server_bin}
