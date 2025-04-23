# Concepts

## QUIC

RFC 9000

https://quicwg.org/

https://www.cse.wustl.edu/~jain/cse570-21/ftp/quic/index.html

- 0 RTT 连接恢复
- 多路复用，无队头阻塞
  - HTTP/2 虽然支持多 stream，但重传会导致阻塞
  - QUIC 独立 stream
- 规避缓冲区/拥塞控制限制？
- 连接标识符，网络切换无需重连

## HTTP/3 over QUIC

> HTTP/3 -> QUIC (TLS 1.3) -> UDP -> IP

TCP + TLS: 1(TCP) + 1~2(TLS) RTT

- TLS 1.2 (2 RTT)

## NAT 穿透

- 完全锥形
- 受限锥形
- 端口受限锥形
- 对称型

## OPUS

FEC 优先，PLC 兜底

- G.711 延迟更低？(0.125ms)
- 6kbps~510kbps 之间之间动态调整码率

- 分割成 20ms 的帧，50hz
- 码率自适应

- FEC (Forward Error Correction) 发送冗余包
  - XOR（单帧），A ^ B ^ C = F
  - Reed - Solomon（RS码）
  - 喷泉码 - Raptor
  - 低丢包 RS，高丢包 Fountain
  - 独立媒体编码（添加冗余包）
- PLC (Package Loss Concealment) 算法隐藏
  - LPC 线性插帧
  - WSLOA 拉伸
- VBR, 可变码率
  - 主动探测音频信号复杂度
  - 根据 RMS 值识别音频能量（静音检测，音量调整）
- VAD，语音活动检测
  - 静音/白噪音降低码率
- 动态切换
  - SILK 模式（语音）
  - CELT 模式（音乐）
- RTCP 反馈

### SRT 协议（高丢包网络）

- 直播推流
- 自带加密
- 高丢包，ARQ，内置 NACK 机制
- 含拥塞控制

### RTP 协议（低丢包网络）

- 序列号
- Timestamp
- SSRC（StreamID)

## RTCP

- RTP 的控制协议
- 秒级发送间隔
- QoS 监控，丢包率/抖动/延迟等网络指标
  - SR（Sender Report）
    - 结合 NTP 时间戳估算 RTT
  - RR（Receiver Report）
    - 调整码率或开启 FEC

### 加密（SRTP）

- 类似 TLS，密钥交换（AES
- 特性小包高频
- 头部开销小

### 高丢包网络

1. FEC
   1. In-Band FEC (SILK) 嵌入前一帧的低码率版本
   2. FlexFEC，利用 RTP 冗余包修复连续丢包
2. PLC
   1. WSOLA，复制相似帧
   2. CNG 舒适噪音
3. RTCP 自适应码率
   1. 根据 RR 反馈，在 6~510kbps 间动态调整
   2. 带宽探测（指数退避） 

### 高抖动网络

1. Jitter Buffer
   1. 自适应缓冲，预留 20~200ms
   2. 连续乱序临时扩展缓冲容量
2. RTP 时间戳优化
   1. Opus Frame header 自带高精度时间戳，与 NTP 时钟对齐
   2. 乱序包重组
3. 弹性帧处理
   1. 抖动高时使用短帧，降低丢包影响
   2. 抖动低时使用长帧，提高编码效率
   3. 动态分裂与合并帧

- 客户端 PLC (Package Loss Concealment) 算法
  - 短帧 LPC（预测），长帧 WSLOA（拉伸）
  - CELP？

- WSOLA (Waveform Similarity based Overlap-Add) 应用场景？
  - PLC（丢包隐藏）
  - 语音变速不变调


## 语音 vs 音乐

- 窄带 vs 全频带
- 浊音占比高，而浊音具有高度周期性
- 声带单一频率 vs 多乐器/和弦频率叠加
- 瞬态成分，爆破音 vs 打击乐

结论：语音高度周期性，音乐因复杂叠加周期性弱

## SRT

- RTMS，延迟 1~3s
- 支持 ARQ，但使用 UDP 传输，解决队头阻塞问题
- Jitter buffer，自适应缓冲，通过时间戳对齐数据
  - 高抖动网络增大缓冲区
- PLC（丢包隐藏）
  - repeat, 机器人声音
  - LPC（算法预测）
  - WSLOA（时间缩放），长时丢包
  - AI 生成

## WebRTC

P2P communication

- Signaling server (websocket)

- NAT 穿透
  - STUN - NAT Hole Punching
  - TURN
    - https://github.com/coturn/coturn
  - ICE (STUN 后降级为 TURN)
  - UPnP/IGD（自动端口映射）

## Audio codecs

https://opus-codec.org/

- Opus
  - 5ms - 200ms 可调 frame 
  - 自适应码率 (6 kb/s to 510 kb/s)
- AAC
- G.711
- EVS