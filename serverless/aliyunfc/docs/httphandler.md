HTTP 请求处理程序（HTTP Handler）

# 基本信息

请求处理程序分为事件请求处理程序（Event Handler）和 HTTP 请求处理程序（HTTP Handler），其中事件请求由各种事件源触发生成，HTTP
请求则由 HTTP 触发器触发生成。

本文关注 HTTP Handler。编程语言使用 Go。

# HTTP Handler

## 内置运行时

### 示例项目

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

需要注意的是，`codeUri` 目录中的内容是最终的交付物，函数计算最终会把此目录下的内容拷贝到容器 `/code` 目录下。使用 Go 语言时的交付物是一个二进制可执行文件，因此我们要确保这个二进制文件出现在 `codeUri` 指定的目录中，而且文件名是 `handler` 所指定的值。这里通过 `pre-deploy` 指定部署之前的 actions，在部署之前进行编译。

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

可以使用 `s local start` 命令把 HTTP 函数部署在本地，方便调试。但是此命令并不会执行 pre-deploy 中的 actions。

我尝试了 `s build` 命令，发现其也不会执行 pre-deploy，可能 Go 语言并不在 `s build` 命令的考虑中。这里只能自己编译（省略部分输出结果，用 ... 代替）：

```shell
$ GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o code/main code/main.go
$ s local start
...
[2023-10-26 23:18:36] [INFO] [FC-LOCAL-INVOKE] - CustomDomain auto of http-handler-builtin-runtime-example-service/http-handler-builtin-runtime-example-function was registered
        url: http://localhost:7342/
        methods: GET,POST
        authType: anonymous
...
function compute app listening on port 7342!
```

使用 `curl` 命令进行本地请求：

```shell
$  curl http://localhost:7342/              
2023-10-27 09:31:05
Request Method: GET
$
$ curl -X POST -d 'test POST body' http://localhost:7342/ 
2023-10-27 09:32:25
Request Method: POST
test POST body
```

第一次请求时，可能需要拉取 go1 运行时的镜像并且创建容器，速度可能会慢点。

### 部署

本地调试好后，可以将函数部署到线上了。

这里介绍使用 `s` 命令进行部署。因为涉及到本地和远端通信，要先用 `s config` 命令[配置密钥](https://docs.serverless-devs.com/serverless-devs/command/config)。

使用 `s deploy --use-local -y` 命令部署（省略部分输出结果，用 ... 代替）：

```shell
$ s deploy --use-local -y   
[2023-10-27 09:46:38] [INFO] [S-CORE] - Start the pre-action
[2023-10-27 09:46:38] [INFO] [S-CORE] - Action: go mod tidy
[2023-10-27 09:46:39] [INFO] [S-CORE] - Action: GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go
[2023-10-27 09:46:39] [INFO] [S-CORE] - End the pre-action
...
Tips for next step
======================
* Display information of the deployed resource: s info
* Invoke remote function: s invoke
...
  url: 
    system_url:          https://http-hafunction-http-ha-service-syfmwcasfs.cn-zhangjiakou.fcapp.run
    system_intranet_url: https://http-hafunction-http-ha-service-syfmwcasfs.cn-zhangjiakou-vpc.fcapp.run
    custom_domain: 
      - 
        domain: http://http-handler-builtin-runtime-example-function.http-handler-builtin-runtime-example-service.1810657881264284.cn-zhangjiakou.fc.devsapp.net
```

部署后，输出了云函数的公网 url（system_url）和自定义域名 url（custom_domain）；我们也可以使用 `s info` 命令查看云函数信息；也可以在阿里云控制台查看。

### 本地调用云函数

因为这是一个 HTTP handler，我们可以用各种方式发起 HTTP 请求，但这里还是介绍下用 `s invoke` 命令进行远程调用。

首先使用 `s cli fc-event http` 命令生成调用参数的模版：

```shell
$ s cli fc-event http
      👓 Parameter Template Path: event-template/http-parameter.json
      You could user fc component invoke method and specify the event.
      E.g: [s projectName invoke --event-file  event-template/http-parameter.json]
$ cat event-template/http-parameter.json 
{
  "path": "string",
  "method": "POST",
  "headers": {
    "key": "value"
  },
  "queries": {
    "key": "value"
  },
  "body": "body"
}                     
```

把模版修改成这样：

```shell
$ cat event-template/http-parameter.json 
{
  "path": "/",
  "method": "POST",
  "headers": {
    "key": "value"
  },
  "queries": {
    "key": "value"
  },
  "body": "invoke body"
}
```

发起远程调用：

```shell
$ s invoke -f event-template/http-parameter.json
Reading event file content:
{
  "path": "/",
  "method": "POST",
  "headers": {
    "key": "value"
  },
  "queries": {
    "key": "value"
  },
  "body": "invoke body"
}

Request url: https://http-hafunction-http-ha-service-syfmwcasfs.cn-zhangjiakou.fcapp.run/

FC Invoke instanceId: c-653b1c76-267e16d848714db98c75

FC Invoke Result:
2023-10-27 10:12:15
Request Method: POST
invoke body

End of method: invoke
```

### 登录实例

可以在阿里云控制台登录实例，也可以在本地登录。如果现在没有实例，可以发起一个请求，让函数计算创建一个实例。

```shell
$ s instance list                               
http-handler-builtin-runtime-example-service: 
  http-handler-builtin-runtime-example-function: 
    instances: 
      - 
        instanceId: c-653b221b-67d7232869314a88a7f9
        versionId:  0
$ s instance exec  c-653b221b-67d7232869314a88a7f9 -it /bin/bash
root@sr-653ae858-9b5d81f96fda4b2bbebf:/# ls
bin   code  etc   lib    media  opt   root  sbin  sys  usr
boot  dev   home  lib64  mnt    proc  run   srv   tmp  var
root@sr-653ae858-9b5d81f96fda4b2bbebf:/# cd code/
root@sr-653ae858-9b5d81f96fda4b2bbebf:/code# ls
main  main.go
```

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
