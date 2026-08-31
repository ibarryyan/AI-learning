# 🚀 AI 学习指南 | AI Learning Guide

[![](https://img.shields.io/github/stars/ibarryyan/AI-learning.svg?style=flat)](https://github.com/ibarryyan/AI-learning/stargazers)
![Static Badge](https://img.shields.io/badge/AI%20Learning-闫同学-8A2BE2)
<a href="https://github.com/ibarryyan/golang-tips-100/blob/master/img/wechat.jpg"><img src="https://img.shields.io/badge/%E5%85%AC%E4%BC%97%E5%8F%B7-%E6%89%AF%E7%BC%96%E7%A8%8B%E7%9A%84%E6%B7%A1-blue" alt="公众号"></a>

> 全面、系统的 AI 学习资源汇总，内容对齐 2026 年业界最新进展。无论你是刚接触 AI 的新手，还是有一定经验的开发者，这个仓库都能帮你少走弯路。

---

## 📋 目录导航

- [📖 项目概览](#-项目概览)
- [🚀 快速开始](#-快速开始)
- [🤖 开源大模型](#-开源大模型)
- [🛠️ 开发工具](#️-开发工具)
- [📚 学习教程](#-学习教程)
- [📄 核心论文](#-核心论文)
- [💡 应用产品](#-应用产品)
- [🤖 Agent 专辑](#-agent-专辑)
- [📖 博客资源](#-博客资源)
- [🗂️ 项目仓库](#️-项目仓库)
- [🔧 专题深入](#-专题深入)
- [📎 联系作者](#-联系作者)

---

## 📖 项目概览

本仓库收集和整理与 AI 学习相关的资源，涵盖以下几类。

| 分类 | 描述 |
|------|------|
| 🤖 开源大模型 | DeepSeek、Qwen3、Kimi K2、GPT-OSS、Llama 4 等主流开源大语言模型 |
| 🛠️ 开发工具 | PyTorch、Transformers、vLLM、SGLang、LLaMA-Factory 等 AI 开发框架 |
| 📚 学习教程 | 微软、Stanford、Hugging Face 等优质教程 |
| 📄 核心论文 | Transformer、BERT、GPT 等奠基性论文与最新进展 |
| 💡 应用产品 | GPT-5、Claude、Gemini 3、DeepSeek 等产品与编程 Agent |
| 🤖 Agent 专辑 | LangGraph、MCP、Claude Code、Dify 等 Agent 框架与平台 |
| 📖 博客资源 | 实战教程、入门指南、部署经验 |
| 🗂️ 项目仓库 | LangChain、LLaMA-Factory、MCP 等热门开源项目 |

---

## 🚀 快速开始

### 🤖 想要使用大模型？

| 类型 | 推荐 |
|------|------|
| 🔥 最强开源 | [DeepSeek](https://chat.deepseek.com/) · [通义千问](https://tongyi.aliyun.com/) · [Kimi](https://kimi.com/) |
| 🌐 闭源前沿 | [ChatGPT](https://chat.openai.com/) · [Claude](https://claude.ai/) · [Gemini](https://gemini.google.com/) |
| 🇨🇳 中文优化 | [文心一言](https://yiyan.baidu.com/) · [智谱清言](https://chatglm.cn/) · [豆包](https://www.doubao.com/) |

### 💻 想要本地部署？

```bash
# 使用 Ollama 一键本地运行开源模型
ollama run deepseek-r1:7b

# 使用 vLLM 部署高性能推理服务
pip install vllm
vllm serve Qwen/Qwen3-8B

# 使用 LLaMA-Factory 微调模型
git clone https://github.com/hiyouga/LLaMA-Factory
cd LLaMA-Factory
pip install -e ".[torch,metrics]"
```

### 📚 想要系统学习？

| 水平 | 推荐资源 |
|------|----------|
| 🟢 初学者 | [微软 AI for Beginners](./tutorial/README.md) · [零基础 AI 入门](./blogs/README.md) |
| 🟡 进阶 | [LLM Cookbook](./tutorial/README.md) · [CS224n](./tutorial/README.md) |
| 🔴 深入 | [DeepSeek 论文](./paper/README.md) · [Transformer 家族](./paper/README.md) |

---

## 🤖 开源大模型

2025 年以来开源模型全面走向 MoE 架构，单 H100 即可跑起前沿级模型，与闭源前沿的差距显著缩小。

| 模型 | 链接 | 简介 | 推荐 |
|:-----|:-----|:-----|:----:|
| **DeepSeek-V4** | [GitHub](https://github.com/deepseek-ai/DeepSeek-V3) | 深度求索最新一代开源模型，推理与 Agent 能力持续强化 | ⭐⭐⭐⭐⭐ |
| **DeepSeek-V3.2** | [GitHub](https://github.com/deepseek-ai/DeepSeek-V3) | 671B 总参 / 37B 激活，平衡推理与输出长度，对标 GPT-5 | ⭐⭐⭐⭐⭐ |
| **DeepSeek-R1** | [GitHub](https://github.com/deepseek-ai/DeepSeek-R1) | 纯强化学习激励推理能力，数学和逻辑表现出色 | ⭐⭐⭐⭐⭐ |
| **Qwen3-Max** | [GitHub](https://github.com/QwenLM/Qwen) | 阿里首个万亿参数旗舰，MoE 架构，代码与 Agent 能力突出 | ⭐⭐⭐⭐⭐ |
| **Qwen3** | [GitHub](https://github.com/QwenLM/Qwen) | 通义千问新一代全系列，覆盖 0.6B 到 235B，支持思考模式 | ⭐⭐⭐⭐⭐ |
| **Kimi K2 Thinking** | [GitHub](https://github.com/MoonshotAI/Kimi-K2) | 月之暗面 1T 总参 / 32B 激活，HLE 等推理榜单刷新 | ⭐⭐⭐⭐⭐ |
| **GPT-OSS** | [GitHub](https://github.com/openai/gpt-oss) | OpenAI 首批开源权重模型，120b 与 20b 两档，Apache 2.0 | ⭐⭐⭐⭐⭐ |
| **Llama 4** | [GitHub](https://github.com/meta-llama/llama) | Meta 首次采用 MoE，Scout 单卡可跑，原生 10M 上下文 | ⭐⭐⭐⭐ |
| **GLM-4.6** | [GitHub](https://github.com/zhipuai-ai/GLM-4.6) | 智谱开源对话与 Agent 模型，中英双语表现稳定 | ⭐⭐⭐⭐ |
| **Mistral** | [GitHub](https://github.com/mistralai/mistral-src) | 欧洲主流开源模型，轻量化部署友好 | ⭐⭐⭐⭐ |

> 💡 更多模型对比和技术细节请查看 [AI Coding IDE 专题](./docs/ai-coding-ide.md)

---

## 🛠️ 开发工具

### 深度学习框架

| 工具 | 链接 | 简介 | 推荐 |
|:-----|:-----|:-----|:----:|
| **PyTorch** | [官网](https://pytorch.org/) | Meta 主导，灵活性高，研究首选 | ⭐⭐⭐⭐⭐ |
| **JAX** | [官网](https://jax.readthedocs.io/) | Google 高性能 ML 框架，擅长并行与自动微分 | ⭐⭐⭐⭐ |
| **TensorFlow** | [官网](https://www.tensorflow.org/) | Google 生态完善，工业部署成熟 | ⭐⭐⭐ |

### 模型与推理

| 工具 | 链接 | 简介 | 推荐 |
|:-----|:-----|:-----|:----:|
| **Hugging Face Transformers** | [GitHub](https://github.com/huggingface/transformers) | 预训练模型事实标准库，生态最全 | ⭐⭐⭐⭐⭐ |
| **vLLM** | [GitHub](https://github.com/vllm-project/vllm) | 高性能 LLM 推理框架，PagedAttention 加速 | ⭐⭐⭐⭐⭐ |
| **SGLang** | [GitHub](https://github.com/sgl-project/sglang) | 结构化生成与推理框架，复杂调用场景表现强 | ⭐⭐⭐⭐⭐ |
| **LLaMA-Factory** | [GitHub](https://github.com/hiyouga/LLaMA-Factory) | 一站式模型微调工具，低门槛上手 | ⭐⭐⭐⭐⭐ |
| **DeepSpeed** | [GitHub](https://github.com/microsoft/DeepSpeed) | 微软分布式训练加速，ZeRO 系列优化 | ⭐⭐⭐⭐ |
| **Ollama** | [官网](https://ollama.com/) | 本地大模型一键运行，桌面端体验好 | ⭐⭐⭐⭐⭐ |
| **llama.cpp** | [GitHub](https://github.com/ggml-org/llama.cpp) | C++ 推理引擎，CPU 与边缘设备首选 | ⭐⭐⭐⭐ |

### Agent 与应用框架

| 工具 | 链接 | 简介 | 推荐 |
|:-----|:-----|:-----|:----:|
| **LangChain / LangGraph** | [GitHub](https://github.com/langchain-ai/langchain) | LLM 应用与多步 Agent 开发框架 | ⭐⭐⭐⭐⭐ |
| **MCP** | [官网](https://modelcontextprotocol.io/) | 模型上下文协议，AI 工具调用事实标准 | ⭐⭐⭐⭐⭐ |
| **Claude Agent SDK** | [GitHub](https://github.com/anthropics/claude-code) | Anthropic 官方 Agent SDK，复用 Claude Code 内核 | ⭐⭐⭐⭐ |
| **AutoGen** | [GitHub](https://github.com/microsoft/autogen) | 微软多智能体协作框架 | ⭐⭐⭐⭐ |

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
| **LLM Cookbook** | [GitHub](https://github.com/datawhalechina/llm-cookbook) | DataWhale 大模型教程，中文友好 | ⭐⭐⭐⭐⭐ |
| **Hugging Face 教程** | [官网](https://huggingface.co/learn) | NLP 模型使用与训练 | ⭐⭐⭐⭐ |
| **Stanford CS224n** | [课程](http://web.stanford.edu/class/cs224n/) | 斯坦福 NLP 经典课程 | ⭐⭐⭐⭐⭐ |
| **Stanford CS25** | [课程](https://web.stanford.edu/class/cs25/) | Transformers 统一前沿讲座 | ⭐⭐⭐⭐ |
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

### 2024-2026 最新进展

| 论文 | 链接 | 简介 | 年份 |
|:-----|:-----|:-----|:----:|
| **DeepSeek-V3** | [arxiv](https://arxiv.org/abs/2412.19437) | 超强开源 MoE 模型技术报告 | 2024 |
| **DeepSeek-R1** | [arxiv](https://arxiv.org/abs/2501.12948) | 纯强化学习激励推理，RL 路线代表 | 2025 |
| **Qwen2.5** | [arxiv](https://arxiv.org/abs/2412.15115) | 阿里新一代开源模型 | 2024 |
| **Llama 4** | [模型卡](https://huggingface.co/meta-llama/Llama-4-Scout-17B-16E-Original) | Meta 首个 MoE 开源模型，原生多模态 | 2025 |
| **GPT-OSS** | [模型卡](https://cdn.openai.com/pdf/419b6906-9da6-406c-a19d-1bb078ac7637/oai_gpt-oss_model_card.pdf) | OpenAI 首批开源权重模型 | 2025 |
| **Kimi K2** | [arxiv](https://arxiv.org/abs/2507.20534) | 万亿参数开源 Agent 模型，MuonClip 优化器 | 2025 |

更多论文请查看 [Papers](./paper/README.md) 专辑。

---

## 💡 应用产品

### 通用对话

| 产品 | 链接 | 特色 | 推荐 |
|:-----|:-----|:-----|:----:|
| **ChatGPT** | [官网](https://chat.openai.com/) | OpenAI 通用对话，GPT-5 驱动 | ⭐⭐⭐⭐⭐ |
| **Claude** | [官网](https://claude.ai/) | Anthropic 出品，Opus 4.5 编程与写作强 | ⭐⭐⭐⭐⭐ |
| **Gemini** | [官网](https://gemini.google.com/) | Google 多模态，3 Pro 原生百万上下文 | ⭐⭐⭐⭐⭐ |
| **DeepSeek** | [官网](https://chat.deepseek.com/) | 长文本理解与代码生成，性价比高 | ⭐⭐⭐⭐⭐ |
| **Kimi** | [官网](https://kimi.com/) | 月之暗面，超长上下文与深度搜索 | ⭐⭐⭐⭐ |
| **文心一言** | [官网](https://yiyan.baidu.com/) | 百度多模态大模型 | ⭐⭐⭐⭐ |
| **通义千问** | [官网](https://tongyi.aliyun.com/) | 阿里云 Qwen3-Max 商业版 | ⭐⭐⭐⭐ |
| **智谱清言** | [官网](https://chatglm.cn/) | ChatGLM 商业版 | ⭐⭐⭐⭐ |

### AI 编程

| 产品 | 链接 | 特色 | 推荐 |
|:-----|:-----|:-----|:----:|
| **Claude Code** | [GitHub](https://github.com/anthropics/claude-code) | 终端 Agent，仓库级多文件改动与回滚 | ⭐⭐⭐⭐⭐ |
| **Codex CLI** | [GitHub](https://github.com/openai/codex) | OpenAI 终端 Agent，Apache 2.0 开源 | ⭐⭐⭐⭐⭐ |
| **Cursor** | [官网](https://cursor.sh/) | AI 增强版 VS Code，交互流畅 | ⭐⭐⭐⭐⭐ |
| **Windsurf** | [官网](https://codeium.com/windsurf) | 大型代码库跨文件感知强 | ⭐⭐⭐⭐ |
| **Trae** | [官网](https://www.trae.ai/) | 字节跳动 AI IDE，国内可用 | ⭐⭐⭐⭐ |
| **GitHub Copilot** | [官网](https://github.com/features/copilot) | GitHub 官方 AI 编程，生态广 | ⭐⭐⭐⭐ |
| **通义灵码** | [官网](https://tongyi.aliyun.com/lingma) | 阿里 AI 编程助手，免费额度友好 | ⭐⭐⭐⭐ |

### AI 图像

| 产品 | 链接 | 特色 | 推荐 |
|:-----|:-----|:-----|:----:|
| **Midjourney** | [官网](https://www.midjourney.com/) | 艺术级 AI 绘画 | ⭐⭐⭐⭐⭐ |
| **Stable Diffusion** | [官网](https://stability.ai/) | 开源图像生成，可本地部署 | ⭐⭐⭐⭐⭐ |
| **DALL-E** | [官网](https://openai.com/dall-e-3/) | OpenAI 图像生成 | ⭐⭐⭐⭐ |
| **通义万相** | [官网](https://wanxiang.aliyun.com/) | 阿里 AI 绘画 | ⭐⭐⭐⭐ |

更多产品请查看 [Product](./product/README.md) 专辑。

---

## 🤖 Agent 专辑

| 类别 | 代表项目 | 简介 |
|:-----|:---------|:-----|
| **Agent 框架** | LangGraph · AutoGen · CrewAI | 状态图编排与多智能体协作 |
| **协议标准** | MCP (Model Context Protocol) | AI 工具调用事实标准，已被 OpenAI/Google 采纳 |
| **终端 Agent** | Claude Code · Codex CLI · Cline | 仓库级自主编程 Agent |
| **Agent 平台** | Dify · Coze · 腾讯元器 | 零代码可视化智能体构建 |

更多 Agent 资源请查看 [Agent](./agent/README.md) 专辑。

---

## 📖 博客资源

| 文章 | 链接 | 简介 | 推荐 |
|:-----|:-----|:-----|:----:|
| 普通人如何使用 AI？ | [CSDN](https://blog.csdn.net/z551646/article/details/143848325) | AI 基本概念和日常使用 | ⭐⭐⭐ |
| 高效使用 AI 的 21 种技巧 | [知乎](https://zhuanlan.zhihu.com/p/14780214495) | AI 提示词技巧汇总 | ⭐⭐⭐ |
| DeepSeek 本地部署 | [CSDN](https://blog.csdn.net/zhishi0000/article/details/145450612) | 详细本地部署教程 | ⭐⭐⭐⭐ |
| 零基础 AI 入门指南 | [廖雪峰](https://www.liaoxuefeng.com/article/1543329456062498) | 适合新手的入门指南 | ⭐⭐⭐ |
| DeepSeek 搭建知识库 | [CSDN](https://blog.csdn.net/python12345_/article/details/145450272) | 企业级私有知识库 | ⭐⭐⭐ |

更多博客请查看 [Blogs](./blogs/README.md) 专辑。

---

## 🗂️ 项目仓库

| 项目 | 链接 | 简介 | 推荐 |
|:-----|:-----|:-----|:----:|
| **HuggingFace Transformers** | [GitHub](https://github.com/huggingface/transformers) | 主流模型 PyTorch 实现 | ⭐⭐⭐⭐⭐ |
| **LangChain** | [GitHub](https://github.com/langchain-ai/langchain) | LLM 应用开发框架 | ⭐⭐⭐⭐⭐ |
| **LLaMA-Factory** | [GitHub](https://github.com/hiyouga/LLaMA-Factory) | 一站式微调工具 | ⭐⭐⭐⭐⭐ |
| **vLLM** | [GitHub](https://github.com/vllm-project/vllm) | 高性能推理框架 | ⭐⭐⭐⭐⭐ |
| **SGLang** | [GitHub](https://github.com/sgl-project/sglang) | 结构化生成与推理框架 | ⭐⭐⭐⭐⭐ |
| **MCP** | [GitHub](https://github.com/modelcontextprotocol) | 模型上下文协议，工具调用标准 | ⭐⭐⭐⭐⭐ |
| **Claude Code** | [GitHub](https://github.com/anthropics/claude-code) | 终端 AI 编程 Agent | ⭐⭐⭐⭐⭐ |
| **DeepSpeed** | [GitHub](https://github.com/microsoft/DeepSpeed) | 分布式训练加速 | ⭐⭐⭐⭐ |
| **OpenAI Cookbook** | [GitHub](https://github.com/openai/openai-cookbook) | OpenAI 官方示例 | ⭐⭐⭐⭐ |
| **Prompt Engineering Guide** | [GitHub](https://github.com/dair-ai/Prompt-Engineering-Guide) | 提示工程指南 | ⭐⭐⭐⭐ |

更多仓库请查看 [Repository](./repository/README.md) 专辑。

---

## 🔧 专题深入

| 专题 | 描述 | 路径 |
|:-----|:-----|:-----|
| 🤖 **AI 学习路线** | 系统化的 AI 学习路径规划 | [查看](./docs/ai-learning-roadmap.md) |
| 💻 **AI Coding IDE** | 主流 AI 编程 IDE 与 Agent 对比 | [查看](./docs/ai-coding-ide.md) |
| 🖥️ **GPU 知识** | GPU 原理与环境配置指南 | [查看](./docs/gpu-knowledge.md) |
| 🛠️ **AI 技术栈** | 完整的技术栈文档 | [查看](./docs/ai-tech-stack.md) |
| 🎨 **IDE AI 插件** | 主流 IDE AI 插件收集 | [查看](./docs/ide-ai-plugins.md) |
| 🐚 **OpenClaw** | OpenClaw 使用指南 | [查看](./docs/openclaw.md) |
| 🎯 **AI 工具补充** | UI 生成等工具 | [查看](./docs/ai-tools-supplement.md) |
| 🏗️ **平台对比** | Dify 与 n8n 平台对比 | [查看](./platform/README.md) |

---

## 📎 联系作者

| 联系方式 | 详情 |
|:---------|:-----|
| 📧 E-mail | yanmingxin.boy@gmail.com |
| 💬 WeChat | 扯编程的淡 |

<img src="assets/wx.png" title="" alt="公众号二维码" width="246">

---

> ⭐ 如果这个项目对你有帮助，请 star 支持一下！