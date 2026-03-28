# 🚀 AI 学习指南 | AI Learning Guide

[![](https://img.shields.io/github/stars/ibarryyan/AI-learning.svg?style=flat)](https://github.com/ibarryyan/AI-learning/stargazers)
![Static Badge](https://img.shields.io/badge/AI%20Learning-闫同学-8A2BE2)
<a href="https://github.com/ibarryyan/golang-tips-100/blob/master/img/wechat.jpg"><img src="https://img.shields.io/badge/%E5%85%AC%E4%BC%97%E5%8F%B7-%E6%89%AF%E7%BC%96%E7%A8%8B%E7%9A%84%E6%B7%A1-blue" alt="公众号"></a>

> 全面、系统的 AI 学习资源汇总。无论你是 AI 领域的新手还是有一定经验的开发者，这个仓库都能帮助你更好地学习和应用 AI 技术。

---

## 📋 目录导航

- [📖 项目概览](#-项目概览)
- [🚀 快速开始](#-快速开始)
- [🤖 开源大模型](#-开源大模型)
- [🛠️ 开发工具](#️-开发工具)
- [📚 学习教程](#-学习教程)
- [📄 核心论文](#-核心论文)
- [💡 应用产品](#-应用产品)
- [📖 博客资源](#-博客资源)
- [🗂️ 项目仓库](#️-项目仓库)
- [🔧 专题深入](#-专题深入)
- [📎 联系作者](#-联系作者)

---

## 📖 项目概览

本仓库收集和整理与 AI 学习相关的资源，涵盖：

| 分类 | 描述 |
|------|------|
| 🤖 开源大模型 | LLaMA、DeepSeek、Qwen 等主流开源大语言模型 |
| 🛠️ 开发工具 | PyTorch、Transformers、vLLM 等 AI 开发框架 |
| 📚 学习教程 | 微软、Stanford、Hugging Face 等优质教程 |
| 📄 核心论文 | Transformer、BERT、GPT 等奠基性论文 |
| 💡 应用产品 | ChatGPT、DeepSeek、文心一言等产品 |
| 📖 博客资源 | 实战教程、入门指南、部署经验 |
| 🗂️ 项目仓库 | LangChain、LLaMA-Factory 等热门开源项目 |

---

## 🚀 快速开始

### 🤖 想要使用大模型？

| 类型 | 推荐 |
|------|------|
| 🔥 最强开源 | [DeepSeek-V3](https://chat.deepseek.com/) · [Qwen](https://tongyi.aliyun.com/) |
| 🌐 通用对话 | [ChatGPT](https://chat.openai.com/) · [Claude](https://claude.ai/) |
| 🇨🇳 中文优化 | [文心一言](https://yiyan.baidu.com/) · [智谱清言](https://chatglm.cn/) |

### 💻 想要本地部署？

```bash
# 使用 LLaMA-Factory 微调模型
git clone https://github.com/hiyouga/LLaMA-Factory
cd LLaMA-Factory
pip install -r requirements.txt

# 使用 vLLM 部署推理服务
pip install vllm
vllm serve Qwen/Qwen2-7B-Instruct
```

### 📚 想要系统学习？

| 水平 | 推荐资源 |
|------|----------|
| 🟢 初学者 | [微软 AI for Beginners](./tutorial/README.md) · [零基础 AI 入门](./blogs/README.md) |
| 🟡 进阶 | [LLM Cookbook](./tutorial/README.md) · [CS224n](./tutorial/README.md) |
| 🔴 深入 | [DeepSeek 论文](./paper/README.md) · [Transformer 家族](./paper/README.md) |

---

## 🤖 开源大模型

| 模型 | 链接 | 简介 | 推荐 |
|:-----|:-----|:-----|:----:|
| **DeepSeek-V3** | [GitHub](https://github.com/deepseek-ai/DeepSeek-V3) | 深度求索开源最强对话模型，长文本理解与代码生成能力强 | ⭐⭐⭐⭐⭐ |
| **DeepSeek-R1** | [GitHub](https://github.com/deepseek-ai/DeepSeek-R1) | 通过强化学习激励推理能力，数学和逻辑能力出色 | ⭐⭐⭐⭐⭐ |
| **Qwen2.5** | [GitHub](https://github.com/QwenLM/Qwen) | 阿里开源通义千问，中文优化出色的多系列模型 | ⭐⭐⭐⭐⭐ |
| **LLaMA 3** | [GitHub](https://github.com/facebookresearch/llama) | Meta 开源，羊驼家族最新力作 | ⭐⭐⭐⭐ |
| **DeepSeek-MoE** | [GitHub](https://github.com/deepseek-ai/DeepSeek-MoE) | 混合专家架构，推理成本低 | ⭐⭐⭐⭐ |
| **Vicuna** | [GitHub](https://github.com/lm-sys/FastChat) | 基于 LLaMA 的对话优化模型 | ⭐⭐⭐⭐ |
| **ChatGLM** | [GitHub](https://github.com/THUDM/ChatGLM-6B) | 清华开源中英双语对话模型 | ⭐⭐⭐⭐ |
| **Mistral** | [GitHub](https://github.com/mistralai/mistral-src) | 欧洲最强开源模型 | ⭐⭐⭐⭐ |

> 💡 更多模型对比和技术细节请查看 [AI Coding IDE 专题](./docs/ai-coding-ide.md)

---

## 🛠️ 开发工具

### 深度学习框架

| 工具 | 链接 | 简介 | 推荐 |
|:-----|:-----|:-----|:----:|
| **PyTorch** | [官网](https://pytorch.org/) | Facebook 开发，灵活性高，适合研究 | ⭐⭐⭐⭐⭐ |
| **TensorFlow** | [官网](https://www.tensorflow.org/) | Google 开发，生态完善 | ⭐⭐⭐⭐ |
| **JAX** | [官网](https://jax.readthedocs.io/) | Google 高性能 ML 框架 | ⭐⭐⭐⭐ |

### 模型与推理

| 工具 | 链接 | 简介 | 推荐 |
|:-----|:-----|:-----|:----:|
| **Hugging Face Transformers** | [GitHub](https://github.com/huggingface/transformers) | 提供丰富预训练模型 | ⭐⭐⭐⭐⭐ |
| **vLLM** | [GitHub](https://github.com/vllm-project/vllm) | 高性能 LLM 推理框架 | ⭐⭐⭐⭐⭐ |
| **LLaMA-Factory** | [GitHub](https://github.com/hiyouga/LLaMA-Factory) | 一站式模型微调工具 | ⭐⭐⭐⭐⭐ |
| **DeepSpeed** | [GitHub](https://github.com/microsoft/DeepSpeed) | 微软分布式训练加速 | ⭐⭐⭐⭐ |
| **Ollama** | [官网](https://ollama.com/) | 本地大模型运行工具 | ⭐⭐⭐⭐⭐ |
| **LangChain** | [GitHub](https://github.com/langchain-ai/langchain) | LLM 应用开发框架 | ⭐⭐⭐⭐ |

### 计算机视觉

| 工具 | 链接 | 简介 | 推荐 |
|:-----|:-----|:-----|:----:|
| **OpenCV** | [官网](https://opencv.org/) | 计算机视觉基础库 | ⭐⭐⭐⭐ |
| **Pillow** | [官网](https://pillow.readthedocs.io/) | Python 图像处理 | ⭐⭐⭐⭐ |

---

## 📚 学习教程

| 教程 | 链接 | 简介 | 推荐 |
|:-----|:-----|:-----|:----:|
| **微软 AI for Beginners** | [GitHub](https://github.com/microsoft/AI-For-Beginners) | 面向初学者的系统课程 | ⭐⭐⭐⭐ |
| **LLM Cookbook** | [GitHub](https://github.com/datawhalechina/llm-cookbook) | DataWhale 大模型教程 | ⭐⭐⭐⭐⭐ |
| **Hugging Face 教程** | [官网](https://huggingface.co/learn) | NLP 模型使用与训练 | ⭐⭐⭐⭐ |
| **Stanford CS224n** | [课程](http://web.stanford.edu/class/cs224n/) | 斯坦福 NLP 课程 | ⭐⭐⭐⭐⭐ |
| **PyTorch 官方教程** | [官网](https://pytorch.org/tutorials/) | 深度学习入门必读 | ⭐⭐⭐⭐ |

更多教程资源请查看 [Tutorial](./tutorial/README.md) 专辑。

---

## 📄 核心论文

### 奠基之作

| 论文 | 链接 | 简介 | 年份 |
|:-----|:-----|:-----|:----:|
| **Attention Is All You Need** | [arxiv](https://arxiv.org/abs/1706.03762) | Transformer 架构奠基论文 | 2017 |
| **BERT** | [arxiv](https://arxiv.org/abs/1810.04805) | 预训练语言模型里程碑 | 2018 |
| **GPT** | [arxiv](https://arxiv.org/abs/1801.04586) | 生成式预训练先驱 | 2018 |
| **GPT-3** | [arxiv](https://arxiv.org/abs/2005.14165) | 提示学习的里程碑 | 2020 |

### 2024-2025 最新进展

| 论文 | 链接 | 简介 | 年份 |
|:-----|:-----|:-----|:----:|
| **DeepSeek-V3** | [arxiv](https://arxiv.org/abs/2412.19437) | 超强开源 MoE 模型 | 2024 |
| **DeepSeek-R1** | [arxiv](https://arxiv.org/abs/2501.12948) | 纯强化学习推理模型 | 2025 |
| **Qwen2.5** | [arxiv](https://arxiv.org/abs/2412.15115) | 阿里新一代开源模型 | 2024 |
| **LLaMA 3** | [arxiv](https://arxiv.org/abs/2407.21783) | Meta 开源 Llama 最新版 | 2024 |

更多论文请查看 [Papers](./paper/README.md) 专辑。

---

## 💡 应用产品

### 通用对话

| 产品 | 链接 | 特色 | 推荐 |
|:-----|:-----|:-----|:----:|
| **DeepSeek** | [官网](https://chat.deepseek.com/) | 长文本理解、代码生成强 | ⭐⭐⭐⭐⭐ |
| **ChatGPT** | [官网](https://chat.openai.com/) | OpenAI 通用对话 | ⭐⭐⭐⭐⭐ |
| **Claude** | [官网](https://claude.ai/) | Anthropic 出品，更安全 | ⭐⭐⭐⭐⭐ |
| **文心一言** | [官网](https://yiyan.baidu.com/) | 百度多模态大模型 | ⭐⭐⭐⭐ |
| **通义千问** | [官网](https://tongyi.aliyun.com/) | 阿里云行业大模型 | ⭐⭐⭐⭐ |
| **智谱清言** | [官网](https://chatglm.cn/) | ChatGLM 商业版 | ⭐⭐⭐⭐ |

### AI 编程

| 产品 | 链接 | 特色 | 推荐 |
|:-----|:-----|:-----|:----:|
| **Cursor** | [官网](https://cursor.sh/) | AI 增强版 VS Code | ⭐⭐⭐⭐⭐ |
| **Windsurf** | [官网](https://codeium.com/windsurf) | AI 编程新秀 | ⭐⭐⭐⭐⭐ |
| **GitHub Copilot** | [官网](https://github.com/features/copilot) | GitHub 官方 AI 编程 | ⭐⭐⭐⭐ |
| **通义灵码** | [官网](https://tongyi.aliyun.com/lingma) | 阿里 AI 编程助手 | ⭐⭐⭐⭐ |

### AI 图像

| 产品 | 链接 | 特色 | 推荐 |
|:-----|:-----|:-----|:----:|
| **Midjourney** | [官网](https://www.midjourney.com/) | 艺术级 AI 绘画 | ⭐⭐⭐⭐⭐ |
| **Stable Diffusion** | [官网](https://stability.ai/) | 开源图像生成 | ⭐⭐⭐⭐⭐ |
| **DALL-E** | [官网](https://openai.com/dall-e-3/) | OpenAI 图像生成 | ⭐⭐⭐⭐ |
| **通义万相** | [官网](https://wanxiang.aliyun.com/) | 阿里 AI 绘画 | ⭐⭐⭐⭐ |

更多产品请查看 [Product](./product/README.md) 专辑。

---

## 📖 博客资源

| 文章 | 链接 | 简介 | 推荐 |
|:-----|:-----|:-----|:----:|
| 普通人如何使用 AI？ | [CSDN](https://blog.csdn.net/z551646/article/details/143848325) | AI 基本概念和日常使用 | ⭐⭐⭐ |
| AI 赋能第一次使用 AI | [知乎](https://zhuanlan.zhihu.com/p/684988103) | 保姆级 AI 使用教程 | ⭐⭐⭐ |
| 高效使用 AI 的 21 种技巧 | [知乎](https://zhuanlan.zhihu.com/p/14780214495) | AI 提示词技巧汇总 | ⭐⭐⭐ |
| DeepSeek R1 本地部署 | [CSDN](https://blog.csdn.net/zhishi0000/article/details/145450612) | 详细本地部署教程 | ⭐⭐⭐⭐ |
| 零基础 AI 入门指南 | [廖雪峰](https://www.liaoxuefeng.com/article/1543329456062498) | 适合新手的入门指南 | ⭐⭐⭐ |
| DeepSeek v3 搭建知识库 | [CSDN](https://blog.csdn.net/python12345_/article/details/145450272) | 企业级私有知识库 | ⭐⭐⭐ |

更多博客请查看 [Blogs](./blogs/README.md) 专辑。

---

## 🗂️ 项目仓库

| 项目 | 链接 | 简介 | 推荐 |
|:-----|:-----|:-----|:----:|
| **HuggingFace Transformers** | [GitHub](https://github.com/huggingface/transformers) | 主流模型 PyTorch 实现 | ⭐⭐⭐⭐⭐ |
| **LangChain** | [GitHub](https://github.com/langchain-ai/langchain) | LLM 应用开发框架 | ⭐⭐⭐⭐⭐ |
| **LLaMA-Factory** | [GitHub](https://github.com/hiyouga/LLaMA-Factory) | 一站式微调工具 | ⭐⭐⭐⭐⭐ |
| **vLLM** | [GitHub](https://github.com/vllm-project/vllm) | 高性能推理框架 | ⭐⭐⭐⭐⭐ |
| **DeepSpeed** | [GitHub](https://github.com/microsoft/DeepSpeed) | 分布式训练加速 | ⭐⭐⭐⭐ |
| **Prompt Engineering Guide** | [GitHub](https://github.com/dair-ai/Prompt-Engineering-Guide) | 提示工程指南 | ⭐⭐⭐⭐ |
| **OpenAI Cookbook** | [GitHub](https://github.com/openai/openai-cookbook) | OpenAI 官方示例 | ⭐⭐⭐⭐ |
| **Chinese-LLaMA-Alpaca** | [GitHub](https://github.com/ymcui/Chinese-LLaMA-Alpaca) | 中文 LLaMA 增强 | ⭐⭐⭐⭐ |

更多仓库请查看 [Repository](./repository/README.md) 专辑。

---

## 🔧 专题深入

| 专题 | 描述 | 路径 |
|:-----|:-----|:-----|
| 🤖 **AI 学习路线** | 系统化的 AI 学习路径规划 | [查看](./docs/ai-learning-roadmap.md) |
| 💻 **AI Coding IDE** | 主流 AI 编程 IDE 对比 | [查看](./docs/ai-coding-ide.md) |
| 🖥️ **GPU 知识** | GPU 原理与环境配置指南 | [查看](./docs/gpu-knowledge.md) |
| 🛠️ **AI 技术栈** | 完整的技术栈文档 | [查看](./docs/ai-tech-stack.md) |
| 🎨 **IDE AI 插件** | 主流 IDE AI 插件收集 | [查看](./docs/ide-ai-plugins.md) |
| 🐚 **OpenClaw** | OpenClaw 使用指南 | [查看](./docs/openclaw.md) |
| 🎯 **AI 工具补充** | UI 生成等工具 | [查看](./docs/ai-tools-supplement.md) |

---

## 📎 联系作者

| 联系方式 | 详情 |
|:---------|:-----|
| 📧 E-mail | yanmingxin.boy@gmail.com |
| 💬 WeChat | 扯编程的淡 |

<img src="assets/wx.png" title="" alt="公众号二维码" width="246">

---

> ⭐ 如果这个项目对你有帮助，请 star 支持一下！
