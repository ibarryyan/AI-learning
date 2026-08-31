# 🛠️ 工具专辑 | Tools

> AI 开发与使用相关工具，内容更新至 2026 年。

---

## 🗂️ 目录分类

- [模型部署](#模型部署)
- [数据处理](#数据处理)
- [训练监控](#训练监控)
- [可视化工具](#可视化工具)
- [Prompt 管理](#prompt-管理)
- [API 网关](#api-网关)

---

## 模型部署

| 工具 | 链接 | 简介 |
|:-----|:-----|:-----|
| **Ollama** | [官网](https://ollama.com/) | 本地模型一键管理，最易上手 |
| **LM Studio** | [官网](https://lmstudio.ai/) | 桌面端本地模型管理 |
| **vLLM** | [GitHub](https://github.com/vllm-project/vllm) | 高吞吐推理服务 |
| **SGLang** | [GitHub](https://github.com/sgl-project/sglang) | 结构化生成框架 |
| **llama.cpp** | [GitHub](https://github.com/ggml-org/llama.cpp) | C++ 推理，CPU/GPU 通用 |
| **LocalAI** | [GitHub](https://github.com/mudler/LocalAI) | OpenAI 兼容本地服务 |

---

## 数据处理

| 工具 | 链接 | 简介 |
|:-----|:-----|:-----|
| **Label Studio** | [GitHub](https://github.com/HumanSignalAI/label-studio) | 多模态数据标注 |
| **Docling** | [GitHub](https://github.com/DS4SD/docling) | 文档解析，IBM 出品 |
| **Unstructured** | [GitHub](https://github.com/Unstructured-IO/unstructured) | 非结构化数据解析 |
| **LlamaParse** | [官网](https://llamaindex.ai/llamaparse) | LlamaIndex 文档解析 |
| **Snorkel** | [官网](https://snorkel.ai/) | 程序化数据标注 |

---

## 训练监控

| 工具 | 链接 | 简介 |
|:-----|:-----|:-----|
| **Weights & Biases** | [官网](https://wandb.ai/) | 实验跟踪与可视化 |
| **MLflow** | [官网](https://mlflow.org/) | 开源 ML 生命周期管理 |
| **TensorBoard** | [官网](https://www.tensorflow.org/tensorboard) | TensorFlow 可视化 |
| **ClearML** | [官网](https://clear.ml/) | 开源实验管理 |

---

## 可视化工具

| 工具 | 链接 | 简介 |
|:-----|:-----|:-----|
| **Netron** | [GitHub](https://github.com/lutzroeder/netron) | 模型结构可视化 |
| **BertViz** | [GitHub](https://github.com/jessevig/bertviz) | 注意力可视化 |
| **Transformers Interpret** | [GitHub](https://github.com/cdpierse/transformers-interpret) | 模型解释 |
| **LangSmith** | [官网](https://smith.langchain.com/) | LangChain 追踪调试 |

---

## Prompt 管理

| 工具 | 链接 | 简介 |
|:-----|:-----|:-----|
| **Promptfoo** | [GitHub](https://github.com/promptfoo/promptfoo) | Prompt 测试评估 |
| **Promptflow** | [GitHub](https://github.com/microsoft/promptflow) | 微软 Prompt 工作流 |
| **LangSmith** | [官网](https://smith.langchain.com/) | LangChain Prompt 管理 |
| **Helicone** | [GitHub](https://github.com/Helicone/helicone) | LLM 可观测性 |

---

## API 网关

| 工具 | 链接 | 简介 |
|:-----|:-----|:-----|
| **LiteLLM** | [GitHub](https://github.com/BerriAI/litellm) | 统一 LLM API 网关 |
| **OpenRouter** | [官网](https://openrouter.ai/) | 多模型 API 聚合 |
| **One API** | [GitHub](https://github.com/songquanpeng/one-api) | 开源 API 管理 |
| **Portkey** | [官网](https://portkey.ai/) | AI 网关与可观测 |

---

## 📊 工具分类图

```
AI 工具生态:
│
├── 🚀 部署
│   ├── 本地: Ollama, LM Studio, llama.cpp
│   └── 服务: vLLM, SGLang, LocalAI
│
├── 📊 数据
│   ├── 标注: Label Studio, Snorkel
│   └── 解析: Docling, Unstructured, LlamaParse
│
├── 📈 监控
│   └── W&B, MLflow, TensorBoard, ClearML
│
├── 👁️ 可视化
│   └── Netron, BertViz, LangSmith
│
├── 📝 Prompt
│   └── Promptfoo, Promptflow, LangSmith
│
└── 🌐 API
    └── LiteLLM, OpenRouter, One API, Portkey
```

---

> 💡 **提示**: 本地部署首选 Ollama；服务端高吞吐用 vLLM；文档解析用 Docling；实验跟踪用 W&B；多模型 API 统一用 LiteLLM 或 OpenRouter。