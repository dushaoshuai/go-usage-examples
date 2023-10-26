HTTP 请求处理程序（HTTP Handler）

# 基本信息

请求处理程序分为事件请求处理程序（Event Handler）和 HTTP 请求处理程序（HTTP Handler），其中事件请求由各种事件源触发生成，HTTP
请求则由 HTTP 触发器触发生成。

本文关注 HTTP Handler。编程语言使用 Go。

# HTTP Handler

## 内置运行时

这是一个使用内置运行时 go1 实现的 HTTP Handler。

项目结构如下，其中 main.go 实现了 handler，s.yaml 描述云函数的资源、行为：

```shell
$ tree .
.
├── code
│   └── main.go
└── s.yaml

2 directories, 2 files
```

s.yaml 内容如下:

```yaml
edition: 1.0.0 # 命令行YAML规范版本，遵循语义化版本（Semantic Versioning）规范
name: http-handler-builtin-runtime-example # 项目/应用名称
access: default # 密钥别名

services: # 应用所包含的服务，可以包含多个
  http-handler-builtin-runtime-example-service: # 服务/模块名称
    component: devsapp/fc # 组件名称，这里使用阿里云函数计算（FC）组件
    actions: # 自定义执行逻辑
      pre-deploy: # 在 deploy 之前运行，把 main.go 编译为可执行的二进制文件
        - run: go mod tidy
          path: ./code
        - run: GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go
          path: ./code
    props: # 组件的属性值
      region: cn-zhangjiakou # 地域
      service: # 服务配置
        name: http-handler-builtin-runtime-example-service # service 名称
        description: "http handler builtin runtime example service" # Service 的简短描述
        internetAccess: false # 设为 true 让 function 可以访问公网
        tracingConfig: Disable # 链路追踪，可取值：Enable、Disable
        role: acs:ram::1810657881264284:role/aliyunfcdefaultrole # 授予函数计算所需权限的RAM role
        logConfig: null # log配置，function产生的log会写入这里配置的logstore
        vpcConfig: null # VPC配置, 配置后function可以访问指定VPC
        nasConfig: null # NAS配置, 配置后function可以访问指定NAS
        ossMountConfig: null # OSS挂载配置, 配置后function可以访问指定OSS bucket
        vpcBinding: null # 仅允许指定 VPC 调用函数
      function: # 函数配置
        name: http-handler-builtin-runtime-example-function # function 名称
        description: "http handler builtin runtime example function" # function 的简短描述
        codeUri: ./code # 代码位置
        handler: main # function 执行的入口，具体格式和语言相关
        memorySize: 128 # function 的内存规格
        runtime: go1 # 运行时
        timeout: 10 # function 运行的超时时间
        cpu: 0.05 # 函数的 CPU 规格，单位为 vCPU，为 0.05 vCPU 的倍数
        diskSize: 512 # 函数的磁盘规格，单位为 MB，可选值为 512 MB 或 10240 MB
        instanceConcurrency: 10 # 单实例多并发，一个函数实例可以并发处理这么多请求
        instanceSoftConcurrency: 7 # 扩容并发度。扩容并发度用于优雅扩容，
          # 当实例上并发数超过扩容并发度时，会触发实例扩容。
          # 例如，您的实例启动较慢，可以通过设置合适的扩容并发度提前启动实例。
          # 注意：扩容并发度的值不能大于实例并发度，最小值为1。
        # 线上存在此配置，但是yaml中没有配置，则默认为和 instanceConcurrency 值一致。
        instanceType: e1 # 函数实例类型，可选值为：e1（弹性实例）、c1（性能实例）、fc.gpu.tesla.1（GPU T4实例）、fc.gpu.ampere.1（GPU A10实例）。
        environmentVariables: # 环境变量
          TZ: "Asia/Shanghai" # 设置时区为东 8 区
      triggers: # 触发器配置
        - name: httptrigger # 触发器名称
          type: http # 触发器类型
          qualifier: LATEST # 触发器函数的版本或者别名，默认 LATEST
          config: # 触发器配置
            authType: anonymous # 鉴权类型，可选值：anonymous、function
            disableURLInternet: false # 是否禁用公网访问 URL，默认为 false
            methods: # HTTP 触发器支持的访问方法，可选值：GET、POST、PUT、DELETE、PATCH、HEAD、OPTIONS
              - GET
              - POST
      customDomains: # 自定义域名
        - domainName: auto # 域名，如果是 auto 取值，系统则会默认分配域名
          protocol: HTTP # 协议，取值：HTTP, HTTP,HTTPS
```

main.go 内容如下:

```go
package main

import (
	"context"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
)

func HandleHttpRequest(_ context.Context, w http.ResponseWriter, req *http.Request) error {
	resp := []string{
		time.Now().Format(time.DateTime),
		"Request Method: " + req.Method,
	}
	w.Write([]byte(strings.Join(resp, "\n")))
	w.Write([]byte{'\n'})

	body, err := io.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "text/plain")
		w.Write([]byte(err.Error()))
		return nil
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/plain")
	w.Write(body)
	return nil
}

func main() {
	fc.StartHttp(HandleHttpRequest)
}
```

### 本地调试

可以使用 `s local start` 命令把 HTTP 函数部署在本地，方便调试。但是此命令并不会执行 pre-deploy 中的 action，把 Go
程序编译为二进制可执行文件。

因为使用 Go 语言时的交付物是一个二进制可执行文件，因此要先编译。我尝试了执行 `s build` 命令，发现其也不会执行 pre-deploy 中的
action，似乎 Go 语言并不在 `s build` 命令的考虑中。这里只能自己编译。

```shell
$ GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o code/main code/main.go
$ s local start
✔ devsapp_fc-local-invoke.zip file decompression completed
[2023-10-26 23:18:36] [INFO] [FC-LOCAL-INVOKE] - Using trigger for start: 
name: httptrigger
type: http
qualifier: LATEST
config:
  authType: anonymous
  disableURLInternet: false
  methods:
    - GET
    - POST

The local command for go1 runtime is in public test. If you have any questions, welcome to join DingTalk Group: 33947367
[2023-10-26 23:18:36] [INFO] [FC-LOCAL-INVOKE] - CustomDomain auto of http-handler-builtin-runtime-example-service/http-handler-builtin-runtime-example-function was registered
        url: http://localhost:7648/
        methods: GET,POST
        authType: anonymous

Tips for next step
======================
* Deploy Resources: s deploy
http-handler-builtin-runtime-example-service: 
  status: succeed
function compute app listening on port 7648!
```

使用 `curl` 对命令进行调试：

## 自定义运行时

TBD

## 自定义容器运行时

TBD

# 参见

* [函数类型选型](https://help.aliyun.com/zh/fc/product-overview/overview-30)
* [请求处理程序（Handler）](https://help.aliyun.com/zh/fc/user-guide/handlers-1?)
* [HTTP请求处理程序（HTTP Handler）](https://help.aliyun.com/zh/fc/http-handlers-3)
* [代码开发 - Go](https://help.aliyun.com/zh/fc/user-guide/go-1)
* [Serverless Devs 描述文件（Yaml）规范](https://docs.serverless-devs.com/serverless-devs/yaml)
* [ 函数计算（FC）组件 Yaml 规范](https://docs.serverless-devs.com/fc/yaml/readme)
