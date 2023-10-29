事件请求处理程序（Event Handler）

# 基本信息

请求处理程序分为事件请求处理程序（Event Handler）和 HTTP 请求处理程序（HTTP Handler）；其中事件请求由各种事件源触发生成，HTTP 请求则由 HTTP 触发器触发生成。

本文关注 Event Handler。编程语言使用 Go。

# Event Handler

## 内置运行时

### 示例项目

这是一个使用内置运行时 go1 实现的 Event Handler。

项目结构如下，其中 main.go 实现了 handler，s.yaml 描述云函数的资源、行为：

```shell
tree .
.
├── code
│   ├── go.mod
│   ├── go.sum
│   └── main.go
└── s.yaml

2 directories, 4 files
```

s.yaml 内容如下:

```yaml
edition: 1.0.0 # 命令行YAML规范版本，遵循语义化版本（Semantic Versioning）规范
name: event-handler-builtin-runtime-example # 项目/应用名称
access: default # 密钥别名

services: # 应用所包含的服务，可以包含多个
  event-handler-builtin-runtime-example-service: # 服务/模块名称
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
        name: event-handler-builtin-runtime-example-service # service 名称
        description: "event handler builtin runtime example service" # Service 的简短描述
        internetAccess: false # 设为 true 让 function 可以访问公网
        tracingConfig: Disable # 链路追踪，可取值：Enable、Disable
        role: acs:ram::1810657881264284:role/aliyunfcdefaultrole # 授予函数计算所需权限的RAM role
        logConfig: null # log配置，function产生的log会写入这里配置的logstore
        vpcConfig: null # VPC配置, 配置后function可以访问指定VPC
        nasConfig: null # NAS配置, 配置后function可以访问指定NAS
        ossMountConfig: null # OSS挂载配置, 配置后function可以访问指定OSS bucket
        vpcBinding: null # 仅允许指定 VPC 调用函数
      function: # 函数配置
        name: event-handler-builtin-runtime-example-function # function 名称
        description: "event handler builtin runtime example function" # function 的简短描述
        codeUri: ./code # 代码位置，目录下的内容是最终的交付物
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
      triggers: null # 触发器配置
```

