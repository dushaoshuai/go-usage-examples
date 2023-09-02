k8s 服务发现和 gRPC 负载均衡

# gRPC 负载均衡的问题

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
