**不同 AI 有不同长处，真正的 AI 生产力来自“模型/工具组合”，而不是押注一个万能模型。**

这里是按照 **2026 年 9 月的现状**整理。由于 AI 产品更新非常快，**“免费”这里指目前存在免费层/免费额度，不代表永久免费或可以无限使用**；具体额度也可能因地区、活动和套餐调整。

---

# 一、通用大模型：负责“思考、写作、分析、规划”

| 工具 / 模型              | 厂商        | 主要优势           | 适合做什么             | 免费情况         | 商用/付费         |
| -------------------- | --------- | -------------- | ----------------- | ------------ | ------------- |
| **ChatGPT / GPT 系列** | OpenAI    | 综合能力、推理、多模态、生态 | 写作、分析、代码、研究、Agent | ✅ 有 Free     | Plus/Pro 等    |
| **Claude**           | Anthropic | 长文本、代码、写作、复杂任务 | 写文章、代码、文档分析、Agent | ✅ 有 Free     | Pro 起         |
| **Gemini**           | Google    | 多模态、搜索生态、长上下文  | 研究、文档、图片、代码       | ✅ 有免费层       | Google AI Pro |
| **Grok**             | xAI       | 实时信息、社交内容、推理   | 热点、搜索、内容分析        | ⚠️ 有限免费/依套餐  | 付费            |
| **Qwen / 通义千问**      | 阿里        | 中文、开源生态、Agent  | 中文写作、代码、Agent     | ✅ 有免费产品/开源模型 | API 按量        |
| **GLM**              | 智谱 / Z.ai | 中文、Agent、代码    | 中文内容、代码、Agent     | ✅ 有免费使用渠道    | API/会员        |
| **Kimi**             | 月之暗面      | 长文本、中文、研究      | 长文档、资料分析、研究       | ✅ 有免费层       | 会员/API        |
| **DeepSeek**         | DeepSeek  | 推理、代码、性价比、开源生态 | 推理、编程、数学          | ✅ 有免费产品      | API           |
| **豆包**               | 字节跳动      | 中文、多模态、生态      | 内容、图片、视频、办公       | ✅ 有免费产品      | 部分能力收费        |
| **讯飞星火**             | 科大讯飞      | 中文、语音、办公       | 中文内容、语音、办公        | ✅ 有免费层       | 会员/API        |

ChatGPT 官方目前提供 Free 层，包含有限的文件分析、图片生成、Deep Research 等能力；Claude 也提供免费层，付费 Pro 则增加使用额度、Claude Code、Research 等能力。([OpenAI][1])

**这一类 AI 的核心价值：**

> **负责“大脑”。**

它们最适合放在整个 AI 工作流的最上游：

```text
需求
 ↓
大模型分析
 ↓
制定方案
 ↓
调用其他专业 AI
```

---

# 二、AI 搜索 / Research：负责“找资料”

这是非常值得单独列出来的一类。

| 工具                        | 核心能力       |    免费 | 最适合      |
| ------------------------- | ---------- | ----: | -------- |
| **Perplexity**            | AI 搜索 + 引用 |     ✅ | 查资料、行业研究 |
| **Gemini Deep Research**  | 深度研究       | ⚠️ 有限 | 长篇研究     |
| **ChatGPT Deep Research** | 深度研究       | ⚠️ 有限 | 市场/行业分析  |
| **Claude Research**       | 深度研究       |    ⚠️ | 长文研究     |
| **Grok Search**           | 实时搜索       |    ⚠️ | 热点/社交媒体  |
| **Kimi**                  | 中文搜索/研究    |     ✅ | 中文资料     |
| **秘塔 AI 搜索**              | 中文搜索       |     ✅ | 中文资料、论文  |
| **360 AI 搜索**             | 中文搜索       |     ✅ | 国内信息检索   |

Perplexity 当前免费层可以进行基础搜索，并提供有限的 Pro Search；Pro 进一步提供 Research、文件分析、图片生成等能力。([Perplexity AI][2])

### 这一类 AI 的长处

不是“写东西”。

而是：

> **帮你找到东西。**

所以一个非常典型的组合就是：

```text
Perplexity
    ↓
收集资料
    ↓
Claude / GPT
    ↓
分析资料
    ↓
AI 写作
```

这就已经比“直接让 ChatGPT 写文章”高级很多。

---

# 三、AI 图片生成：负责“视觉”

这一类特别适合你的文章，因为它可以和《后西游记》直接关联。

