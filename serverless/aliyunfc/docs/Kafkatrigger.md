# 基本信息

消息队列 Kafka ----> 事件总线 EventBridge 事件流 ----> 函数计算

# 并发配置

### 消费任务并发数

官方解释：

> 通过设置消费并发数，您可以配置源 Kafka 实例的消费者数量。当您的 Kafka 实例有多个分区时，配置和分区数相同的消费并发数可以提高 Kafka 触发函数的并发情况。

> 消费者的并发数量，取值范围为[1,Topic的分区数]。

> 您可以通过设置并发消费线程数提高吞吐，目前仅Kafka触发器支持设置并发配额，云消息队列 Kafka 版并发消费需配合Topic分区共同实现，包括以下几种场景。
>
> -  Topic分区数=并发消费数：一个线程消费一个Topic分区。
> -  Topic分区数>并发消费数：多个并发消费会均摊所有分区消费。
> -  Topic分区数<并发消费数：一个线程消费一个Topic分区，多出的消费数无效。
>
> 说明：为保证您的资源被充分利用，建议您选择Topic分区数=并发消费数或Topic分区数>并发消费数场景。

我的理解：通过设置 Kafka 触发器内部并发消费线程数，可以让我们直观感受到消费速度提升（如果云函数处理事件速度够快的话）。

### 投递并发最大值

官方解释：

> Kafka 投递到函数计算的并发最大值。

> Kafka 消息投递到函数计算的并发最大值，取值范围为 1~300。该参数仅对同步调用生效。如果需要更高的并发，请进入 EventBridge 配额中心申请配额名称为 EventStreaming FC Sink 同步投递最大并发数的配额。

我的理解：把投递理解为请求，这个参数可以理解为 Kafka 触发器向函数计算发起的请求的并发最大值。

### （批量）推送配置

官方解释：

> 批量推送条数：一次调用函数发送的最大批量消息条数，当积压的消息数量到达设定值时才会发送请求，取值范围为 [1, 10000]。
> 批量推送间隔：调用函数的间隔时间，系统每到间隔时间点会将消息聚合后发给函数计算，取值范围为 [0,15]，单位秒。0 秒表示无等待时间，直接投递。

注意：

- 两个条件满足其中一个时，触发函数执行
- 需结合 body 大小限制决定是否减少聚合消息数
    - 同步调用：32 MB
    - 异步调用：128 KB

场景分析（TBD）

### 函数实例并发度

官方解释：

> 函数计算支持一个实例同时并发执行多个请求，这个值用来配置单个函数实例可以同时处理多少个请求。

我的理解：一次请求可以发送多个事件（根据批量推送配置而定）。

### 各参数间关系

投递并发最大值决定了 Kafka 触发器向函数计算发起的请求的并发最大值，而是否能达到这个最大值，是由消费任务并发数决定的，Kafka 触发器内部并发消费线程数越多，越有可能达到这个最大值。如果把 Kafka 触发器向函数计算发起的请求的实际并发值，叫做实际投递并发值，那么函数计算实际创建的函数实例数为：

* `⌈<实际投递并发值> / <函数实例并发度>⌉`

而函数计算创建的函数实例数上限为：

* `⌈<投递并发最大值> / <函数实例并发度>⌉`

# 示例项目

这是一个使用内置运行时 go1 实现的 Event Handler，配置了 Kafka 触发器。

## 项目结构

```shell
$ tree .                   
.
├── code
│   ├── go.mod
│   ├── go.sum
│   └── main.go
└── s.yaml

2 directories, 4 files
```

