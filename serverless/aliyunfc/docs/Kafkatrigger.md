Kafka 触发器

# 基本信息

消息队列 Kafka ----> 事件总线 EventBridge 事件流 ----> 函数计算

```yaml
triggers: # 触发器配置
  - name: example-Kafka-trigger # 触发器名称
    type: eventbridge # 触发器类型
    qualifier: LATEST # 触发器函数的版本或者别名，默认 LATEST
    config: # 触发器配置
      triggerEnable: true # 触发器禁用开关
      asyncInvocationType: false # 触发器调用函数的方式。目前支持同步调用以及异步调用。
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
      eventRuleFilterPattern: "{}" # 事件模式，JSON 格式
      eventSinkConfig: # 事件目标配置
        deliveryOption: # 事件投递参数
          mode: event-streaming # 事件投递模型，与 runOptions 中的 mode 参数含义相同，但是优先级更低
          eventSchema: CloudEvents # 指定函数入口参数，有两种取值。CloudEvents: 以通用格式描述事件数据的规范，RawData: 只投递 CloudEvents 中 $data 引用的数据。
          concurrency: 20 # Kafka 投递到函数计算的并发最大值，即函数实例的最大数量
      runOptions: # 触发器运行时参数
        mode: event-streaming # 事件投递模型，可选值有 event-driven 以及 event-streaming。事件源为 Kafka 时，只支持 event-streaming 模式。runOptions 中参数只有在 mode 为 event-streaming 时才有效。
        maximumTasks: 1 # 并发消费者数量，只有在指定 Kafka 事源时该参数有效
        errorsTolerance: 'ALL' # 容错策略，即发生错误时是否选择容错。ALL:允许容错；NONE:禁止容错。
        retryStrategy: # 事件推送失败时的重试策略相关参数
          PushRetryStrategy: 'BACKOFF_RETRY' # 事件推送失败时的重试策略。BACKOFF_RETRY: 退避重试策略。EXPONENTIAL_DECAY_RETRY: 指数衰减重试。
        deadLetterQueue: # 死信队列配置，若配置了该配置，超过重试策略后的事件将被放入该队列中
          Arn: acs:mns:cn-qingdao:123:/queues/queueName
        batchWindow: # 调用函数时的批处理参数
          CountBasedWindow: 1 # 一次调用函数发送的最大批量消息条数，当积压的消息数量到达设定值时才会发送请求，取值范围为 [1, 10000]。
          TimeBasedWindow: 0 # 调用函数的间隔时间，系统每到间隔时间点会将消息聚合后发给函数计算，取值范围为 [0,15]，单位秒。0 秒表示无等待时间，直接投递。
```

* 关于可配置的参数，见 [triggers 字段](https://docs.serverless-devs.com/fc/yaml/triggers)
* 其中有些并发相关的配置，见并发配置部分
* 关于重试策略（PushRetryStrategy）和容错策略（errorsTolerance），见重试和容错部分

# 并发配置

### 消费任务并发数

官方解释：通过设置消费并发数，您可以配置源 Kafka 实例的消费者数量。当您的 Kafka 实例有多个分区时，配置和分区数相同的消费并发数可以提高 Kafka 触发函数的并发情况。

官方解释：消费者的并发数量，取值范围为[1,Topic的分区数]。

官方解释：

> 您可以通过设置并发消费线程数提高吞吐，目前仅Kafka触发器支持设置并发配额，云消息队列 Kafka 版并发消费需配合Topic分区共同实现，包括以下几种场景。
>
> * Topic分区数=并发消费数：一个线程消费一个Topic分区。
>
> * Topic分区数>并发消费数：多个并发消费会均摊所有分区消费。
>
> * Topic分区数<并发消费数：一个线程消费一个Topic分区，多出的消费数无效。
>
> 说明：为保证您的资源被充分利用，建议您选择Topic分区数=并发消费数或Topic分区数>并发消费数场景。

我的理解：通过设置触发器内部并发消费线程数，可以让我们直观感受到消费速度提升（如果云函数处理事件速度够快的话）。因为设置的是触发器，所以在云函数这里我们观察不到任何资源数量的上升。

### 投递并发最大值

官方解释：Kafka 投递到函数计算的并发最大值。

官方解释：Kafka 消息投递到函数计算的并发最大值，取值范围为 1~300。该参数仅对同步调用生效。如果需要更高的并发，请进入 EventBridge 配额中心申请配额名称为 EventStreaming FC Sink 同步投递最大并发数的配额。

我的理解：可以理解为函数实例的最大数量，如果有需要，函数计算会创建最多这么多个函数实例进行并发计算。

### （批量）推送配置 

批量推送条数：一次调用函数发送的最大批量消息条数，当积压的消息数量到达设定值时才会发送请求，取值范围为 [1, 10000]。

批量推送间隔：调用函数的间隔时间，系统每到间隔时间点会将消息聚合后发给函数计算，取值范围为 [0,15]，单位秒。0 秒表示无等待时间，直接投递。

注意：

* 两个条件满足其中一个时，触发函数执行
* 需结合 body 大小限制决定是否减少聚合消息数
  * 同步调用：32 MB
  * 异步调用：128 KB
  
场景分析（TBD）

### 函数实例并发度

官方解释：函数计算支持一个实例同时并发执行多个请求，这个值用来配置单个函数实例可以同时处理多少个请求。

我的理解：一次请求可以发送多个事件（根据批量推送配置而定）。

### 并发能力计算

根据请求处理程序（Handler）对事件数组处理的不同，云函数可以并发处理的事件数量上限不同：

* Handler 对数组中的事件顺序处理：<投递并发最大值> × <函数实例并发度>
* Handler 对数组中的事件并发处理：<投递并发最大值> × <函数实例并发度> × <批量推送条数>

# 重试和容错

### 重试策略

函数执行出错时可进行重试。本来函数计算同步调用是不支持重试的，但是 Kafka 等触发器支持重试。重试策略选项如下：

* 退避重试：重试 3 次，每次重试的时间间隔为介于 10s~20s 的随机值。
* 指数衰减重试：默认重试策略。重试 176 次，每次重试的时间间隔按照指数递增至 512s，总计重试时间为 24 小时，即重试时间间隔为 1s、2s、4s、8s、16s、32s、64s、128s、256s、512s……512s（共计 167 次间隔 512s）。

### 容错策略

当错误发生时的处理方式：

* 允许容错：请求失败且重试失败后，跳过此请求，继续处理下一条请求。
* 禁止容错：请求失败且重试失败后，消费任务阻塞。（目前禁止容错有 bug）
* 死信队列：仅当开启允许容错时，可配置死信队列。
  * 如果启用死信队列，未被处理或超过重试次数的消息会被投递到死信队列中。
  * 如果未启用死信队列，超过重试次数的消息会被丢弃。

### 重试和容错最佳实践

TBD

# 参见

* [Kafka 触发器](https://help.aliyun.com/zh/fc/apsaramq-for-kafka-trigger)
* [触发器高级功能](https://help.aliyun.com/zh/fc/user-guide/advanced-features-of-triggers)
* [触发器 Event 格式](https://help.aliyun.com/zh/fc/user-guide/formats-of-event-for-different-triggers)
* [事件总线 EventBridge - 操作指南 - 事件流 - 事件源 - 消息队列 Kafka 版](https://help.aliyun.com/document_detail/439526.html)
* [triggers 字段](https://docs.serverless-devs.com/fc/yaml/triggers)