| 工具 / 模型            | 厂商                | 长处            |   免费 | 适合          |
| ------------------ | ----------------- | ------------- | ---: | ----------- |
| **Midjourney**     | Midjourney        | 艺术风格、审美       | ❌/⚠️ | 海报、概念图      |
| **Nano Banana 系列** | Google            | 图片编辑、多模态      |   ⚠️ | 图片修改        |
| **Imagen**         | Google            | 写实、视觉质量       |   ⚠️ | 商业视觉        |
| **GPT Image**      | OpenAI            | 综合生成/编辑       |   ⚠️ | 通用图片        |
| **FLUX**           | Black Forest Labs | 写实、开放生态       |   ⚠️ | 商业视觉        |
| **Ideogram**       | Ideogram          | **文字生成能力强**   |    ✅ | 海报、Logo、封面  |
| **Adobe Firefly**  | Adobe             | 商业设计、Adobe 生态 |    ✅ | 商业设计        |
| **Leonardo AI**    | Leonardo          | 游戏/角色/视觉创作    |    ✅ | 人物、资产       |
| **即梦 AI**          | 字节跳动              | 中文、图片、视频      |    ✅ | 国内内容创作      |
| **豆包图片**           | 字节跳动              | 中文理解、易用       |    ✅ | 社媒图片        |
| **通义万相**           | 阿里                | 中文视觉          |    ✅ | 电商/内容       |
| **文心一格**           | 百度                | 中文图片          | ✅/⚠️ | 插画          |
| **LiblibAI**       | 社区平台              | 模型丰富          | ✅/⚠️ | SD/Flux 工作流 |
| **ComfyUI**        | 开源                | 工作流高度可控       |    ✅ | 专业 AI 绘图    |

Ideogram 目前仍有 Free Plan；Adobe Firefly 也提供免费层，并提供每日有限生成额度。([Ideogram Docs][3])

### 这里特别值得你讲一个观点：

> **不是“哪个图片 AI 最好”，而是“我要解决什么视觉问题”。**

例如：

**做电影海报**

→ Midjourney / Ideogram

**做产品广告**

→ Firefly

**做人物一致性**

→ Flux / Leonardo / 专业工作流

**中文海报**

→ Ideogram / 即梦 / 豆包

---

# 四、AI 视频：这应该是你文章的重点

因为你就是从 AI 电视剧切入的。

| 工具 / 模型         | 主要特点        |    免费 | 适合      |
| --------------- | ----------- | ----: | ------- |
| **Seedance**    | 视频生成、动作/镜头  |    ⚠️ | 影视/短视频  |
| **Kling / 可灵**  | 中文生态、视频生成   | ✅ 有额度 | 短片/广告   |
| **Veo**         | Google 视频模型 |    ⚠️ | 高质量视频   |
| **Sora**        | OpenAI      |    ⚠️ | 视频创作    |
| **Runway**      | 专业 AI 视频工作流 |  ✅ 有限 | 影视/广告   |
| **Vidu**        | 视频生成        |     ✅ | 短视频     |
| **PixVerse**    | 视频/特效       |     ✅ | 社媒视频    |
| **Hailuo / 海螺** | 视频生成        |     ✅ | 人物/剧情视频 |
| **即梦视频**        | 国内创作生态      |     ✅ | 短视频     |
| **Pika**        | 视频特效        |     ✅ | 创意短视频   |
| **Luma**        | 视频生成        |    ⚠️ | 影视视觉    |
| **HeyGen**      | 数字人/口播      |     ✅ | 企业视频    |
| **Synthesia**   | 企业数字人       |    ⚠️ | 培训/企业视频 |

Runway 当前有 Free 计划，但属于一次性 125 credits，并不是每月刷新；HeyGen 免费版目前提供每月 3 个视频、单条最长 1 分钟。([Runway][4])

### 这一领域最适合你的观点

**视频 AI 也不是一个模型包打天下。**

可以拆成：

```text
剧本
 ↓
GPT / Claude
 ↓
人物设计
 ↓
Midjourney / 即梦
 ↓
分镜
 ↓
图像模型
 ↓
视频
 ↓
Seedance / Kling / Veo / Runway
 ↓
配音
 ↓
ElevenLabs
 ↓
音乐
 ↓
Suno
 ↓
剪辑
 ↓
剪映
```

这实际上就是：

> **AI 电视剧生产流水线。**

---

# 五、AI 数字人 / AI 口播

| 工具                | 长处        |   免费 |
| ----------------- | --------- | ---: |
| **HeyGen**        | 数字人、口型、翻译 |    ✅ |
| **Synthesia**     | 企业数字人     |   ⚠️ |
| **D-ID**          | 图片数字人     |   ⚠️ |
| **HeyGen Avatar** | 数字人视频     | ✅ 有限 |
| **即梦数字人**         | 国内内容创作    |   ⚠️ |
| **剪映数字人**         | 短视频       |   ⚠️ |
| **腾讯智影**          | 数字人/视频    |   ⚠️ |

