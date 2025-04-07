# 范腾远

上海 ｜ (+86) 15216709049 | stdrickforce@gmail.com

LinkedIn: https://linkedin.com/in/inority ｜ Github: https://github.com/terencefan

## 教育背景

上海交通大学 ｜ 软件工程 ｜ 学士 ｜ 2009.9 ~ 2013.06

## 技术技能

- 编程语言: C# • Golang • Java • C

- 开源工具: Kubernetes (K8s) • Redis • Thrift • ElasticSearch • MySQL

- 前端开发: Node.js • TypeScript • React

- 英语: 流畅阅读技术文档，具备口语和书面交流能力

## 工作经验

### 微软亚太研发中心 - Azure 云计算

- Software Engineer II
- 2019.8 至今

#### Azure SignalR / Web PubSub 云服务开发
  
**SignalR** 是由 Microsoft 研发并维护的基于 websocket 的应用层通讯协议，现作为 .Net Core 的核心组件之一，为服务器和客户端的实时通信提供解决方案。

**Azure SignalR** 是微软基于该协议提供的云服务产品，我们致力于减少客户开发实时应用通讯产品的成本，通过使用我们的服务，用户无需独立维护客户端到服务端的长连接，并可快速拥有跨区域部署 / 自动扩容等 Azure 云服务的特性。

此外，我们还基于 websocket 原生协议开发了长连接服务，并兼容 MQTT / Socket.io 等多种协议，即 **Azure Web PubSub** 服务。

我在本团队的工作内容：

- 主导设计服务端和客户端协议，优化连接稳定性和解决负载不均问题
- 主导实现并维护基于 JWT token 和 OAuth 2.0 的认证模块
- 协助开发多语言 SDK，如 Python / Java
- 其他日常工作，如本地化、日志、监控、oncall 等

在本项目中，我主要使用 C# 作为开发语言，并使用 Redis / K8s 等工具，同时使用 TypeScript/ReactJS 进行前端开发，以及使用 Golang 进行工具和中间件开发。

我们的服务端 SDK 是开源项目，项目地址：

https://github.com/Azure/azure-signalr • https://github.com/Azure/azure-webpubsub

### 美团点评 - 基础研发平台

- 软件开发工程师
- 2018.1 ~ 2019.8

#### 监控系统中间件（CAT）开发

CAT 是由大众点评开发，现由美团点评维护的基于客户端埋点的分布式监控系统，在美团点评的基础中间件框架（如 MVC / RPC / SQL / Redis / Kafka / Zookeeper 等）深度集成，为美团点评各业务线提供性能指标、健康状况、实时告警等。

我在本团队的工作内容：

- 主导开发数据收集系统，支持百万级 QPS
- 主导开发多语言客户端埋点工具
- 主导开发监控大盘产品
- 协助开发秒级告警系统
- 其他日常维护工作，如协助其他部门排查问题

在本团队中，我主要使用 Java 作为开发语言，并使用 Kafka / Hadoop / MySQL 等工具，同时使用 C / Python / Golang 等语言开发客户端埋点 SDK，以及 JavaScript + ECharts 进行简单的前端开发。

我们有开源版本的服务端，项目地址：https://github.com/dianping/cat

### 其他

- 百姓网：基于 PHP6 + swoole + Thrift 开发后台系统

- 小红书：基于 Python 开发社区后台系统和推送系统

- 饿了么：基于 Python + Thrift 协议开发 RPC 服务，基于 AngularJS 开发配送后台系统