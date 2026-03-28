# 📚 AI 学习路线 | AI Learning Roadmap

> 系统化的 AI 学习路径规划，从入门到进阶全覆盖。

---

## 🗺️ 学习路径总览

```
┌─────────────────────────────────────────────────────────────────────────┐
│                           AI 学习路线图                                  │
├─────────────────────────────────────────────────────────────────────────┤
│                                                                         │
│   ┌──────────┐     ┌──────────┐     ┌──────────┐     ┌──────────┐      │
│   │  入门阶段  │ ──▶ │  基础阶段  │ ──▶ │  进阶阶段  │ ──▶ │  实战阶段  │      │
│   │ (1-2月)   │     │ (2-3月)   │     │ (3-6月)   │     │ (持续)    │      │
│   └──────────┘     └──────────┘     └──────────┘     └──────────┘      │
│       │                │                │                │            │
│       ▼                ▼                ▼                ▼            │
│   Python基础      机器学习基础      深度学习理论       项目实战         │
│   数学基础        深度学习入门      大模型原理         Kaggle竞赛       │
│   AI入门          PyTorch入门       Transformer        开源贡献        │
│                                                         部署运维         │
└─────────────────────────────────────────────────────────────────────────┘
```

---

## 🚀 阶段一：入门阶段 (1-2个月)

### 目标
- 掌握 Python 编程基础
- 了解 AI 基本概念
- 具备初步的代码能力

### 学习内容

#### 1.1 Python 基础

