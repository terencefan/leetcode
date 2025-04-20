# 

## Pod

Our service is an Azure service, consisting of two components: SignalR and Web PubSub, both built on top of WebSocket protocols.

SignalR is part of Microsoft's .NET Framework and functions as an RPC protocol between clients (such as mobile apps, browsers, or any JavaScript-enabled devices). It supports bidirectional invocations, which is one of our primary use cases. For example, a client can invoke predefined methods on its app server using this protocol, and the server can invoke methods on the client in return.

Our SignalR service is built on top of this protocol and acts as a proxy in this scenario. Our customers, primarily developers, use the server SDK we provide to establish connections to our service. Clients first negotiate the connection request with the app server, which redirects the request to one of our instances. Depending on whether the instance has connections to the desired server, it may redirect the request to another instance. Once this process is complete, we respond with a 101 Switching Protocol status to the client, upgrading the HTTP request to a WebSocket connection.

In addition to supporting basic client/server invocations, we also provide RESTful APIs for our customers. These APIs allow them to add or remove connections from groups or attach a connection to a user. Using the group/user to connection mapping in redis and the route table shared across our instances, the server can send messages to any connection, even if the connection is not on the same server. It can also broadcast messages to all connections, or connections within a group, or connections associated with a user (a user is treated as a special type of group). This functionality is achieved using Redis channels, which I can explain in more detail if needed.

As mentioned earlier, the SignalR protocol is basically an RPC protocol built on the WebSocket protocol. We have also introduced features like HTTP upstream invocation to further lower the develop costs for our customers to build real-time systems. This part can be an HTTP server instead of WebSocket connections or operate in a completely serverless manner. Our service provides connection management and message-sending methods for clients to invoke, which are largely similar to our REST APIs.



## Redis

Using redis to distribute messages is one of our bottleneck, depends on the customer's behavior, this bottleneck can be in either inbound or outbound. 

### 入口流量 Inbound

When the customer create unique groups for each connection and send messages to these groups, or they send messages directly to other connections that are not on the same instance, the message has to be forwared to other instances through redis. Consider if the redis server has resources has more than 100 units in total connected to it, which in our case, the maximum connection count can be 100K. If the customer send message to a connections in 10 seconds interval, the QPS will be 10K, which nearly reaches the capacity of a single redis instance can handle. Considering the message pub is an expansive operation rather than simple get/set, the real quota could be even less. We solve this problem by introducing client-side sharding algorithm, the pub requests to redis will be partitioned by the channel name so the inbound messages can be distributed almost evenly to backend redis instances. The reason I say almost is our consistent hash algorithm is based on calculation rather than sync, to reduce the cost of development and maintenance. 

### 出口流量 Outbound

如果用户的 group 中存在大量连接，且服务端向 group 中广播消息，出口流量的规模会按照乘方增长。Redis 在这一场景中不会承载压力，因为实际订阅某一 group 的客户端数量 <20 个（极端情况下我们也只会有 20 个 pod）。但同一 pod 上可能会有数千连接需要并发写出，我们通过以下四个办法解决该问题：

1. 出口流量预测
3. 无锁队列转移 I/O 到单线程处理
3. cmap + 客户端引用链表 + 压缩
4. 多消息压缩为大消息

### 故障转移 Failover

一个逻辑 Redis client 由两个 client 组成，分别连接到 primary 和 secondary redis 实例，作为一致性 hash 的一部分。实际操作中我们把哈希槽看作一个环形，secondary 一般是当前槽位对面槽位所对应的实例。

我们只针对 Premium Tier 的客户开启此 feature，用户的后段资源相对独立且分布均匀，目前并不存在某一主节点崩溃后的雪崩问题，如有，我们可以通过实现类似 BRR 的客户端拥塞控制算法进行解决，在 Redis 的 RTT 过高时对客户端进行限流（未实施）。

A logical redis client will subscribe the channel on 2 different redis instance, primary/secondary.

## Resource limitation based on CGroup and standalone GC

Another bottleneck is memory usages. For resources that has only 1 unit, the cpu/memory quota should be 500mc/512MB. It is obviously that we could not let a customer who requested only a few units to use all resources assigned to a pod, which in our case is 6c/14G. So we offer shared units for small customers, which we called the standard Tier. The resource limit can vary between 1~100 units. We use soft resource isolation in the very early stages, like 1000 connections / 4M iops per unit

- Pod container orchestration (Golang)
- WebSocket and HTTP service runtime (C#)

## Questions:

Q: Do you support streams?
A: The original SignalR protcol supports streams, so in theory we support. But we charge in message count and size, use our service with streaming can be extremely expansive so I think there is no actual usages. Customers should find peer-to-peer solutions based on QUIC or other network protocols and use our service solely as a signaling tool.