* 这里没有配置[触发器](https://help.aliyun.com/zh/fc/user-guide/event-triggers/)
* codeUri 目录中的内容是最终的交付物，函数计算最终会把此目录下的内容拷贝到容器 /code 目录下。使用 Go 语言时的交付物是一个二进制可执行文件，因此我们要确保这个二进制文件出现在 codeUri 指定的目录中，而且文件名是 handler 所指定的值。这里通过 pre-deploy 指定部署之前的 actions，在部署之前进行编译。

main.go 内容如下:

```go
package main

import (
	"context"
	"github.com/aliyun/fc-runtime-go-sdk/fc"
)

type Req struct {
	Name       string `json:"name"`
	Age        int8   `json:"age"`
	Department string `json:"department"`
}

type Resp = Req

func HandleRequest(_ context.Context, event Req) (Resp, error) {
	return event, nil
}

func main() {
	fc.Start(HandleRequest)
}
```

* handler 要遵循一定的签名，关于有效的签名，见 [Event Handler 签名](https://help.aliyun.com/zh/fc/event-handlers-6#section-jto-5ty-sh3)。
* 这里的 handler `HandleRequest` 的签名是 `func (context.Context, InputType) (OutputType, error)`。其中 `InputType` 和 `OutputType` 必须与 encoding/json 标准库兼容。函数计算会使用 `json.Unmarshal` 方法对传入的 `InputType` 进行反序列化，使用 `json.Marshal` 方法对返回的 `OutputType` 进行序列化。

### 本地调试

要调用 Event Handler，首先需要准备事件数据，我这里把事件保存在文件 event.json 中：

```json
{
  "name": "shaouai",
  "age": 28,
  "department": "IT"
}
```

如果设置了触发器，则需要使用 `s cli fc-event` 命令生成事件模版，如命令 `s cli fc-event sls` 生成 SLS 触发器事件模版。

然后使用 `local invoke` 命令，进行本地事件函数调试。但是此命令并不会执行 pre-deploy 中的 actions。

我尝试了 s build 命令，发现其也不会执行 pre-deploy，可能 Go 语言并不在 s build 命令的考虑中。这里只能自己编译（省略部分输出结果，用 ... 代替）：

```shell
$ GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o code/main code/main.go
$ s local invoke --event-file event.json 
{
        "name": "shaouai",
        "age": 28,
        "department": "IT"
}
...
FC Invoke Start RequestId: a7949507-8b85-4c94-8a6c-69b3c9ae38bc
FC Invoke End RequestId: a7949507-8b85-4c94-8a6c-69b3c9ae38bc

RequestId: a7949507-8b85-4c94-8a6c-69b3c9ae38bc  Billed Duration: 238 ms   Memory Size: 3933 MB  Max Memory Used: 52 MB

FC Local Invoke Result:
{"name":"shaouai","age":28,"department":"IT"}

End of method: invoke
...
```

第一次请求时，可能需要拉取 go1 运行时的镜像并且创建容器，速度可能会慢点。

关于本地调试，详见 [Local 命令](https://docs.serverless-devs.com/fc/command/local)。

### 部署

本地调试好后，可以将函数部署到线上了。

这里介绍使用 `s` 命令进行部署。因为涉及到本地和远端通信，要先用 `s config` 命令[配置密钥](https://docs.serverless-devs.com/serverless-devs/command/config)。

使用 `s deploy --use-local -y` 命令部署（省略部分输出结果，用 ... 代替）：

```shell
$ s deploy --use-local -y   
[2023-10-27 14:01:12] [INFO] [S-CORE] - Start the pre-action
[2023-10-27 14:01:12] [INFO] [S-CORE] - Action: go mod tidy
[2023-10-27 14:01:12] [INFO] [S-CORE] - Action: GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go
[2023-10-27 14:01:12] [INFO] [S-CORE] - End the pre-action
...
Tips for next step
======================
* Display information of the deployed resource: s info
...
* Invoke remote function: s invoke
...
event-handler-builtin-runtime-example-service: 
  region:   cn-zhangjiakou
  service: 
    name: event-handler-builtin-runtime-example-service
  function: 
    name:       event-handler-builtin-runtime-example-function
    runtime:    go1
    handler:    main
    memorySize: 128
    timeout:    10
    cpu:        0.05
    diskSize:   512
```

部署后，输出了云函数的信息；我们也可以使用 `s info` 命令查看云函数信息；也可以在阿里云控制台查看。

### 本地调用云函数

使用 `s invoke` 命令在本地调用云函数，使用之前的事件数据 event.json：

```shell
$ s invoke -f event.json 
Reading event file content:
{
  "name": "shaouai",
  "age": 28,
  "department": "IT"
}

========= FC invoke Logs begin =========
FC Invoke Start RequestId: 1-653b53cb-84ad25f3928599ca34f226cb
FC Invoke End RequestId: 1-653b53cb-84ad25f3928599ca34f226cb

Duration: 1.89 ms, Billed Duration: 2 ms, Memory Size: 128 MB, Max Memory Used: 6.52 MB
========= FC invoke Logs end =========

FC Invoke instanceId: c-653b53cb-34cba99522524ee1b842

FC Invoke Result:
{"name":"shaouai","age":28,"department":"IT"}

End of method: invoke
```

### 登录实例

可以在阿里云控制台登录实例，也可以在本地登录。如果现在没有实例，可以发起一个请求，让函数计算创建一个实例。

```shell
$ s instance list       
event-handler-builtin-runtime-example-service: 
  event-handler-builtin-runtime-example-function: 
    instances: 
      - 
        instanceId: c-653b552d-2b4854ab13264043b422
        versionId:  0
$ s instance exec  c-653b552d-2b4854ab13264043b422 -it /bin/bash
root@sr-653b1a86-8d256d47bda14a5fb15d:/# ls
bin   code  etc   lib    media  opt   root  sbin  sys  usr
boot  dev   home  lib64  mnt    proc  run   srv   tmp  var
root@sr-653b1a86-8d256d47bda14a5fb15d:/# ls code/
go.mod  go.sum  main  main.go
```

## 自定义运行时

TBD

## 自定义容器运行时

TBD

# 参见

* [函数类型选型](https://help.aliyun.com/zh/fc/product-overview/overview-30)
* [请求处理程序（Handler）](https://help.aliyun.com/zh/fc/user-guide/handlers-1?)
* [事件请求处理程序（Event Handler）](https://help.aliyun.com/zh/fc/event-handlers-6)
* [代码开发 - Go](https://help.aliyun.com/zh/fc/user-guide/go-1)
* [Serverless Devs 描述文件（Yaml）规范](https://docs.serverless-devs.com/serverless-devs/yaml)
* [ 函数计算（FC）组件 Yaml 规范](https://docs.serverless-devs.com/fc/yaml/readme)