| 知识点 | 推荐资源 | 掌握程度 |
|:-------|:---------|:---------|
| 基础语法 | [廖雪峰 Python](https://www.liaoxuefeng.com/wiki/1016959663602400) | ⭐⭐⭐⭐⭐ |
| 数据结构 | Python 官方文档 | ⭐⭐⭐⭐ |
| 函数与面向对象 | Python 教程 | ⭐⭐⭐⭐ |
| 文件操作与异常 | 实战练习 | ⭐⭐⭐ |

#### 1.2 数学基础

| 知识点 | 推荐资源 | 掌握程度 |
|:-------|:---------|:---------|
| 线性代数 | 3Blue1Brown 视频 | ⭐⭐⭐⭐⭐ |
| 概率论 | Khan Academy | ⭐⭐⭐⭐ |
| 微积分 | StatQuest 视频 | ⭐⭐⭐ |
| 优化基础 | 机器学习数学基础 | ⭐⭐⭐⭐ |

#### 1.3 AI 入门

| 知识点 | 推荐资源 | 掌握程度 |
|:-------|:---------|:---------|
| AI 是什么 | [微软 AI for Beginners](https://github.com/microsoft/AI-For-Beginners) | ⭐⭐⭐⭐⭐ |
| 机器学习基础 | 吴恩达 ML 课程 | ⭐⭐⭐⭐⭐ |
| 环境配置 | Anaconda + Jupyter | ⭐⭐⭐⭐⭐ |

---

## 📖 阶段二：基础阶段 (2-3个月)

### 目标
- 掌握机器学习核心算法
- 学会使用深度学习框架
- 能够独立完成简单项目

### 学习内容

#### 2.1 机器学习

| 算法 | 类别 | 推荐资源 |
|:-----|:-----|:---------|
| 线性回归 | 监督学习 | 吴恩达课程 |
| 逻辑回归 | 分类 | sklearn 文档 |
| 决策树 | 树模型 | Kaggle 教程 |
| SVM | 核方法 | 李航《统计学习方法》 |
| KNN | 距离模型 | sklearn 示例 |
| 聚类(K-Means) | 无监督 | 实战练习 |

#### 2.2 深度学习入门

| 框架 | 特点 | 推荐资源 |
|:-----|:-----|:---------|
| **PyTorch** | 灵活、研究首选 | [官方教程](https://pytorch.org/tutorials/) |
| **TensorFlow** | 生态完善、生产可用 | Keras 官方文档 |

#### 2.3 核心概念

```
深度学习核心概念：
├── 神经网络基础
│   ├── 前向传播
│   ├── 激活函数 (ReLU, Sigmoid, Tanh)
│   └── 反向传播与梯度下降
├── 优化方法
│   ├── SGD / Adam / AdamW
│   ├── 学习率调度
│   └── 正则化 (Dropout, L2)
├── 卷积神经网络 (CNN)
│   ├── 卷积层
│   ├── 池化层
│   └── 经典架构 (LeNet, AlexNet, VGG)
└── 循环神经网络 (RNN/LSTM/GRN)
    ├── 序列建模
    ├── 注意力机制
    └── NLP 基础
```

---

## 🔬 阶段三：进阶阶段 (3-6个月)

### 目标
- 深入理解大模型原理
- 掌握 Transformer 架构
- 能够进行模型微调和部署

### 学习内容

#### 3.1 Transformer 深入

| 论文 | 年份 | 重要性 |
|:-----|:----:|:------:|
| Attention Is All You Need | 2017 | ⭐⭐⭐⭐⭐ |
| BERT: Pre-training of Deep Bidirectional Transformers | 2018 | ⭐⭐⭐⭐⭐ |
| GPT: Improving Language Understanding by Generative Pre-Training | 2018 | ⭐⭐⭐⭐⭐ |
| GPT-2: Language Models are Unsupervised Multitask Learners | 2019 | ⭐⭐⭐⭐ |
| GPT-3: Language Models are Few-Shot Learners | 2020 | ⭐⭐⭐⭐⭐ |
| Transformer-XL: Attentive Language Models Beyond a Fixed-Length Context | 2019 | ⭐⭐⭐⭐ |

#### 3.2 大模型专项

| 主题 | 知识点 | 推荐资源 |
|:-----|:-------|:---------|
| **模型架构** | Transformer, MoE, RWKV | [LLaMA 论文](https://arxiv.org/abs/2302.13971) |
| **预训练** | 文本语料、数据处理 | [DeepSpeed](https://github.com/microsoft/DeepSpeed) |
| **微调技术** | LoRA, QLoRA, Prefix Tuning | [LLaMA-Factory](https://github.com/hiyouga/LLaMA-Factory) |
| **提示工程** | Few-shot, CoT, ToT | [Prompt Engineering Guide](https://github.com/dair-ai/Prompt-Engineering-Guide) |
| **推理优化** | vLLM, GGUF, 量化 | [vLLM](https://github.com/vllm-project/vllm) |

#### 3.3 2024-2025 最新技术

| 技术 | 描述 | 论文/项目 |
|:-----|:-----|:---------|
| **DeepSeek-R1** | 纯强化学习推理模型 | [arxiv](https://arxiv.org/abs/2501.12948) |
| **DeepSeek-V3** | 超强 MoE 开源模型 | [arxiv](https://arxiv.org/abs/2412.19437) |
| **Qwen2.5** | 阿里新一代开源模型 | [GitHub](https://github.com/QwenLM/Qwen) |
| **GPT-4o** | OpenAI 旗舰多模态 | OpenAI 官网 |
| **Claude 3.5** | Anthropic 最新模型 | Anthropic 官网 |

---

## 🛠️ 阶段四：实战阶段 (持续)

### 目标
- 完成实际项目
- 参与开源贡献
- 具备生产环境部署能力

### 项目方向

#### 4.1 入门项目
| 项目 | 描述 | 技术栈 |
|:-----|:-----|:-------|
| 垃圾邮件分类 | 二分类任务 | PyTorch, sklearn |
| 手写数字识别 | CNN 入门 | PyTorch, MNIST |
| 情感分析 | NLP 基础 | Transformers, BERT |
| AI 对话机器人 | 简单对话 | LangChain, OpenAI API |

#### 4.2 进阶项目
| 项目 | 描述 | 技术栈 |
|:-----|:-----|:-------|
| 本地知识库 | RAG 系统 | LangChain, Chroma, vLLM |
| AI 写作助手 | 文案生成 | LLM Fine-tuning |
| 多模态应用 | 图生文/文生图 | CLIP, Stable Diffusion |
| AI 编程助手 | 代码补全 | CodeGen, Copilot API |

#### 4.3 高级项目
| 项目 | 描述 | 技术栈 |
|:-----|:-----|:-------|
| 大模型预训练 | 从零训练 | DeepSpeed, Megatron |
| Agent 系统 | 自主 AI | LangChain, AutoGen |
| 生产级部署 | 高并发服务 | vLLM, Kubernetes |
| 模型量化压缩 | 边缘部署 | llama.cpp, AWQ, GPTQ |

---

## 📚 推荐学习资源

### 在线课程

| 课程 | 链接 | 评价 |
|:-----|:-----|:-----|
| 机器学习 - 吴恩达 | [Coursera](https://www.coursera.org/learn/machine-learning) | ⭐⭐⭐⭐⭐ |
| 深度学习 - 吴恩达 | [DeepLearning.AI](https://www.deeplearning.ai/) | ⭐⭐⭐⭐⭐ |
| CS224n NLP | [Stanford](http://web.stanford.edu/class/cs224n/) | ⭐⭐⭐⭐⭐ |
| CS231n CV | [Stanford](http://cs231n.stanford.edu/) | ⭐⭐⭐⭐⭐ |
| 微软 AI for Beginners | [GitHub](https://github.com/microsoft/AI-For-Beginners) | ⭐⭐⭐⭐⭐ |

### 书籍

| 书名 | 作者 | 评价 |
|:-----|:-----|:-----|
| 《深度学习》 | Ian Goodfellow | ⭐⭐⭐⭐⭐ |
| 《机器学习》 | 周志华 | ⭐⭐⭐⭐ |
| 《统计学习方法》 | 李航 | ⭐⭐⭐⭐ |
| 《动手学深度学习》 | 李沐 | ⭐⭐⭐⭐⭐ |
| 《大语言模型》 | 赵鑫 | ⭐⭐⭐⭐⭐ |

### 社区

| 社区 | 链接 | 描述 |
|:-----|:-----|:-----|
| Hugging Face | [huggingface.co](https://huggingface.co/) | ML 社区 |
| Papers with Code | [paperswithcode.com](https://paperswithcode.com/) | 论文+代码 |
| Kaggle | [kaggle.com](https://www.kaggle.com/) | 竞赛平台 |
| ArXiv | [arxiv.org](https://arxiv.org/) | 论文预印本 |

---

## 📅 学习计划建议

### 零基础入门计划 (12周)

| 周次 | 主题 | 目标 |
|:-----|:-----|:-----|
| 1-2 | Python 基础 | 能够编写简单程序 |
| 3-4 | 数学基础 | 理解线性代数概率论基本概念 |
| 5-6 | 机器学习入门 | 掌握 sklearn 使用 |
| 7-8 | 深度学习基础 | 掌握 PyTorch 基础 |
| 9-10 | CNN/RNN | 完成图像/文本分类项目 |
| 11-12 | Transformer 入门 | 理解注意力机制 |

### 进阶提升计划 (24周)

| 周次 | 主题 | 目标 |
|:-----|:-----|:-----|
| 1-4 | Transformer 深入 | 精读经典论文 |
| 5-8 | 大模型微调 | 掌握 LoRA/QLoRA |
| 9-12 | RAG 系统 | 搭建知识库 |
| 13-16 | Agent 开发 | 构建 AI Agent |
| 17-20 | 部署优化 | 掌握 vLLM 部署 |
| 21-24 | 项目实战 | 完成完整项目 |

---

> 💡 **提示**: 学习是一个持续的过程，建议边学边做，通过项目来巩固知识。同时关注最新技术动态，保持学习的主动性。
