k8s 服务发现、gRPC 域名解析、gRPC 负载均衡

# gRPC 负载均衡问题

gRPC 基于 HTTP/2。

对比 HTTP/2 和 HTTP/1.1，两者都使用 TCP 长连接。因为 HTTP/2 有多路复用的特性，使用 HTTP/2 只需要建立一个 TCP 连接。而使用 HTTP/1.1 时，需要建立多个 TCP 连接才能进行并发请求。

HTTP/2 使用 TCP 长连接和多路复用的特性，导致使用 ClusterIP 类型的 k8s Service 时，gRPC 客户端通过 Service 的负载均衡（Round Robin）和 Sericve 所代理的某个 Pod 建立长连接，失去了负载均衡的作用。

为了解决这个问题，gRPC 客户端需要和所有的 Pod 都建立连接，在这些连接间进行请求级别的负载均衡。

# gRPC 域名解析

gRPC 域名解析和服务发现有关。gRPC 提供了若干种名字解析机制[^1][^2]：

[^1]: https://grpc.io/docs/guides/custom-name-resolution/#:~:text=various%20other%20name%20resolution%20mechanisms
[^2]: https://github.com/grpc/grpc/blob/master/doc/naming.md

* DNS（默认）：service/name 或者 dns:///service/name
* Unix Domain Socket：uds:///run/containerd/containerd.sock
* xDS：xds:///wallet.grpcwallet.io
* IPv4：ipv4:198.51.100.123:50051
* [自定义解析器](https://grpc.io/docs/guides/custom-name-resolution/#custom-name-resolvers)

解析无 scheme 的服务名（如 my-service）时，grpc 客户端默认使用 DNS，否则，根据 scheme 选择对应的域名解析器。

用户可以自定义解析器，例如，scheme 为 `my-resolver` 的解析器，可以解析以 `my-resolver:` 开头的服务名，如 `my-resolver:///my-service`。使用自定义解析器可以基于 watch 持续更新服务端信息。

# gRPC 自定义域名解析器

自定义解析器的机制详见[这里](https://grpc.io/docs/guides/custom-name-resolution/#life-of-a-target-string)。使用 [grpc-go](https://pkg.go.dev/google.golang.org/grpc) 时，需要使用 [resolver](https://pkg.go.dev/google.golang.org/grpc@v1.57.0/resolver) 包：

1. 使用 resolver.Register() 函数注册一个自定义的 resolver.Builder 
2. resolver.Builder 的 Build() 方法返回一个自定义的 resolver.Resolver
3. resolver.Resolver watch 目标服务的更新，包括地址更新和 service config 更新

# 客户端负载均衡

## DNS 服务发现

* 客户端负载均衡
* k8s Headless Service
* DNS 解析器

https://github.com/grpc/grpc-go/issues/3170#issuecomment-552517779

## k8s API 服务发现

// TODO.

# Proxy 负载均衡（也叫服务端负载均衡）

// TODO.

# 参见

* [gRPC Load Balancing](https://grpc.io/blog/grpc-load-balancing/#proxy-or-client-side)
* [gRPC Load Balancing on Kubernetes without Tears](https://kubernetes.io/blog/2018/11/07/grpc-load-balancing-on-kubernetes-without-tears/)
* [gRPC on HTTP/2 Engineering a Robust, High-performance Protocol](https://grpc.io/blog/grpc-on-http2/#resolvers-and-load-balancers)
* [grpc-by-example-java/kubernetes-lb-example](https://github.com/saturnism/grpc-by-example-java/tree/master/kubernetes-lb-example)
* [Custom Name Resolution](https://grpc.io/docs/guides/custom-name-resolution/)
* [gRPC Name Resolution](https://github.com/grpc/grpc/blob/master/doc/naming.md)
* [Custom Load Balancing Policies](https://grpc.io/docs/guides/custom-load-balancing/)
* [Load Balancing in gRPC](https://github.com/grpc/grpc/blob/master/doc/load-balancing.md)
* [Headless Services](https://kubernetes.io/docs/concepts/services-networking/service/#headless-services)
* [k8s 服务发现 以及 gRPC 长连接负载均衡](https://segmentfault.com/a/1190000039361024)
* [K8S容器编排之Headless浅谈](https://zhuanlan.zhihu.com/p/54153164)
* [grpc-go/examples/features/name_resolving/](https://github.com/grpc/grpc-go/tree/master/examples/features/name_resolving)


* [grpc Guides](https://grpc.io/docs/guides/)
* [grpc Blog](https://grpc.io/blog/)
* [grpc-go/examples](https://github.com/grpc/grpc-go/tree/master/examples)
* [grpc-go/Documentation](https://github.com/grpc/grpc-go/tree/master/Documentation)
