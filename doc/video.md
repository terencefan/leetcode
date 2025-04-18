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
  - 5ms - 200ms 可调 fram   
  - 自适应码率 (6 kb/s to 510 kb/s)
- AAC
- G.711
- EVS