```yaml
edition: 1.0.0 # 命令行YAML规范版本，遵循语义化版本（Semantic Versioning）规范
name: dev_event-handler-builtin-runtime-Kafka-trigger-example # 项目/应用名称
access: default # 密钥别名

services: # 应用所包含的服务，可以包含多个
  dev_event-handler-builtin-runtime-Kafka-trigger-example-service: # 服务/模块名称
    component: devsapp/fc # 组件名称，这里使用阿里云函数计算（FC）组件
    actions: # 自定义执行逻辑
      pre-deploy: # 在 deploy 之前运行，把 main.go 编译为可执行的二进制文件
        - run: go mod tidy
          path: ./code
        - run: GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o target/main main.go
          path: ./code
    props: # 组件的属性值
      region: cn-shanghai # 地域
      service: # 服务配置
        name: dev_event-handler-builtin-runtime-Kafka-trigger-example-service # service 名称
        description: "dev event handler builtin runtime Kafka trigger example service" # Service 的简短描述
        internetAccess: false # 设为 true 让 function 可以访问公网
        tracingConfig: Disable # 链路追踪，可取值：Enable、Disable
        role: acs:ram::xxx:role/aliyunfcdefaultrole # 授予函数计算所需权限的RAM role
        logConfig: # log配置，function产生的log会写入这里配置的logstore
          logstore: xxx # loghub中的logstore名称
          project: xxx # loghub中的project名称
          enableRequestMetrics: true
          enableInstanceMetrics: true
          logBeginRule: DefaultRegex # 日志是否切分，取值 DefaultRegex/None
        vpcConfig: # VPC配置, 配置后function可以访问指定VPC
          securityGroupId: sg-xxx # 安全组ID
          vpcId: vpc-xxx # VPC ID
          vswitchIds: # 交换机 ID 列表
            - vsw-xxx
        nasConfig: null # NAS配置, 配置后function可以访问指定NAS
        ossMountConfig: null # OSS挂载配置, 配置后function可以访问指定OSS bucket
        vpcBinding: null # 仅允许指定 VPC 调用函数
      function: # 函数配置
        name: dev_event-handler-builtin-runtime-Kafka-trigger-example-function # function 名称
        description: "dev event handler builtin runtime Kafka trigger example function" # function 的简短描述
        codeUri: ./code/target # 代码位置，目录下的内容是最终的交付物
        handler: main # function 执行的入口，具体格式和语言相关
        memorySize: 128 # function 的内存规格
        runtime: go1 # 运行时
        timeout: 120 # function 运行的超时时间
        cpu: 0.1 # 函数的 CPU 规格，单位为 vCPU，为 0.05 vCPU 的倍数
        diskSize: 512 # 函数的磁盘规格，单位为 MB，可选值为 512 MB 或 10240 MB
        instanceConcurrency: 26 # 实例并发度，单实例多并发，一个函数实例可以并发处理这么多请求
        instanceType: e1 # 函数实例类型，可选值为：e1（弹性实例）、c1（性能实例）、fc.gpu.tesla.1（GPU T4实例）、fc.gpu.ampere.1（GPU A10实例）。
        environmentVariables: # 环境变量
          TZ: "Asia/Shanghai" # 设置时区为东 8 区
      triggers: # 触发器配置
        - name: dev_event-handler-builtin-runtime-Kafka-trigger-example-trigger # 触发器名称
          type: eventbridge # 触发器类型
          qualifier: LATEST # 触发器函数的版本或者别名，默认 LATEST
          config: # 触发器配置
            triggerEnable: true # 触发器禁用开关
            asyncInvocationType: false # 触发器调用函数的方式。目前支持同步调用以及异步调用
            eventRuleFilterPattern: "{}" # 事件模式
            eventSinkConfig: # 事件目标配置
              deliveryOption: # 事件投递参数
                mode: event-streaming # 与 runOptions 中的 mode 参数含义相同，但是优先级更低
                eventSchema: CloudEvents # 以通用格式描述事件数据的规范
                concurrency: 26 # 投递并发最大值，Kafka 投递到函数计算的并发最大值
            runOptions: # 触发器运行时参数
              mode: event-streaming # event source 为 Kafka 时，只支持 event-streaming 模式
              maximumTasks: 1 # 并发消费者数量，只有在指定 Kafka 事源时该参数有效
              errorsTolerance: 'ALL' # 容错策略，即发生错误时是否选择容错。ALL:允许容错；NONE:禁止容错。
              retryStrategy: # 事件推送失败时的重试策略相关参数
                PushRetryStrategy: 'BACKOFF_RETRY' # 事件推送失败时的重试策略。BACKOFF_RETRY: 退避重试策略。EXPONENTIAL_DECAY_RETRY: 指数衰减重试。
              deadLetterQueue: # 死信队列配置，若配置了该配置，超过重试策略后的事件将被放入该队列中
                Arn: xxx
              batchWindow: # 调用函数时的批处理参数
                CountBasedWindow: 10 # 一次调用函数发送的最大批量消息条数，当积压的消息数量到达设定值时才会发送请求，取值范围为 [1, 10000]。
                TimeBasedWindow: 15 # 调用函数的间隔时间，系统每到间隔时间点会将消息聚合后发给函数计算，取值范围为 [0,15]，单位秒。0 秒表示无等待时间，直接投递。
            eventSourceConfig: # 事件源配置
              eventSourceType: Kafka # 触发器事件源类型
              eventSourceParameters: # 自定义事件源参数
                sourceKafkaParameters: # 事件源为消息队列 Kafka 时的自定义参数配置
                  RegionId: cn-shanghai # 消息队列 Kafka 版的实例所属地域
                  InstanceId: "xxx" # 消息队列 Kafka 版的实例 ID，需要提前创建
                  Topic: "xxx" # 消息队列 Kafka 版的 Topic 名称，需要提前创建
                  ConsumerGroup: "xxx" # 消息队列 Kafka 版的资源组 ID，需要提前创建
                  OffsetReset: "earliest" # 消息的消费位点，可选值有 lastest 和 earliest，分别表示最新位点以及最早位点
                  Network: "Default" # 所用网络类型，可选值有 PublicNetwork 以及 Default，前者需选择另外的专有网络VPC、交换机和安全组，后者表示默认使用部署Kafka实例时选择的VPC ID和vSwitch ID。

```

