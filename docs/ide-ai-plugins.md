# 🎨 IDE AI 插件 | IDE AI Plugins

> 主流 IDE 的 AI 插件收集，助你提升编程效率。

---

## 📋 目录

1. [VS Code 插件](#1-vs-code-插件)
2. [JetBrains 插件](#2-jetbrains-插件)
3. [其他 IDE 插件](#3-其他-ide-插件)
4. [命令行 AI 工具](#4-命令行-ai-工具)

---

## 1. VS Code 插件

### 1.1 AI 编程插件

| 插件 | 安装量 | 功能 | 推荐度 |
|:-----|:------:|:-----|:------:|
| **Cline** | 5M+ | 开源 VS Code Agent，多模型 | ⭐⭐⭐⭐⭐ |
| **Continue** | 800k+ | 开源 AI 编程助手 | ⭐⭐⭐⭐⭐ |
| **GitHub Copilot** | 5M+ | GitHub 官方 AI，GPT-5 | ⭐⭐⭐⭐ |
| **Cursor** | 500k+ | AI 增强版 VS Code | ⭐⭐⭐⭐⭐ |
| **CodeGPT** | 600k+ | 多模型集成 | ⭐⭐⭐⭐ |
| **Tabnine** | 20M+ | AI 代码补全 | ⭐⭐⭐⭐ |
| **Amazon Q Developer** | 1M+ | AWS 开发助手 | ⭐⭐⭐⭐ |

### 1.2 Cline 插件详解

**特点**: 开源 VS Code Agent，Apache 2.0，安装量超 500 万

```json
// settings.json 配置示例
{
  "cline.apiProvider": "anthropic",
  "cline.model": "claude-opus-4-5",
  "cline.maxTokens": 8192,
  "cline.autoApprove": {
    "enabled": true,
    "patterns": ["*.py", "*.md"]
  }
}
```

**核心功能**:
- 🤖 自主编码：文件编辑、终端操作、浏览器控制
- 📝 结构化计划：任务拆解与逐步执行
- 🔗 MCP 集成：调用外部工具与服务
- 🎯 多模型：Claude 4、GPT-5、DeepSeek、本地模型

### 1.3 Continue 插件详解

**特点**: 开源免费，支持多模型切换

```json
// settings.json 配置示例
{
  "continue.enableQuickActions": true,
  "continue.useEmbeddedCompletions": true,
  "continue.models": [
    {
      "model": "claude-opus-4-5",
      "provider": "anthropic",
      "apiKey": "${env:ANTHROPIC_API_KEY}"
    },
    {
      "model": "gpt-5",
      "provider": "openai",
      "apiKey": "${env:OPENAI_API_KEY}"
    }
  ]
}
```

**快捷键**:
- `Ctrl+L` / `Cmd+L`: 打开 AI 对话
- `Ctrl+I` / `Cmd+I`: 智能编辑
- `Tab`: 接受补全

### 1.4 CodeGPT 插件

**特点**: 界面友好，支持多种模型

```json
// settings.json 配置
{
  "codegpt.selectedProvider": "OpenAI",
  "codegpt.model": "gpt-5",
  "codegpt.apiKey": "${env:OPENAI_API_KEY}"
}
```

### 1.5 其他实用插件

| 插件 | 功能 | 推荐度 |
|:-----|:-----|:------:|
| **Error Lens** | 错误高亮显示 | ⭐⭐⭐⭐ |
| **GitLens** | Git 增强 | ⭐⭐⭐⭐⭐ |
| **Prettier** | 代码格式化 | ⭐⭐⭐⭐⭐ |
| **ESLint** | JS/TS 代码检查 | ⭐⭐⭐⭐⭐ |
| **Python** | Python 支持 | ⭐⭐⭐⭐⭐ |
| **Pylance** | Python 语言服务器 | ⭐⭐⭐⭐⭐ |

---

## 2. JetBrains 插件

### 2.1 JetBrains AI

| 特性 | 说明 |
|:-----|:-----|
| **官方插件** | JetBrains 官方 AI 助手 |
| **支持 IDE** | IntelliJ, PyCharm, WebStorm, etc. |
| **价格** | $30/月 |
| **免费额度** | 5000 次/月 |

**功能**:
- 🤖 代码补全与生成
- 💬 自然语言编程
- 📝 代码解释与重构
- 🔍 Bug 分析与修复

### 2.2 其他 JetBrains 插件

| 插件 | 功能 | 推荐度 |
|:-----|:-----|:------:|
| **AiX** | AI 代码补全 | ⭐⭐⭐⭐ |
| **Bito** | AI 代码助手 | ⭐⭐⭐⭐ |
| **Movier** | AI 编程助手 | ⭐⭐⭐ |
| **aiXcoder** | 智能代码补全 | ⭐⭐⭐⭐ |

---

## 3. 其他 IDE 插件

### 3.1 Vim/Neovim

| 插件 | 功能 | 语言 |
|:-----|:-----|:-----|
| **Copilot.vim** | GitHub Copilot | VimScript |
| **vim-ai** | ChatGPT 集成 | VimScript |
| **ChatGPT.nvim** | Neovim ChatGPT | Lua |
| **CodeGPT.nvim** | Neovim CodeGPT | Lua |

**配置示例** (nvim):
```lua
-- 使用 packer.nvim
use({
  "CopilotC-Nvim/CopilotChat.nvim",
  requires = {
    { "zbirenbaum/copilot.lua" },
  },
  config = function()
    require("CopilotChat").setup()
  end
})
```

### 3.2 Sublime Text

| 插件 | 功能 |
|:-----|:-----|
| **Copilot** | GitHub Copilot |
| **ChatGPT** | OpenAI 集成 |

### 3.3 Emacs

| 插件 | 功能 |
|:-----|:-----|
| **copilot.el** | GitHub Copilot |
| **gptel** | 多模型支持 |

---

## 4. 命令行 AI 工具

### 4.1 AI CLI 工具

| 工具 | 功能 | GitHub |
|:-----|:-----|:-------|
| **Claude Code** | 终端 Agent，仓库级多文件改动 | [link](https://github.com/anthropics/claude-code) |
| **Codex CLI** | OpenAI 终端 Agent，Apache 2.0 | [link](https://github.com/openai/codex) |
| **aider** | 终端 AI 编程助手 | [link](https://github.com/Aider-AI/aider) |
| **aichat** | 多模型 CLI 工具 | [link](https://github.com/sigoden/aichat) |
| **ShellGPT** | Linux Shell AI | [link](https://github.com/TheR1D/shell_gpt) |
| **llama-cli** | 本地模型 CLI | [link](https://github.com/ggml-org/llama.cpp) |

### 4.2 aichat 使用

```bash
# 安装
cargo install aichat

# 使用
aichat -p "解释什么是 Transformer"

# 交互模式
aichat
```

### 4.3 终端 AI 辅助

```bash
# 使用 shell-gpt
sgpt "列出当前目录下的所有文件"

# 使用 AI 命令补全
eval "$(shellgpt --init)"
```

---

## 📊 插件对比

| 插件 | 免费版 | 付费版 | 支持模型 | 离线支持 |
|:-----|:------|:------|:--------|:-------:|
| **Cline** | 按 API 计费 | 无平台费 | Claude/GPT-5/DeepSeek/本地 | ✅ |
| **Continue** | ✅ 无限 | - | 多模型 | ✅ |
| **Cursor** | 500次/槽 | $20/月 | GPT-5/Claude 4.5 | ❌ |
| **Copilot** | 2000次/月 | $10/月 | GPT-5 | ❌ |
| **Tabnine** | 有限 | $12/月 | 自研 | ✅ |
| **JetBrains AI** | 5000次/月 | $30/月 | 自研 | ❌ |

---

## 🎯 选择建议

| 场景 | 推荐插件 |
|:-----|:---------|
| **VS Code Agent** | Cline |
| **免费首选** | Continue |
| **专业开发** | Cursor / Windsurf |
| **企业用户** | Copilot / JetBrains AI |
| **本地部署** | Cline/Continue + Ollama |
| **终端 Agent** | Claude Code / Codex CLI |

---

## 🔧 进阶配置

### Continue 完整配置

```json
{
  "continue.enableQuickActions": true,
  "continue.useEmbeddedCompletions": true,
  "continue.maxTokens": 2048,
  "continue.temperature": 0.7,
  "continue.models": [
    {
      "model": "claude-opus-4-5",
      "provider": "anthropic",
      "apiKey": "${env:ANTHROPIC_API_KEY}"
    },
    {
      "model": "gpt-5",
      "provider": "openai",
      "apiKey": "${env:OPENAI_API_KEY}"
    },
    {
      "model": "deepseek-v3",
      "provider": "deepseek",
      "apiKey": "${env:DEEPSEEK_API_KEY}"
    },
    {
      "model": "llama4",
      "provider": "ollama"
    }
  ],
  "continue.contextProviders": [
    "github",
    "git",
    "search",
    "terminal"
  ]
}
```

### Tabnine 配置

```json
{
  "tabnine.experimentalAutoImports": true,
  "tabnine.maxNumResults": 5,
  "tabnine.automation": {
    "enable": true
  }
}
```

---

## 📚 相关资源

| 资源 | 链接 |
|:-----|:-----|
| Continue 官网 | https://continue.dev/ |
| VS Code Marketplace | https://marketplace.visualstudio.com/ |
| JetBrains Marketplace | https://plugins.jetbrains.com/ |

---

> 💡 **建议**: 2026 年 VS Code 用户首选 Cline，开源 Agent 能力强且支持多模型；如需更专业的体验，可选择 Cursor 或 Windsurf；终端重度用户推荐 Claude Code 或 Codex CLI。