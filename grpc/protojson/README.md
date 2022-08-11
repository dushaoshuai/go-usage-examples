# 目的

观察 google.golang.org/protobuf/encoding/protojson 对 
google.golang.org/protobuf/proto:Message 的 encoding 和 decoding 操作

# 结论

即使 message 中有 oneof 类型的字段，protojson 也可以正确处理

# 链接

[protojson](https://pkg.go.dev/google.golang.org/protobuf@v1.28.1/encoding/protojson)

[proto.Message](https://pkg.go.dev/google.golang.org/protobuf@v1.28.1/proto#Message)
