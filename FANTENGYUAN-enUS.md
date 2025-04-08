---
stylesheet: https://cdnjs.cloudflare.com/ajax/libs/github-markdown-css/2.10.0/github-markdown.min.css
body_class: markdown-body
css: |-
  .page-break { page-break-after: always; }
  .markdown-body { font-size: 10px; }
  .markdown-body pre > code { white-space: pre-wrap; }
pdf_options:
  format: A4
  margin: 10mm
  printBackground: true
---

# Fan Tengyuan

Shanghai, China | (+86) 15216709049 | stdrickforce@gmail.com

LinkedIn: https://linkedin.com/in/inority | Github: https://github.com/terencefan

**Shanghai Jiao Tong University​​**

Bachelor of Software Engineering | Sep 2009 - Jun 2013

## Technical skills

**Programming languages**

C# • Golang • Java • C

**Open-Source tools && frameworks**

Kubernetes (K8s) • Redis • Thrift • ElasticSearch • MySQL

**Frontend development**

Node.js • TypeScript • React

## Experience

### Microsoft - Azure

- Software Engineer II
- ​​August 2019 – Present​​

#### Azure SignalR / Web PubSub service
  
**SignalR**​​ is an application-layer communication protocol based on WebSocket developed and maintained by Microsoft. As one of the components of .NET Core, it is designed to help engineers deal with real-time communication challenges between servers and clients.

​**​Azure SignalR​ service​** is a cloud service based on this protocol. By using our service, engineers no longer need to independently maintain long-lived connections between clients and servers, and can quickly leverage Azure cloud service features such as cross-region replicas and auto-scaling.

Additionally, we have developed another cloud service based on the native WebSocket protocol, which supports multiple subprotocols such as ​​MQTT​​ and ​​Socket.io​​, known as the ​​**Azure Web PubSub​​ service**.

​​My responsibilities in this team include:​​

- Lead the design of server and client protocols for optimizing connection stability and rebalance connections across container pods.

- Own the design, implementation, and maintenance of JWT token and OAuth 2.0 based authentication modules.

- Participate in the development of multi-language client/server SDKs, such as ​​Python​​ and ​​Java​​.

- Handle other routine tasks, including localization, logging, monitoring, and on-call duties.

My primary programming language is **C#** in this project, and frequently work with **Redis** and **Kubernetes (K8s)**. Additionally, I contribute to frontend projects using **Typescript/ReactJS** and develop tools and middlewares using **Golang**.

Our server-side SDKs are open-source projects, available at:

https://github.com/Azure/azure-signalr • https://github.com/Azure/azure-webpubsub

<div class="page-break"></div>

### Meituan-Dianping - Infrastructure

- Software Engineer
- Jan 2018 - Aug 2019​

#### CAT Monitoring Platform Development

​​CAT​​ is a distributed real-time monitoring system developed by Meituan-Dianping and is currently open source and widely-used.

It also deeply integrated with Meituan-Dianping's core middlewares (such as ​​MVC, RPC, SQL, Redis, Kafka, Zookeeper​​, etc.), it provides ​​performance metrics, health status, and real-time alerts​​ for all business units across the organization.

​​My responsibilities in this team include:​​

- Own the development of ​​multi-language client-side instrumentation SDKs​​.
- Own the design and development of a centralized monitoring dashboard​​ with multi-dimensional data views.
- Lead the development of the data collection system​​.
- Contribute to the development of a ​​real-time alerting system​​ with second-level latency

My primary programming language is **Java** in this project, and frequently work with **Kafka**, **Hadoop**, **Elasticsearch**.
Additionally, I use **C** and **Golang** to develop cross-platform client instrumentation SDKs for other programming languages.

We have a open source version on Github, available at:
https://github.com/dianping/cat