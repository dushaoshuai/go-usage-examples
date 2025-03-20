package httpx

//go:generate protoc --go_opt=paths=source_relative -I=./proto --go_out=./proto ./proto/proto.proto

const (
	Port = ":8085"
)
