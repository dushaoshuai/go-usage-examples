函数计算运行时

# 基本信息

运行时即函数执行环境，提供管理、运行函数的安全、隔离的环境。函数计算调用函数时，会分配一个执行环境。

每个运行时都会基于一个 Linux 发行版本制作，目前支持 Debian 9（Stretch）和 Debian 10（Buster）。

函数计算提供了内置运行时、自定义运行时和自定义容器运行时，以满足不同需求：

| 对比项  | 内置运行时                         | 自定义运行时                             | 自定义容器运行时                         |
|------|-------------------------------|------------------------------------|----------------------------------| 
| 适用场景 | 按照函数计算定义的接口编写程序处理事件和 HTTP 请求。 | 基于各个语言的流行框架编写程序，或者迁移已有的框架应用。       | 完全控制程序运行的环境，或者迁移已有的容器应用；使用GPU实例。 |
| 冷启动  | 最快。代码包中不包含运行时，所以冷启动最快。        | 较快。自定义运行时使用公共镜像，没有镜像拉取时间，所以冷启动会较快。 | 较慢。需要拉取镜像，所以冷启动较慢。               |
| ...  | ...                           | ...                                | ...                              |

# 修改运行时环境

可以通过设置层或环境变量来修改运行时环境。

## 修改时区

默认是 UTC 时间，也就是 0 时区，设置环境变量 TZ 的值为 Asia/Shanghai 后，时区被修改为东 8 区，即北京时间。

# 参见

* [函数运行时选型](https://help.aliyun.com/zh/fc/product-overview/function-runtime-selection)
* [运行时](https://help.aliyun.com/zh/fc/user-guide/runtimes/)
* [代码开发概述](https://help.aliyun.com/zh/fc/user-guide/overview-35)
* [Custom Runtime](https://help.aliyun.com/zh/fc/user-guide/custom-runtime/)
* [Custom Container](https://help.aliyun.com/zh/fc/user-guide/custom-container/)