如果做：

> 知识类短视频

其实可以：

```text
Claude
↓
写脚本

HeyGen
↓
数字人

ElevenLabs
↓
声音

剪映
↓
字幕+剪辑
```

一个人就能完成过去需要摄像、演员、剪辑等多个角色才能完成的工作。

---

# 六、AI 语音 / 配音

这个领域非常容易被忽略，但对于“AI 电视剧”非常关键。

| 工具                  | 长处       |   免费 |
| ------------------- | -------- | ---: |
| **ElevenLabs**      | 自然度、声音克隆 |    ✅ |
| **OpenAI Audio**    | 语音理解/生成  |   ⚠️ |
| **Google AI Voice** | 多语言      |   ⚠️ |
| **Azure Speech**    | 企业级语音    |   ⚠️ |
| **讯飞语音**            | 中文语音     | ✅/⚠️ |
| **阿里云 CosyVoice**   | 中文/开源    |   ⚠️ |
| **Fish Audio**      | 语音生成/克隆  | ✅/⚠️ |
| **ChatTTS**         | 中文语音     | ✅ 开源 |
| **CosyVoice**       | 中文 TTS   | ✅ 开源 |
| **GPT-SoVITS**      | 声音克隆     | ✅ 开源 |

ElevenLabs 当前提供 Free 层，每月 10,000 credits；付费 Starter 起提供商业授权等能力。([ElevenLabs][5])

这里特别值得注意：

> **免费 ≠ 可以商用。**

例如 Suno 免费计划目前明确没有商业使用权；Pro 才提供新生成歌曲的商业使用权。([Suno][6])

所以你的文章里最好专门加一列：

> **“免费 ≠ 商用”**

这个提醒很重要。

---

# 七、AI 音乐 / 音效

| 工具                      | 长处       |   免费 |
| ----------------------- | -------- | ---: |
| **Suno**                | AI 作曲/歌曲 |    ✅ |
| **Udio**                | AI 音乐    |   ⚠️ |
| **ElevenLabs Music**    | 音乐/音效    |    ✅ |
| **Stable Audio**        | 音效/音乐    |   ⚠️ |
| **Soundraw**            | 商用背景音乐   |   ⚠️ |
| **AIVA**                | 配乐       |   ⚠️ |
| **Mubert**              | BGM      |   ⚠️ |
| **Adobe Firefly Audio** | 音频       | ✅/⚠️ |

Suno 当前 Free 版每天刷新 50 credits，但免费生成内容不享有商业使用权；Pro/更高套餐提供商业权利。([Suno][6])

---

# 八、AI 编程

这个领域其实是**最容易真正赚钱的领域之一**。

| 工具                 | 长处              |   免费 |
| ------------------ | --------------- | ---: |
| **Claude Code**    | Agent Coding    |   ⚠️ |
| **Codex**          | AI 编程 Agent     |   ⚠️ |
| **Cursor**         | AI IDE          |    ✅ |
| **GitHub Copilot** | IDE 编程辅助        |    ✅ |
| **Gemini CLI**     | CLI Agent       | ✅/⚠️ |
| **Windsurf**       | AI IDE          |    ✅ |
| **Cline**          | 开源 Coding Agent |    ✅ |
| **Roo Code**       | Coding Agent    |    ✅ |
| **Aider**          | CLI Coding      | ✅ 开源 |
| **Continue**       | 开源 AI IDE       |    ✅ |
| **OpenCode**       | 开源 Coding Agent |    ✅ |
| **Qwen Code**      | 中文 Coding Agent | ✅/⚠️ |
| **OpenAI Codex**   | Agent 编程        |   ⚠️ |

GitHub Copilot 目前有 Copilot Free，提供有限功能；Cursor 也有 Hobby 免费计划，但 Agent 请求等有额度限制。([GitHub][7])

这里可以直接得出一个非常适合你的观点：

> **程序员真正应该卖的不是“AI 写代码”，而是“用 AI 快速交付软件”。**

---

# 九、AI 自动化 / Agent

这个领域就是从“使用 AI”走向“让 AI 干活”。

