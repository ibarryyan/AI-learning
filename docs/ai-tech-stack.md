# 🛠️ AI 技术栈 | AI Tech Stack

> 完整的 AI 开发技术栈，从基础框架到生产部署全涵盖。

---

## 📋 目录

1. [基础层](#1-基础层)
2. [模型层](#2-模型层)
3. [框架层](#3-框架层)
4. [应用层](#4-应用层)
5. [部署层](#5-部署层)
6. [工具链](#6-工具链)

---

## 1. 基础层

### 1.1 编程语言

| 语言 | 占比 | 适用场景 | 推荐度 |
|:-----|:----:|:--------|:------:|
| **Python** | 80%+ | AI/ML 主语言 | ⭐⭐⭐⭐⭐ |
| **C/C++** | 10% | 底层优化、推理引擎 | ⭐⭐⭐⭐ |
| **CUDA** | 5% | GPU 编程 | ⭐⭐⭐⭐⭐ |
| **Rust** | 3% | 高性能组件 | ⭐⭐⭐⭐ |
| **JavaScript** | 2% | Web AI 应用 | ⭐⭐⭐ |

### 1.2 核心依赖

```
Python AI 生态:
├── Python 3.10+ (推荐 3.11/3.12)
│
├── 数值计算
│   ├── numpy - 基础数值计算
│   ├── scipy - 科学计算
│   └── pandas - 数据处理
│
├── 深度学习
│   ├── pytorch - 深度学习框架
│   ├── tensorflow - Google ML 框架
│   ├── jax - 高性能 ML 框架
│   └── paddlepaddle - 百度飞桨
│
└── 工具库
    ├── matplotlib - 可视化
    ├── seaborn - 统计可视化
    ├── plotly - 交互式图表
    └── tqdm - 进度条
```

### 1.3 环境管理

| 工具 | 特点 | 推荐度 |
|:-----|:-----|:------:|
| **uv** | 最快的 Python 包管理器 | ⭐⭐⭐⭐⭐ |
| **conda** | 环境管理强大 | ⭐⭐⭐⭐ |
| **poetry** | 依赖管理规范 | ⭐⭐⭐⭐ |
| **pip + venv** | 简单直接 | ⭐⭐⭐ |

---

## 2. 模型层

### 2.1 主流大模型

#### 开源模型

| 模型 | 参数量 | 特点 | GitHub |
|:-----|:------:|:-----|:-------|
| **DeepSeek-V3** | 671B (37B active) | MoE, 最强开源 | [DeepSeek-AI](https://github.com/deepseek-ai/DeepSeek-V3) |
| **DeepSeek-R1** | 671B | 推理模型, RL | [DeepSeek-AI](https://github.com/deepseek-ai/DeepSeek-R1) |
| **Qwen2.5** | 0.5B-72B | 多尺寸, 中文优化 | [QwenLM](https://github.com/QwenLM/Qwen) |
| **LLaMA 3** | 8B-70B | Meta 开源 | [Meta-Llama](https://github.com/meta-llama/llama3) |
| **Mistral** | 7B-123B | 欧洲最强 | [mistralai](https://github.com/mistralai/mistral-src) |
| **ChatGLM** | 6B-130B | 清华开源 | [THUDM](https://github.com/THUDM/ChatGLM-6B) |
| **Yi** | 6B-34B | 零一万物 | [01.AI](https://github.com/01-ai/Yi) |

#### 模型下载

```bash
# Hugging Face CLI
huggingface-cli download meta-llama/Meta-Llama-3-8B-Instruct

# 使用 modelscope (国内快)
modelscope download --model LLM-Research/Meta-Lllama-3-8B-Instruct

# 使用 openbuddy
git clone https://modelscope.cn/models/OpenBuddy/openbuddy-llama3-8b-v21
```

### 2.2 模型架构

```
大模型架构类型:
├── Decoder-Only (GPT 系列)
│   ├── GPT, GPT-2, GPT-3, GPT-4
│   ├── LLaMA, Qwen, Mistral
│   └── 生成式, 指令跟随
│
├── Encoder-Only (BERT 系列)
│   ├── BERT, RoBERTa, ALBERT
│   └── 理解式, 分类/序列标注
│
├── Encoder-Decoder (T5 系列)
│   ├── T5, BART, FLAN-T5
│   └── 序列到序列任务
│
└── Mixture-of-Experts (MoE)
    ├── DeepSeek-MoE, Mixtral
    └── 稀疏激活, 推理高效
```

---

## 3. 框架层

### 3.1 深度学习框架

| 框架 | GitHub | Star | 特点 |
|:-----|:-------|:----:|:-----|
| **PyTorch** | [link](https://github.com/pytorch/pytorch) | 79k | 灵活, 研究首选 |
| **TensorFlow** | [link](https://github.com/tensorflow/tensorflow) | 181k | 生态完善, 生产级 |
| **JAX** | [link](https://github.com/google/jax) | 27k | 高性能, 函数式 |
| **PaddlePaddle** | [link](https://github.com/PaddlePaddle/Paddle) | 21k | 百度开源, 中文 |

### 3.2 大模型开发框架

| 框架 | 功能 | GitHub |
|:-----|:-----|:-------|
| **Hugging Face Transformers** | 模型加载/微调 | [link](https://github.com/huggingface/transformers) |
| **LLaMA-Factory** | 一站式微调 | [link](https://github.com/hiyouga/LLaMA-Factory) |
| **LangChain** | LLM 应用开发 | [link](https://github.com/langchain-ai/langchain) |
| **LlamaIndex** | 知识库构建 | [link](https://github.com/run-llama/llama_index) |
| **AutoGen** | Multi-Agent | [link](https://github.com/microsoft/autogen) |

### 3.3 训练加速

| 工具 | 功能 | GitHub |
|:-----|:-----|:-------|
| **DeepSpeed** | 分布式训练, ZeRO | [link](https://github.com/microsoft/DeepSpeed) |
| **Megatron-LM** | 大规模训练 | [link](https://github.com/NVIDIA/Megatron-LM) |
| **Fairscale** | PyTorch 扩展 | [link](https://github.com/facebookresearch/fairscale) |
| **ColossalAI** | 大模型训练 | [link](https://github.com/hpcaitech/ColossalAI) |

---

## 4. 应用层

### 4.1 LLM 应用模式

```
LLM 应用模式:
├── Basic
│   ├── API 调用 (OpenAI, Anthropic)
│   ├── Prompt Engineering
│   └── 简单问答
│
├── Intermediate
│   ├── RAG (检索增强生成)
│   ├── Agent (自主 AI)
│   └── Fine-tuning (微调)
│
└── Advanced
    ├── Multi-Agent 系统
    ├── Agentic Workflow
    └── 自主规划与执行
```

### 4.2 RAG 技术栈

| 组件 | 工具 | 说明 |
|:-----|:-----|:-----|
| **文档加载** | LangChain, Unstructured | PDF, Word, HTML 解析 |
| **文本分片** | LangChain, tiktoken | 智能分块 |
| **向量存储** | Chroma, Milvus, Weaviate | 向量数据库 |
| **相似度检索** | LangChain, LlamaIndex | 语义搜索 |
| **生成** | OpenAI, Claude, 本地模型 | 答案生成 |

### 4.3 Agent 框架

| 框架 | 特点 | GitHub |
|:-----|:-----|:-------|
| **LangChain Agents** | 生态丰富 | [link](https://github.com/langchain-ai/langchain) |
| **AutoGen** | Multi-Agent, 微软 | [link](https://github.com/microsoft/autogen) |
| **CrewAI** | 多代理编排 | [link](https://github.com/joaomdmoura/crewAI) |
| **AgentVerse** | 多代理协作 | [link](https://github.com/agents-ai/agentverse) |
| **OpenAI Swarm** | 轻量级多代理 | [link](https://github.com/openai/swarm) |

---

## 5. 部署层

### 5.1 推理框架

| 框架 | 特点 | 支持模型 |
|:-----|:-----|:--------|
| **vLLM** | 高吞吐量, PagedAttention | LLaMA, Qwen, ChatGLM |
| **Text Generation Inference** | Hugging Face 官方 | Transformers 全模型 |
| **llama.cpp** | 本地运行, 量化 | GGUF, GGML |
| **Ollama** | 本地大模型运行 | 主流开源模型 |
| **FastChat** | Vicuna 推理服务 | 对话模型 |
| **LMDeploy** | 量化推理, TurboMind | LLaMA, Qwen |

### 5.2 量化技术

| 方法 | 工具 | 精度 | 加速比 |
|:-----|:-----|:----:|:------:|
| **FP16** | 原生支持 | 损失极小 | 1.5x |
| **INT8** | transformers, bitsandbytes | 损失较小 | 2-3x |
| **INT4** | bitsandbytes, GPTQ | 中等损失 | 4-6x |
| **GGUF** | llama.cpp | 可控损失 | 4-8x |

### 5.3 部署方案

| 方案 | 适用场景 | 工具 |
|:-----|:---------|:-----|
| **本地部署** | 个人/小团队 | Ollama, llama.cpp |
| **API 服务** | 企业内部分享 | vLLM, FastAPI |
| **Kubernetes** | 生产环境 | KServe, Triton |
| **Serverless** | 弹性伸缩 | Lambda, Cloud Run |

---

## 6. 工具链

### 6.1 开发工具

| 类别 | 工具 | 说明 |
|:-----|:-----|:-----|
| **IDE** | Cursor, Windsurf, VS Code | AI 增强编程 |
| **Notebook** | Jupyter, VS Code Notebook | 交互式开发 |
| **实验管理** | MLflow, Weights & Biases | 实验追踪 |
| **版本控制** | Git, DVC | 代码+数据版本 |

### 6.2 数据处理

| 任务 | 工具 |
|:-----|:-----|
| **数据清洗** | Pandas, Polars |
| **数据标注** | Label Studio, Doccano |
| **数据增强** | Albumentations, NLPAug |
| **特征工程** | Featuretools, tsfresh |

### 6.3 模型评估

| 工具 | 功能 |
|:-----|:-----|
| **lm-evaluation-harness** | 大模型评测 |
| **open-compass** | CompassHub 评测 |
| **BERT-Score** | 文本生成评估 |
| **ROUGE/L BLEU** | 传统 NLP 评估 |

### 6.4 2024-2025 新兴工具

| 工具 | 类别 | 说明 |
|:-----|:-----|:-----|
| **uv** | 包管理 | 最快的 Python 包管理器 |
| **Triton** | 编译优化 | GPU 内核优化 |
| **Vary** | CV 大模型 | 视觉开源模型 |
| **Qwen-VL** | 多模态 | 阿里视觉模型 |
| **Janus** | 多模态 | 幻方多模态 |

---

## 🔗 技术栈图谱

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                            AI 技术栈全景图                                   │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  ┌─────────────┐    ┌─────────────┐    ┌─────────────┐                   │
│  │   应用层      │    │   框架层      │    │   模型层      │                   │
│  │             │    │             │    │             │                   │
│  │  LangChain  │    │  Transformers│    │ DeepSeek-V3 │                   │
│  │  LlamaIndex │    │ LLaMA-Factory│    │    Qwen     │                   │
│  │   AutoGen   │    │   DeepSpeed  │    │   LLaMA 3   │                   │
│  │   CrewAI    │    │    vLLM      │    │   Mistral   │                   │
│  └──────┬──────┘    └──────┬──────┘    └──────┬──────┘                   │
│         │                  │                  │                            │
│         ▼                  ▼                  ▼                            │
│  ┌─────────────────────────────────────────────────────────────────┐        │
│  │                           基础层                                 │        │
│  │                                                                 │        │
│  │   Python 3.11+  │  PyTorch  │  NumPy  │  CUDA  │  Docker  │        │
│  │                                                                 │        │
│  └─────────────────────────────────────────────────────────────────┘        │
│                                                                             │
│  ┌─────────────┐    ┌─────────────┐    ┌─────────────┐                   │
│  │   部署层      │    │   工具链      │    │   数据层      │                   │
│  │             │    │             │    │             │                   │
│  │   vLLM      │    │   Cursor    │    │  Chroma     │                   │
│  │  llama.cpp  │    │    uv       │    │  Milvus     │                   │
│  │   Ollama    │    │   MLflow    │    │  Weaviate   │                   │
│  └─────────────┘    └─────────────┘    └─────────────┘                   │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## 📚 学习路径建议

### 入门技术栈
```
1. Python 基础
   ↓
2. PyTorch 入门
   ↓
3. Hugging Face Transformers
   ↓
4. LangChain 基础
   ↓
5. RAG 实战
```

### 进阶技术栈
```
1. LLaMA-Factory 微调
   ↓
2. vLLM 部署
   ↓
3. Agent 开发
   ↓
4. DeepSpeed 分布式
   ↓
5. 生产级部署
```

---

## 🔗 相关资源

| 资源 | 链接 |
|:-----|:-----|
| Hugging Face | https://huggingface.co/ |
| PyTorch 文档 | https://pytorch.org/docs/ |
| LangChain 文档 | https://python.langchain.com/ |
| vLLM 文档 | https://docs.vllm.ai/ |
| LLaMA-Factory | https://github.com/hiyouga/LLaMA-Factory |

---

> 💡 **提示**: 技术栈选择应根据实际项目需求来定，入门建议从 Hugging Face + LangChain 组合开始。
