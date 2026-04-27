# ComfyUI

## Node

| Name             | Description                           | ddd |
| ---------------- | ------------------------------------- | --- |
| ChekcpointLoader | 选择模型文件                          |     |
| KSampler         | 核心采样器                            |     |
| CLIP Text Encode | 对 Prompt 进行编码，生成向量          |     |
| VAE Encode       | 将图像编码为 Latent Space，用于图生图 |
| VAE Decode       | 将 Latent Space 解码为最终图像        |     |

## KSampler

采样节点

从随机噪声中逐步生成高质量的图像

### Latent Space