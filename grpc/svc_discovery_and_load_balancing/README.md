k8s 服务发现、gRPC 域名解析、gRPC 负载均衡

# gRPC 负载均衡问题

gRPC 基于 HTTP/2。

对比 HTTP/2 和 HTTP/1.1，两者都使用 TCP 长连接。因为 HTTP/2 有多路复用的特性，使用 HTTP/2 只需要建立一个 TCP 连接。而使用 HTTP/1.1 时，需要建立多个 TCP 连接才能进行并发请求。

HTTP/2 使用 TCP 长连接和多路复用的特性，导致使用 ClusterIP 类型的 k8s Service 时，gRPC 客户端通过 Service 的负载均衡（Round Robin）和 Sericve 所代理的某个 Pod 建立长连接，失去了负载均衡的作用。

为了解决这个问题，gRPC 客户端需要和所有的 Pod 都建立连接，在这些连接间进行请求级别的负载均衡。

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