| 工具                | 类型               |     免费 | 适合           |
| ----------------- | ---------------- | -----: | ------------ |
| **n8n**           | Workflow         |  ✅ 自托管 | 自动化          |
| **Dify**          | Agent / Workflow |   ✅ 开源 | 企业 AI        |
| **Coze / 扣子**     | Agent            |      ✅ | 国内 Agent     |
| **LangChain**     | 开发框架             |      ✅ | AI 应用        |
| **LangGraph**     | Agent Runtime    |      ✅ | Agent        |
| **Pydantic AI**   | Agent Framework  |      ✅ | Python Agent |
| **CrewAI**        | Multi-Agent      |      ✅ | 多 Agent      |
| **AutoGen**       | Agent Framework  |      ✅ | 多 Agent      |
| **MCP**           | 工具协议             | ✅ 开放协议 | Agent 工具连接   |
| **Claude Skills** | Agent 能力         |     ⚠️ | 专业任务         |
| **OpenClaw**      | Agent            |     ⚠️ | 自动化          |
| **Zapier AI**     | Automation       |     ⚠️ | 企业自动化        |
| **Make**          | Workflow         |     ⚠️ | 自动化          |

这一层是我认为你文章里**最应该重点讲的地方**。

因为：

```text
AI 工具
   ↓
AI 工作流
   ↓
AI Agent
   ↓
AI 产品
```

是一个非常清晰的升级路径。

---

# 十、AI 设计 / PPT / 办公

| 工具                      | 长处          | 免费 |
| ----------------------- | ----------- | -: |
| **Canva AI**            | PPT、海报、设计   |  ✅ |
| **Gamma**               | PPT/网页      |  ✅ |
| **Napkin AI**           | 信息图         |  ✅ |
| **Manus**               | Agent/办公    | ⚠️ |
| **Beautiful.ai**        | PPT         | ⚠️ |
| **Microsoft Copilot**   | Office      | ⚠️ |
| **Google Workspace AI** | Docs/Sheets | ⚠️ |
| **Adobe Express**       | 设计          |  ✅ |
| **Kimi**                | 文档/PPT      |  ✅ |
| **WPS AI**              | Office      | ⚠️ |
| **腾讯文档 AI**             | 办公          | ⚠️ |

这一类工具的赚钱方式非常直接：

> AI + PPT → PPT 外包
> AI + 设计 → 设计服务
> AI + Excel → 数据分析
> AI + 文档 → 企业知识服务

---

# 十一、AI 3D / 游戏 / 数字资产

这个领域你也可以简单带一下。

| 工具                | 主要用途       | 免费 |
| ----------------- | ---------- | -: |
| **Meshy**         | 文字/图片 → 3D |  ✅ |
| **Tripo AI**      | 3D 模型      |  ✅ |
| **Rodin**         | 3D         | ⚠️ |
| **Spline AI**     | 3D Web     | ⚠️ |
| **Scenario**      | 游戏资产       | ⚠️ |
| **Leonardo**      | 游戏视觉资产     |  ✅ |
| **Blockade Labs** | 3D/360°    | ⚠️ |

---

# 十二、AI 开源生态

如果你的文章想体现一点技术深度，这一块值得单独讲。

| 平台 / 工具              | 作用            |           免费 |
| -------------------- | ------------- | -----------: |
| **Hugging Face**     | 模型/数据集/Spaces |            ✅ |
| **Ollama**           | 本地运行 LLM      |            ✅ |
| **LM Studio**        | 本地运行模型        |            ✅ |
| **ComfyUI**          | 图像/视频工作流      |            ✅ |
| **Stable Diffusion** | 图片生成          |       ✅ 开源生态 |
| **Flux**             | 图片模型          | ⚠️ 部分开源/不同许可 |
| **Qwen**             | 开源模型          |            ✅ |
| **Llama**            | 开放模型          |            ✅ |
| **Mistral**          | 开放模型          |            ✅ |
| **DeepSeek**         | 开源模型          |            ✅ |
| **Gemma**            | Google 开放模型   |            ✅ |

这类工具最大的优势不是：

> “免费。”

而是：

> **你可以自己控制模型、数据和工作流。**

这对企业尤其重要。

[1]: https://openai.com/chatgpt/pricing?utm_source=chatgpt.com "ChatGPT Pricing | OpenAI"
[2]: https://www.perplexity.ai/help-center/zh-CN/articles/11187416-na-ge-perplexity-dingyue-fangan-shihezhe-ni?utm_source=chatgpt.com "哪种 Perplexity 订阅方案适合你？ | Perplexity Help Center"
[3]: https://docs.ideogram.ai/plans-and-pricing/available-plans?utm_source=chatgpt.com "Available Plans | Ideogram"
[4]: https://help.runwayml.com/hc/en-us/articles/50404627334547-Free-plan-details?utm_source=chatgpt.com "Free plan details – Runway"
[5]: https://elevenlabs.io/pricing?utm_source=chatgpt.com "ElevenLabs Pricing for Creators & Businesses of All Sizes"
[6]: https://suno.com/pricing?utm_source=chatgpt.com "Suno | Pricing"
[7]: https://github.com/features/copilot/plans?utm_source=chatgpt.com "GitHub Copilot · Plans & pricing · GitHub"