- 关于可配置的参数，见 [triggers 字段](https://docs.serverless-devs.com/fc/yaml/triggers)

```go
package main

import (
	"context"
	"encoding/json"
	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/aliyun/fc-runtime-go-sdk/fccontext"
)

type kafkaMsg map[string]any

// 这里 Kafka 消息格式是 JSON
func (k *kafkaMsg) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	var v map[string]any
	err = json.Unmarshal([]byte(s), &v)
	if err != nil {
		return err
	}

	*k = v
	return nil
}

// https://help.aliyun.com/zh/fc/apsaramq-for-kafka-trigger?spm=a2c4g.11186623.0.0.14994dddWPAIcY#:~:text=%E5%8D%95%E5%87%BB%E7%A1%AE%E5%AE%9A%E3%80%82-,event,-%E6%A0%BC%E5%BC%8F%E5%A6%82%E4%B8%8B%E6%89%80%E7%A4%BA%EF%BC%9A
type event struct {
	SpecVersion     string    `json:"specversion"`
	ID              string    `json:"id"`
	Source          string    `json:"source"`
	Type            string    `json:"type"`
	Subject         string    `json:"subject"`
	DataContentType string    `json:"datacontenttype"`
	Time            string    `json:"time"`
	AliyunAccountID string    `json:"aliyunaccountid"`
	Data            eventData `json:"data"`
}

type eventData struct {
	Topic     string          `json:"topic"`
	Partition int             `json:"partition"`
	Offset    int64           `json:"offset"`
	Timestamp int64           `json:"timestamp"`
	Headers   eventDataHeader `json:"headers"`
	Key       string          `json:"key"`
	Value     kafkaMsg        `json:"value"`
}

type eventDataHeader struct {
	Headers    []string `json:"headers"`
	IsReadOnly bool     `json:"isReadOnly"`
}

func main() {
	fc.Start(HandleRequest)
}

func HandleRequest(ctx context.Context, cloudEvents []event) ([]event, error) {
	fctx, _ := fccontext.FromContext(ctx)
	flog := fctx.GetLogger()
	flog.Infof("Start processing %d events", len(cloudEvents))

	for _, e := range cloudEvents {
		flog.Infof("Start processing event: %+v", e)
	}

	return cloudEvents, nil
}
```

## 本地调试

准备事件数据，我把事件保存在一个文件中，事件格式参考[文档](https://help.aliyun.com/zh/fc/apsaramq-for-kafka-trigger?spm=a2c4g.11186623.0.0.3f042b3c6b5Ooo#:~:text=%E5%8D%95%E5%87%BB%E7%A1%AE%E5%AE%9A%E3%80%82-,event,-%E6%A0%BC%E5%BC%8F%E5%A6%82%E4%B8%8B%E6%89%80%E7%A4%BA%EF%BC%9A)，是一个 JSON 数组：

```shell
$ cat event.json              
[
    {
        "specversion":"1.0",
        "id":"8e215af8-ca18-4249-8645-f96c1026****",
        "source":"acs:alikafka",
        "type":"alikafka:Topic:Message",
        "subject":"acs:alikafka_pre-cn-i7m2t7t1****:topic:mytopic",
        "datacontenttype":"application/json; charset=utf-8",
        "time":"2022-06-23T02:49:51.589Z",
        "aliyunaccountid":"164901546557****",
        "data":{
            "topic":"xxtopic",
            "partition":7,
            "offset":25,
            "timestamp":1655952591589,
            "headers":{
                "headers":[
                ],
                "isReadOnly":false
            },
            "key":"keytest",
            "value": "{\"name\":\"shaouai\",\"age\":28,\"department\":\"IT\"}"
        }
    }
]
```

把 Go 程序编译成二进制文件：

```shell
$ GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o code/target/main code/main.go
```

本地调用，省略了部分结果：

```shell
$  s local invoke -f event.json              
[
    {
        "specversion":"1.0",
        "id":"8e215af8-ca18-4249-8645-f96c1026****",
        "source":"acs:alikafka",
        "type":"alikafka:Topic:Message",
        "subject":"acs:alikafka_pre-cn-i7m2t7t1****:topic:mytopic",
        "datacontenttype":"application/json; charset=utf-8",
        "time":"2022-06-23T02:49:51.589Z",
        "aliyunaccountid":"164901546557****",
        "data":{
            "topic":"xxtopic",
            "partition":7,
            "offset":25,
            "timestamp":1655952591589,
            "headers":{
                "headers":[
                ],
                "isReadOnly":false
            },
            "key":"keytest",
            "value": "{\"name\":\"shaouai\",\"age\":28,\"department\":\"IT\"}"
        }
    }
]

FC Invoke Start RequestId: f2847ae1-ef5e-4382-a88b-74ec3305158b
2023-10-29T09:24:07.491Z f2847ae1-ef5e-4382-a88b-74ec3305158b [INFO] main.go:65: Start processing 1 events
2023-10-29T09:24:07.494Z f2847ae1-ef5e-4382-a88b-74ec3305158b [INFO] main.go:68: Start processing event: {SpecVersion:1.0 ID:8e215af8-ca18-4249-8645-f96c1026**** Source:acs:alikafka Type:alikafka:Topic:Message Subject:acs:alikafka_pre-cn-i7m2t7t1****:topic:mytopic DataContentType:application/json; charset=utf-8 Time:2022-06-23T02:49:51.589Z AliyunAccountID:164901546557**** Data:{Topic:xxtopic Partition:7 Offset:25 Timestamp:1655952591589 Headers:{Headers:[] IsReadOnly:false} Key:keytest Value:map[age:28 department:IT name:shaouai]}}
FC Invoke End RequestId: f2847ae1-ef5e-4382-a88b-74ec3305158b

RequestId: f2847ae1-ef5e-4382-a88b-74ec3305158b          Billed Duration: 199 ms         Memory Size: 3933 MB    Max Memory Used: 65 MB

FC Local Invoke Result:
[{"specversion":"1.0","id":"8e215af8-ca18-4249-8645-f96c1026****","source":"acs:alikafka","type":"alikafka:Topic:Message","subject":"acs:alikafka_pre-cn-i7m2t7t1****:topic:mytopic","datacontenttype":"application/json; charset=utf-8","time":"2022-06-23T02:49:51.589Z","aliyunaccountid":"164901546557****","data":{"topic":"xxtopic","partition":7,"offset":25,"timestamp":1655952591589,"headers":{"headers":[],"isReadOnly":false},"key":"keytest","value":{"age":28,"department":"IT","name":"shaouai"}}}]

End of method: invoke

Tips for next step
======================
* Deploy Resources: s deploy
dev_event-handler-builtin-runtime-Kafka-trigger-example-service: 
  status: succeed
```

可以看到，调用成功，返回了原始输入的事件数据。

这次用 3 个事件调用看看（`s local invoke` 命令会在开始打印事件数据）：

```shell
$ s local invoke -f event.json
[
    {
        "specversion":"1.0",
        "id":"8e215af8-ca18-4249-8645-f96c1026****",
        "source":"acs:alikafka",
        "type":"alikafka:Topic:Message",
        "subject":"acs:alikafka_pre-cn-i7m2t7t1****:topic:mytopic",
        "datacontenttype":"application/json; charset=utf-8",
        "time":"2022-06-23T02:49:51.589Z",
        "aliyunaccountid":"164901546557****",
        "data":{
            "topic":"xxtopic",
            "partition":7,
            "offset":25,
            "timestamp":1655952591589,
            "headers":{
                "headers":[
                ],
                "isReadOnly":false
            },
            "key":"keytest",
            "value": "{\"name\":\"shaouai\",\"age\":28,\"department\":\"IT\"}"
        }
    },
    {
        "specversion":"1.0",
        "id":"8e215af8-ca18-4249-8645-f96c1026****",
        "source":"acs:alikafka",
        "type":"alikafka:Topic:Message",
        "subject":"acs:alikafka_pre-cn-i7m2t7t1****:topic:mytopic",
        "datacontenttype":"application/json; charset=utf-8",
        "time":"2022-06-23T02:49:51.589Z",
        "aliyunaccountid":"164901546557****",
        "data":{
            "topic":"xxtopic",
            "partition":7,
            "offset":26,
            "timestamp":1655952591589,
            "headers":{
                "headers":[
                ],
                "isReadOnly":false
            },
            "key":"keytest",
            "value": "{\"name\":\"shaouai\",\"age\":29,\"department\":\"IT\"}"
        }
    },
    {
        "specversion":"1.0",
        "id":"8e215af8-ca18-4249-8645-f96c1026****",
        "source":"acs:alikafka",
        "type":"alikafka:Topic:Message",
        "subject":"acs:alikafka_pre-cn-i7m2t7t1****:topic:mytopic",
        "datacontenttype":"application/json; charset=utf-8",
        "time":"2022-06-23T02:49:51.589Z",
        "aliyunaccountid":"164901546557****",
        "data":{
            "topic":"xxtopic",
            "partition":7,
            "offset":27,
            "timestamp":1655952591589,
            "headers":{
                "headers":[
                ],
                "isReadOnly":false
            },
            "key":"keytest",
            "value": "{\"name\":\"shaouai\",\"age\":30,\"department\":\"IT\"}"
        }
    }
]

[2023-10-29 17:40:13] [INFO] [FC-CORE] - Skip pulling image registry.cn-beijing.aliyuncs.com/aliyunfc/runtime-go1:1.10.9...
Aliyun FunctionComputer runtime emulator start.
FC runtime init Duration: 47 ms
FC Invoke Start RequestId: 3ab1bd3c-03db-42b6-82aa-dce17f71db37
2023-10-29T09:40:14.323Z 3ab1bd3c-03db-42b6-82aa-dce17f71db37 [INFO] main.go:65: Start processing 3 events
2023-10-29T09:40:14.325Z 3ab1bd3c-03db-42b6-82aa-dce17f71db37 [INFO] main.go:68: Start processing event: {SpecVersion:1.0 ID:8e215af8-ca18-4249-8645-f96c1026**** Source:acs:alikafka Type:alikafka:Topic:Message Subject:acs:alikafka_pre-cn-i7m2t7t1****:topic:mytopic DataContentType:application/json; charset=utf-8 Time:2022-06-23T02:49:51.589Z AliyunAccountID:164901546557**** Data:{Topic:xxtopic Partition:7 Offset:25 Timestamp:1655952591589 Headers:{Headers:[] IsReadOnly:false} Key:keytest Value:map[age:28 department:IT name:shaouai]}}
2023-10-29T09:40:14.325Z 3ab1bd3c-03db-42b6-82aa-dce17f71db37 [INFO] main.go:68: Start processing event: {SpecVersion:1.0 ID:8e215af8-ca18-4249-8645-f96c1026**** Source:acs:alikafka Type:alikafka:Topic:Message Subject:acs:alikafka_pre-cn-i7m2t7t1****:topic:mytopic DataContentType:application/json; charset=utf-8 Time:2022-06-23T02:49:51.589Z AliyunAccountID:164901546557**** Data:{Topic:xxtopic Partition:7 Offset:26 Timestamp:1655952591589 Headers:{Headers:[] IsReadOnly:false} Key:keytest Value:map[age:29 department:IT name:shaouai]}}
2023-10-29T09:40:14.325Z 3ab1bd3c-03db-42b6-82aa-dce17f71db37 [INFO] main.go:68: Start processing event: {SpecVersion:1.0 ID:8e215af8-ca18-4249-8645-f96c1026**** Source:acs:alikafka Type:alikafka:Topic:Message Subject:acs:alikafka_pre-cn-i7m2t7t1****:topic:mytopic DataContentType:application/json; charset=utf-8 Time:2022-06-23T02:49:51.589Z AliyunAccountID:164901546557**** Data:{Topic:xxtopic Partition:7 Offset:27 Timestamp:1655952591589 Headers:{Headers:[] IsReadOnly:false} Key:keytest Value:map[age:30 department:IT name:shaouai]}}
FC Invoke End RequestId: 3ab1bd3c-03db-42b6-82aa-dce17f71db37

RequestId: 3ab1bd3c-03db-42b6-82aa-dce17f71db37          Billed Duration: 192 ms         Memory Size: 3933 MB    Max Memory Used: 54 MB

FC Local Invoke Result:
[{"specversion":"1.0","id":"8e215af8-ca18-4249-8645-f96c1026****","source":"acs:alikafka","type":"alikafka:Topic:Message","subject":"acs:alikafka_pre-cn-i7m2t7t1****:topic:mytopic","datacontenttype":"application/json; charset=utf-8","time":"2022-06-23T02:49:51.589Z","aliyunaccountid":"164901546557****","data":{"topic":"xxtopic","partition":7,"offset":25,"timestamp":1655952591589,"headers":{"headers":[],"isReadOnly":false},"key":"keytest","value":{"age":28,"department":"IT","name":"shaouai"}}},{"specversion":"1.0","id":"8e215af8-ca18-4249-8645-f96c1026****","source":"acs:alikafka","type":"alikafka:Topic:Message","subject":"acs:alikafka_pre-cn-i7m2t7t1****:topic:mytopic","datacontenttype":"application/json; charset=utf-8","time":"2022-06-23T02:49:51.589Z","aliyunaccountid":"164901546557****","data":{"topic":"xxtopic","partition":7,"offset":26,"timestamp":1655952591589,"headers":{"headers":[],"isReadOnly":false},"key":"keytest","value":{"age":29,"department":"IT","name":"shaouai"}}},{"specversion":"1.0","id":"8e215af8-ca18-4249-8645-f96c1026****","source":"acs:alikafka","type":"alikafka:Topic:Message","subject":"acs:alikafka_pre-cn-i7m2t7t1****:topic:mytopic","datacontenttype":"application/json; charset=utf-8","time":"2022-06-23T02:49:51.589Z","aliyunaccountid":"164901546557****","data":{"topic":"xxtopic","partition":7,"offset":27,"timestamp":1655952591589,"headers":{"headers":[],"isReadOnly":false},"key":"keytest","value":{"age":30,"department":"IT","name":"shaouai"}}}]

End of method: invoke

Tips for next step
======================
* Deploy Resources: s deploy
dev_event-handler-builtin-runtime-Kafka-trigger-example-service: 
  status: succeed
```

## 部署

`s deploy`

## 本地调用云函数

用之前的事件数据调用部署好的云函数：

```shell
$ s invoke -f event.json
Reading event file content:
[
    {
        "specversion":"1.0",
        "id":"8e215af8-ca18-4249-8645-f96c1026****",
        "source":"acs:alikafka",
        "type":"alikafka:Topic:Message",
        "subject":"acs:alikafka_pre-cn-i7m2t7t1****:topic:mytopic",
        "datacontenttype":"application/json; charset=utf-8",
        "time":"2022-06-23T02:49:51.589Z",
        "aliyunaccountid":"164901546557****",
        "data":{
            "topic":"xxtopic",
            "partition":7,
            "offset":25,
            "timestamp":1655952591589,
            "headers":{
                "headers":[
                ],
                "isReadOnly":false
            },
            "key":"keytest",
            "value": "{\"name\":\"shaouai\",\"age\":28,\"department\":\"IT\"}"
        }
    },
    {
        "specversion":"1.0",
        "id":"8e215af8-ca18-4249-8645-f96c1026****",
        "source":"acs:alikafka",
        "type":"alikafka:Topic:Message",
        "subject":"acs:alikafka_pre-cn-i7m2t7t1****:topic:mytopic",
        "datacontenttype":"application/json; charset=utf-8",
        "time":"2022-06-23T02:49:51.589Z",
        "aliyunaccountid":"164901546557****",
        "data":{
            "topic":"xxtopic",
            "partition":7,
            "offset":26,
            "timestamp":1655952591589,
            "headers":{
                "headers":[
                ],
                "isReadOnly":false
            },
            "key":"keytest",
            "value": "{\"name\":\"shaouai\",\"age\":29,\"department\":\"IT\"}"
        }
    },
    {
        "specversion":"1.0",
        "id":"8e215af8-ca18-4249-8645-f96c1026****",
        "source":"acs:alikafka",
        "type":"alikafka:Topic:Message",
        "subject":"acs:alikafka_pre-cn-i7m2t7t1****:topic:mytopic",
        "datacontenttype":"application/json; charset=utf-8",
        "time":"2022-06-23T02:49:51.589Z",
        "aliyunaccountid":"164901546557****",
        "data":{
            "topic":"xxtopic",
            "partition":7,
            "offset":27,
            "timestamp":1655952591589,
            "headers":{
                "headers":[
                ],
                "isReadOnly":false
            },
            "key":"keytest",
            "value": "{\"name\":\"shaouai\",\"age\":30,\"department\":\"IT\"}"
        }
    }
]


========= FC invoke Logs begin =========
ad3bfcd17471
2023-10-29T09:53:15.456Z 1-653e2b8b-23402f30bde5ad3bfcd17471 [INFO] main.go:65: Start processing 3 events
2023-10-29T09:53:15.456Z 1-653e2b8b-23402f30bde5ad3bfcd17471 [INFO] main.go:68: Start processing event: {SpecVersion:1.0 ID:8e215af8-ca18-4249-8645-f96c1026**** Source:acs:alikafka Type:alikafka:Topic:Message Subject:acs:alikafka_pre-cn-i7m2t7t1****:topic:mytopic DataContentType:application/json; charset=utf-8 Time:2022-06-23T02:49:51.589Z AliyunAccountID:164901546557**** Data:{Topic:xxtopic Partition:7 Offset:25 Timestamp:1655952591589 Headers:{Headers:[] IsReadOnly:false} Key:keytest Value:map[age:28 department:IT name:shaouai]}}
2023-10-29T09:53:15.456Z 1-653e2b8b-23402f30bde5ad3bfcd17471 [INFO] main.go:68: Start processing event: {SpecVersion:1.0 ID:8e215af8-ca18-4249-8645-f96c1026**** Source:acs:alikafka Type:alikafka:Topic:Message Subject:acs:alikafka_pre-cn-i7m2t7t1****:topic:mytopic DataContentType:application/json; charset=utf-8 Time:2022-06-23T02:49:51.589Z AliyunAccountID:164901546557**** Data:{Topic:xxtopic Partition:7 Offset:26 Timestamp:1655952591589 Headers:{Headers:[] IsReadOnly:false} Key:keytest Value:map[age:29 department:IT name:shaouai]}}
2023-10-29T09:53:15.456Z 1-653e2b8b-23402f30bde5ad3bfcd17471 [INFO] main.go:68: Start processing event: {SpecVersion:1.0 ID:8e215af8-ca18-4249-8645-f96c1026**** Source:acs:alikafka Type:alikafka:Topic:Message Subject:acs:alikafka_pre-cn-i7m2t7t1****:topic:mytopic DataContentType:application/json; charset=utf-8 Time:2022-06-23T02:49:51.589Z AliyunAccountID:164901546557**** Data:{Topic:xxtopic Partition:7 Offset:27 Timestamp:1655952591589 Headers:{Headers:[] IsReadOnly:false} Key:keytest Value:map[age:30 department:IT name:shaouai]}}
FC Invoke End RequestId: 1-653e2b8b-23402f30bde5ad3bfcd17471

Duration: 1.49 ms, Billed Duration: 2 ms, Memory Size: 128 MB, Max Memory Used: 12.98 MB
========= FC invoke Logs end =========

FC Invoke instanceId: c-653e2a96-754e88d40299440db4ff

FC Invoke Result:
[{"specversion":"1.0","id":"8e215af8-ca18-4249-8645-f96c1026****","source":"acs:alikafka","type":"alikafka:Topic:Message","subject":"acs:alikafka_pre-cn-i7m2t7t1****:topic:mytopic","datacontenttype":"application/json; charset=utf-8","time":"2022-06-23T02:49:51.589Z","aliyunaccountid":"164901546557****","data":{"topic":"xxtopic","partition":7,"offset":25,"timestamp":1655952591589,"headers":{"headers":[],"isReadOnly":false},"key":"keytest","value":{"age":28,"department":"IT","name":"shaouai"}}},{"specversion":"1.0","id":"8e215af8-ca18-4249-8645-f96c1026****","source":"acs:alikafka","type":"alikafka:Topic:Message","subject":"acs:alikafka_pre-cn-i7m2t7t1****:topic:mytopic","datacontenttype":"application/json; charset=utf-8","time":"2022-06-23T02:49:51.589Z","aliyunaccountid":"164901546557****","data":{"topic":"xxtopic","partition":7,"offset":26,"timestamp":1655952591589,"headers":{"headers":[],"isReadOnly":false},"key":"keytest","value":{"age":29,"department":"IT","name":"shaouai"}}},{"specversion":"1.0","id":"8e215af8-ca18-4249-8645-f96c1026****","source":"acs:alikafka","type":"alikafka:Topic:Message","subject":"acs:alikafka_pre-cn-i7m2t7t1****:topic:mytopic","datacontenttype":"application/json; charset=utf-8","time":"2022-06-23T02:49:51.589Z","aliyunaccountid":"164901546557****","data":{"topic":"xxtopic","partition":7,"offset":27,"timestamp":1655952591589,"headers":{"headers":[],"isReadOnly":false},"key":"keytest","value":{"age":30,"department":"IT","name":"shaouai"}}}]


End of method: invoke
```

输出结果和本地调用时相差不大。

在阿里云控制台也能看到相应日志：

![image](https://github.com/dushaoshuai/dushaoshuai.github.io/assets/56815047/89328b3a-2468-4651-940f-30af9cf30548)

## 并发配置测试

这部分关注 3 个配置参数间的关系：

- （Kafka 触发器）消费任务并发数
- （Kafka 触发器）投递并发最大值
- （函数）实例并发度

注意在下面的过程中，（Kafka 触发器）消费任务并发数一直是 1，因为源 Kafka Topic 只有 1 个分区，设置再多的消费任务并发数也是没用的。

最初的配置值是：

- 投递并发最大值：26
- 实例并发度：26

观察到函数计算只创建了 1 个函数实例：

![image](https://github.com/dushaoshuai/dushaoshuai.github.io/assets/56815047/436d022b-2752-4376-8b1b-2fdbb6988311)

删除 trigger：

```shell
$ s remove trigger
Need to delete the resource in the cn-shanghai area, the operation service is dev_event-handler-builtin-runtime-Kafka-trigger-example-service:

Trigger:

  ┌──────────────────────────────────────────────────────────────────┬─────────────────────────────────────────────────────────────────┬─────────────┬───────────┐
  │                           functionName                           │                           triggerName                           │ triggerType │ qualifier │
  ├──────────────────────────────────────────────────────────────────┼─────────────────────────────────────────────────────────────────┼─────────────┼───────────┤
  │ dev_event-handler-builtin-runtime-Kafka-trigger-example-function │ dev_event-handler-builtin-runtime-Kafka-trigger-example-trigger │ eventbridge │ LATEST    │
  └──────────────────────────────────────────────────────────────────┴─────────────────────────────────────────────────────────────────┴─────────────┴───────────┘
? Are you sure you want to delete these resources? yes
✔ Delete trigger dev_event-handler-builtin-runtime-Kafka-trigger-example-service/dev_event-handler-builtin-runtime-Kafka-trigger-example-function/dev_event-handler-builtin-runtime-Kafka-trigger-example-trigger success.
End of method: remove
```
在阿里云控制台重置 group `dev_test-Kafka-trigger`的消费位点为 0。

改变配置：

- 投递并发最大值：26
- 实例并发度：13

使用 `s deploy` 命令部署。

观察到函数计算创建了 2 个函数实例：

![image](https://github.com/dushaoshuai/dushaoshuai.github.io/assets/56815047/2441fcf6-8a71-4ebd-9a92-195f51441112)

重复删除 trigger、重置消费位点、更改并发配置、部署的步骤（后面不再提及），这次的配置改为：

- 投递并发最大值：26
- 实例并发度：5

观察到函数计算创建了 6 个函数实例：

![image](https://github.com/dushaoshuai/dushaoshuai.github.io/assets/56815047/8e1abaf9-3f86-4f3d-ab3e-4626126729dc)

将配置更改为：

- 投递并发最大值：26
- 实例并发度：2

观察到函数计算创建了 11 个函数实例：

![image](https://github.com/dushaoshuai/dushaoshuai.github.io/assets/56815047/7a65c5b2-abbd-4303-9752-ff7c2b3d7f55)

将配置更改为：

- 投递并发最大值：26
- 实例并发度：1

观察到函数计算创建了 11 个函数实例：

![image](https://github.com/dushaoshuai/dushaoshuai.github.io/assets/56815047/c97bf7f3-1de1-41a9-b75f-246cbcc3bfe9)

汇总上面的数据：

| （Kafka 触发器）消费任务并发数 | （Kafka 触发器）投递并发最大值 | （函数）实例并发度 | 函数计算创建的函数实例数 | 理论上应该创建的函数实例数 |
|--------------------|--------------------|-----------|--------------|---------------|
| 1                  | 26                 | 26        | 1            | 26/26 = 1     |
| 1                  | 26                 | 13        | 2            | 26/13 = 2     |
| 1                  | 26                 | 5         | 6            | 26/5 = 5.2    |
| 1                  | 26                 | 2         | 11           | 26/2 = 13     |
| 1                  | 26                 | 1         | 11           | 26/1 = 26     |

观察到在（函数）实例并发度为 2 和 1 时，函数计算创建的函数实例数不如理论上那么多，这是因为（Kafka 触发器）消费任务并发数只有 1，虽然设置了其投递并发最大值为 26，但受限于其投递能力有限（或者说函数计算处理事件速度较快），并不会给函数计算造成并发 26 个请求的压力，因此函数计算用少于理论值个数的函数实例就可以应付 Kafka 触发器的请求了。

下面我把 handler 实现改一下，处理每个事件时 `time.Sleep(500 * time.Millisecond)`，这样函数计算的压力就比较大了，其创建的函数实例数量就会增加。

将配置更改为：

- 投递并发最大值：26
- 实例并发度：2

观察到函数计算创建了 13 个函数实例：

![image](https://github.com/dushaoshuai/dushaoshuai.github.io/assets/56815047/e140b43a-795e-418a-9184-deefb7b42287)

配置更改为：

- 投递并发最大值：26
- 实例并发度：1

观察到函数计算创建了 26 个函数实例：

![image](https://github.com/dushaoshuai/dushaoshuai.github.io/assets/56815047/40e6fc02-9c46-42e3-b927-b37d8cfc8f4d)

汇总以上数据：

| （Kafka 触发器）消费任务并发数 | （Kafka 触发器）投递并发最大值 | （函数）实例并发度 | 函数计算创建的函数实例数 | 理论上应该创建的函数实例数 |
|--------------------|--------------------|-----------|--------------|---------------|
| 1                  | 26                 | 2         | 13           | 26/2 = 13     |
| 1                  | 26                 | 1         | 26           | 26/1 = 26     |

看到增加函数计算的压力，其创建的函数实例数量就会增加。且函数计算最多创建的函数实例数量遵循这个公式：

- `⌈<Kafka 触发器投递并发最大值> / <函数实例并发度>⌉`

# 重试和容错

### 重试策略

函数执行出错时可进行重试。本来函数计算同步调用是不支持重试的，但是 Kafka 等触发器支持重试。重试策略选项如下：

- 退避重试：重试 3 次，每次重试的时间间隔为介于 10s~20s 的随机值。
- 指数衰减重试：默认重试策略。重试 176 次，每次重试的时间间隔按照指数递增至 512s，总计重试时间为 24 小时，即重试时间间隔为 1s、2s、4s、8s、16s、32s、64s、128s、256s、512s……512s（共计 167 次间隔 512s）。

### 容错策略

当错误发生时的处理方式：

- 允许容错：请求失败且重试失败后，跳过此请求，继续处理下一条请求。
- 禁止容错：请求失败且重试失败后，消费任务阻塞。（**目前禁止容错有 bug**）
- 死信队列：仅当开启允许容错时，可配置死信队列。
    - 如果启用死信队列，未被处理或超过重试次数的消息会被投递到死信队列中。
    - 如果未启用死信队列，超过重试次数的消息会被丢弃。

### 重试和容错最佳实践

TBD

# 参见

- [Kafka 触发器](https://help.aliyun.com/zh/fc/apsaramq-for-kafka-trigger)
- [触发器高级功能](https://help.aliyun.com/zh/fc/user-guide/advanced-features-of-triggers)
- [触发器 Event 格式](https://help.aliyun.com/zh/fc/user-guide/formats-of-event-for-different-triggers)
- [事件总线 EventBridge - 操作指南 - 事件流 - 事件源 - 消息队列 Kafka 版](https://help.aliyun.com/document_detail/439526.html)
- [triggers 字段](https://docs.serverless-devs.com/fc/yaml/triggers)