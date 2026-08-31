# 🤖 Agent 专辑 | AI Agents

> AI Agent 框架、协议与开发资源，内容更新至 2026 年。

---

## 🗂️ 目录分类

- [Agent 框架](#agent-框架)
- [协议与标准](#协议与标准)
- [终端 Agent](#终端-agent)
- [Agent 平台](#agent-平台)
- [学习资源](#学习资源)

---

## Agent 框架

| 框架 | 链接 | 简介 | Star |
|:-----|:-----|:-----|:----:|
| **LangGraph** | [GitHub](https://github.com/langchain-ai/langgraph) | 状态图 Agent 编排，当前主流 | 12k+ |
| **LangChain** | [GitHub](https://github.com/langchain-ai/langchain) | LLM 应用开发框架 | 95k+ |
| **AutoGen** | [GitHub](https://github.com/microsoft/autogen) | 微软多智能体框架 | 35k+ |
| **CrewAI** | [GitHub](https://github.com/joaomdmoura/crewai) | 多 Agent 协作框架 | 25k+ |
| **Claude Agent SDK** | [GitHub](https://github.com/anthropics/claude-code-sdk-python) | Claude Code 内核 SDK | 5k+ |
| **LlamaIndex** | [GitHub](https://github.com/run-llama/llama_index) | RAG 与 Agent 框架 | 37k+ |
| **OpenHands** | [GitHub](https://github.com/All-Hands-AI/OpenHands) | 开源软件工程 Agent | 50k+ |

---

## 协议与标准

### MCP (Model Context Protocol)

**官方地址**: https://modelcontextprotocol.io/

MCP 是 Anthropic 于 2024 年底提出的开放协议，2025 年已成为 AI 工具调用的事实标准。它让 LLM 以统一方式连接外部数据源和工具，已被 OpenAI、Google、Microsoft 等主流厂商采纳，MCP 服务器数量超过 1 万个。

| 资源 | 链接 | 简介 |
|:-----|:-----|:-----|
| **MCP 官方组织** | [GitHub](https://github.com/modelcontextprotocol) | 协议官方仓库 |
| **MCP Servers** | [GitHub](https://github.com/modelcontextprotocol/servers) | 官方服务器集合 |
| **Python SDK** | [GitHub](https://github.com/modelcontextprotocol/python-sdk) | Python 开发包 |
| **TypeScript SDK** | [GitHub](https://github.com/modelcontextprotocol/typescript-sdk) | TypeScript 开发包 |

---

## 终端 Agent

| 工具 | 链接 | 简介 | Star |
|:-----|:-----|:-----|:----:|
| **Claude Code** | [GitHub](https://github.com/anthropics/claude-code) | Anthropic 终端 Agent | 12k+ |
| **Codex CLI** | [GitHub](https://github.com/openai/codex) | OpenAI 终端 Agent，Apache 2.0 | 20k+ |
| **Cline** | [GitHub](https://github.com/cline/cline) | 开源 VS Code Agent | 47k+ |
| **aider** | [GitHub](https://github.com/Aider-AI/aider) | 终端 AI 编程助手 | 30k+ |

---

## Agent 平台

| 平台 | 链接 | 简介 |
|:-----|:-----|:-----|
| **Dify** | [官网](https://dify.ai/) | 开源 AI 应用构建平台，可视化工作流 |
| **Coze** | [官网](https://www.coze.com/) | 字节跳动智能体平台 |
| **腾讯元器** | [官网](https://yuanqi.tencent.com/) | 零代码智能体平台 |
| **文心智能体** | [官网](https://agents.baidu.com/) | 百度智能体 |
| **n8n** | [官网](https://n8n.io/) | 开源自动化工作流，可接 LLM |

---

## 学习资源

| 资源 | 链接 | 简介 |
|:-----|:-----|:-----|
| **MCP 官方文档** | [链接](https://modelcontextprotocol.io/) | Model Context Protocol 教程 |
| **LangGraph Tutorial** | [链接](https://langchain-ai.github.io/langgraph/tutorials/) | 状态图 Agent 编排 |
| **Claude Agent SDK 文档** | [文档](https://docs.anthropic.com/en/docs/claude-code/sdk) | Claude Agent 官方文档 |
| **AutoGen Studio** | [链接](https://microsoft.github.io/autogen/) | 微软多智能体教程 |
| **AI Agent 研究综述** | [arxiv](https://arxiv.org/abs/2308.11432) | Lilian Weng Agent 综述 |

---

## 📊 Agent 生态图

```
AI Agent 生态:
│
├── 🏗️ 框架
│   ├── 编排: LangGraph, LangChain
│   ├── 多智能体: AutoGen, CrewAI
│   └── SDK: Claude Agent SDK
│
├── 🔗 协议
│   └── MCP (Model Context Protocol)
│       ├── 官方 SDK: Python, TypeScript
│       └── 服务器: 10000+ MCP Servers
│
├── 💻 终端
│   ├── Claude Code (Anthropic)
│   ├── Codex CLI (OpenAI, Apache 2.0)
│   └── Cline (开源 VS Code Agent)
│
├── 🌐 平台
│   ├── 开源: Dify, n8n
│   └── 商业: Coze, 腾讯元器, 文心智能体
│
└── 📚 学习
    ├── MCP 官方文档
    ├── LangGraph Tutorial
    └── Lilian Weng Agent 综述
```

---

## 🔗 相关链接

- [AI 技术栈](../docs/ai-tech-stack.md)
- [AI Coding IDE](../docs/ai-coding-ide.md)
- [产品专辑](../product/README.md)

---

> 💡 **提示**: 2026 年开发 Agent 首选 LangGraph 做编排，通过 MCP 连接外部工具，终端场景用 Claude Code 或 Codex CLI，零代码场景用 Dify 或 Coze。