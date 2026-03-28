# 🖥️ GPU 知识指南 | GPU Knowledge Guide

> GPU 原理、深度学习环境配置与优化指南。

---

## 📋 目录

1. [GPU 基础概念](#1-gpu-基础概念)
2. [深度学习 GPU 选择](#2-深度学习-gpu-选择)
3. [环境配置](#3-环境配置)
4. [GPU 监控与优化](#4-gpu-监控与优化)
5. [常见问题解决](#5-常见问题解决)

---

## 1. GPU 基础概念

### 1.1 什么是 GPU？

**GPU (Graphics Processing Unit)** 最初用于图形渲染，后来因并行计算能力强大而被广泛应用于深度学习。

| 对比项 | CPU | GPU |
|:------:|:---:|:---:|
| **核心数** | 几到几十核心 | 几千个小型核心 |
| **设计目标** | 通用计算 | 并行计算 |
| **内存** | 高速缓存为主 | 高带宽显存 |
| **擅长** | 逻辑控制、串行任务 | 大规模矩阵运算 |

### 1.2 GPU 架构发展 (NVIDIA)

| 架构 | 代号 | 年份 | 代表产品 |
|:-----|:-----|:----:|:--------|
| Ampere | GA100 | 2020 | A100, RTX 30xx |
| Ada Lovelace | AD102 | 2022 | H100, RTX 40xx |
| Hopper | H100 | 2023 | H100, H200 |
| Blackwell | B100 | 2024 | B100, B200 |

### 1.3 关键指标

| 指标 | 说明 | 对深度学习的影响 |
|:-----|:-----|:----------------|
| **显存 (VRAM)** | GPU 专用内存 | 决定能加载多大的模型 |
| **CUDA 核心数** | 并行计算单元 | 影响训练速度 |
| **Tensor 核心** | 专为矩阵运算设计 | 大幅加速矩阵乘法 |
| **带宽** | 显存传输速度 | 影响数据传输效率 |
| **功耗 (TDP)** | 热设计功耗 | 决定散热需求 |

---

## 2. 深度学习 GPU 选择

### 2.1 消费级 GPU 对比

| GPU | 显存 | CUDA 核心 | Tensor 核心 | 适合场景 |
|:----|:----:|:---------:|:-----------:|:--------|
| RTX 4090 | 24GB | 16384 | 512 | 入门级大模型 |
| RTX 4080 SUPER | 16GB | 10240 | 320 | 中等模型微调 |
| RTX 4070 Ti SUPER | 16GB | 8448 | 264 | 学习实验 |
| RTX 3090 | 24GB | 10496 | 328 | 性价比之选 |
| RTX 3080 | 10GB | 8704 | 272 | 入门学习 |

### 2.2 专业级 GPU

| GPU | 显存 | 特点 | 价格区间 |
|:----|:-----|:-----|:---------|
| **H100** | 80GB | 当前最强训练卡 | ¥150,000+ |
| **A100** | 40/80GB | 性价比高 | ¥60,000+ |
| **A6000** | 48GB | 专业图形工作站 | ¥40,000+ |
| **L40** | 48GB | 数据中心级别 | ¥35,000+ |

### 2.3 选择建议

```
选择决策树:

├── 预算 < ¥5000
│   └── RTX 4070 Ti SUPER (16GB)
│
├── 预算 ¥5000-15000
│   ├── 单一用户 → RTX 4090 (24GB)
│   └── 多用户 → RTX 3090 (24GB) × 2
│
├── 预算 ¥15000-50000
│   ├── 预训练 → A100 (40GB)
│   └── 微调 → H100 (80GB) 或 A100 × 2
│
└── 预算 > ¥50000
    └── H100 × 4+ (多卡并行)
```

---

## 3. 环境配置

### 3.1 基础环境检查

```bash
# 检查 NVIDIA 驱动
nvidia-smi

# 输出示例:
# +-----------------------------------------------------------------------------+
# | NVIDIA-SMI 535.154.05   Driver Version: 535.154.05   CUDA Version: 12.2     |
# |-------------------------------+----------------------+----------------------+
# | GPU  Name        Persistence-M| Bus-Id        Disp.A | Volatile Uncorr. ECC |
# | Fan  Temp  Perf  Pwr:Usage/Cap|         Memory-Usage | GPU-Util  Compute M. |
# |===============================+======================+======================|
# |   0  NVIDIA GeForce ...  Off  | 00000000:01:00.0 Off |                  N/A |
# |  0%   34C    P8    11W / 450W |      1MiB / 24576MiB |      0%      Default |
# +-------------------------------+----------------------+----------------------+
```

### 3.2 CUDA 环境变量

```bash
# ~/.bashrc 或 ~/.zshrc 添加
export CUDA_HOME=/usr/local/cuda
export PATH=$CUDA_HOME/bin:$PATH
export LD_LIBRARY_PATH=$CUDA_HOME/lib64:$LD_LIBRARY_PATH
```

### 3.3 PyTorch GPU 确认

```python
import torch

# 检查 CUDA 是否可用
print(f"CUDA available: {torch.cuda.is_available()}")

# 检查 GPU 数量
print(f"GPU count: {torch.cuda.device_count()}")

# 检查当前 GPU
if torch.cuda.is_available():
    print(f"Current GPU: {torch.cuda.get_device_name(0)}")
    print(f"GPU Memory: {torch.cuda.get_device_properties(0).total_memory / 1024**3:.2f} GB")
```

### 3.4 常用深度学习镜像

```docker
# PyTorch + CUDA 官方镜像
docker pull pytorch/pytorch:2.2.0-cuda12.1-cudnn8-runtime

# Hugging Face 镜像
docker pull huggingface/transformers-pytorch-gpu:latest

# 包含常用包的镜像
docker pull nvidia/cuda:12.1.0-cudnn8-devel-ubuntu22.04
```

---

## 4. GPU 监控与优化

### 4.1 实时监控工具

```bash
# 基础监控
nvidia-smi

# 持续监控 (每秒刷新)
nvidia-smi -l 1

# 详细监控 (包括PCIe带宽)
nvidia-smi dmon

# 进程详情
nvidia-smi pmon

# 查询GPU详情
nvidia-smi -q
```

### 4.2 Python 监控

```python
import pynvml
import time

pynvml.nvmlInit()

# 获取 GPU 数量
device_count = pynvml.nvmlDeviceGetCount()

for i in range(device_count):
    handle = pynvml.nvmlDeviceGetHandleByIndex(i)
    name = pynvml.nvmlDeviceGetName(handle)
    
    while True:
        mem_info = pynvml.nvmlDeviceGetMemoryInfo(handle)
        util = pynvml.nvmlDeviceGetUtilizationRates(handle)
        
        print(f"GPU: {name}")
        print(f"Memory: {mem_info.used / 1024**3:.2f}GB / {mem_info.total / 1024**3:.2f}GB")
        print(f"Utilization: {util.gpu}%")
        
        time.sleep(1)
```

### 4.3 常用优化技巧

#### 梯度累积 (Gradient Accumulation)

```python
# 显存不足时使用
accumulation_steps = 4
optimizer.zero_grad()

for step, (inputs, labels) in enumerate(dataloader):
    outputs = model(inputs)
    loss = criterion(outputs, labels)
    loss.backward()
    
    if (step + 1) % accumulation_steps == 0:
        optimizer.step()
        optimizer.zero_grad()
```

#### 混合精度训练 (Mixed Precision)

```python
from torch.cuda.amp import autocast, GradScaler

scaler = GradScaler()

for inputs, labels in dataloader:
    optimizer.zero_grad()
    
    with autocast(dtype=torch.float16):
        outputs = model(inputs)
        loss = criterion(outputs, labels)
    
    scaler.scale(loss).backward()
    scaler.step(optimizer)
    scaler.update()
```

#### 梯度检查点 (Gradient Checkpointing)

```python
# 节省显存
from torch.utils.checkpoint import checkpoint_sequential

# 分段保存梯度
model = checkpoint_sequential(module, segments, input_var)
```

### 4.4 模型量化

| 量化方法 | 精度损失 | 显存减少 | 速度提升 |
|:--------|:-------:|:-------:|:-------:|
| FP16 | 极小 | ~50% | 1.5-2x |
| INT8 | 较小 | ~75% | 2-4x |
| INT4 | 中等 | ~87% | 4-8x |

```python
# 使用 bitsandbytes 量化
from transformers import AutoModelForCausalLM, BitsAndBytesConfig

quantization_config = BitsAndBytesConfig(
    load_in_4bit=True,
    bnb_4bit_compute_dtype=torch.float16
)

model = AutoModelForCausalLM.from_pretrained(
    "model_name",
    quantization_config=quantization_config
)
```

---

## 5. 常见问题解决

### 5.1 显存不足 (OOM)

**错误**: `RuntimeError: CUDA out of memory`

**解决方案**:

1. **减小 batch size**
```python
# 原来
batch_size = 32
# 修改为
batch_size = 8
```

2. **启用梯度累积**
```python
accumulation_steps = 4
```

3. **使用混合精度**
```python
with autocast(dtype=torch.float16):
    outputs = model(inputs)
```

4. **使用 LoRA 微调**
```python
from peft import LoraConfig, get_peft_model

lora_config = LoraConfig(r=8, lora_alpha=16, target_modules=["q_proj", "v_proj"])
model = get_peft_model(model, lora_config)
```

5. **使用量化模型**
```python
# GGUF 格式量化模型
model = AutoModelForCausalLM.from_pretrained(
    "model.gguf",
    torch_dtype=torch.float16,
    device_map="auto"
)
```

### 5.2 CUDA 版本不匹配

```bash
# 检查各组件版本
nvcc --version          # CUDA 编译器版本
nvidia-smi              # CUDA 运行时版本
python -c "import torch; print(torch.version.cuda)"  # PyTorch CUDA 版本

# 推荐版本匹配
# PyTorch 2.x + CUDA 12.1
# TensorFlow 2.x + CUDA 11.x / 12.x
```

### 5.3 多卡并行训练

```python
# DataParallel (单节点多卡)
model = torch.nn.DataParallel(model)

# DistributedDataParallel (多节点多卡)
import torch.distributed as dist

dist.init_process_group(backend="nccl")
model = torch.nn.parallel.DistributedDataParallel(model)
```

### 5.4 常用训练脚本

#### DeepSpeed 配置示例

```json
{
  "train_batch_size": "auto",
  "train_micro_batch_size_per_gpu": "auto",
  "gradient_accumulation_steps": "auto",
  "gradient_clipping": 1.0,
  "zero_optimization": {
    "stage": 3,
    "offload_optimizer": {
      "device": "cpu",
      "pin_memory": true
    },
    "offload_param": {
      "device": "cpu",
      "pin_memory": true
    }
  },
  "fp16": {
    "enabled": true
  }
}
```

---

## 📚 相关资源

| 资源 | 链接 |
|:-----|:-----|
| NVIDIA 驱动下载 | https://www.nvidia.com/Download/index.aspx |
| CUDA Toolkit | https://developer.nvidia.com/cuda-toolkit |
| PyTorch GPU 指南 | https://pytorch.org/docs/stable/notes/cuda.html |
| DeepSpeed 文档 | https://www.deepspeed.ai/getting-started/ |
| vLLM 文档 | https://docs.vllm.ai/ |

---

## 🧮 显存需求估算

| 模型 | 7B | 13B | 34B | 70B |
|:----|:--:|:---:|:---:|:---:|
| **FP16 推理** | 14GB | 26GB | 68GB | 140GB |
| **FP16 训练** | 56GB | 104GB | 272GB | 560GB |
| **LoRA (r=8)** | 16GB | 28GB | 72GB | 144GB |
| **QLoRA (INT4)** | 5GB | 10GB | 20GB | 40GB |

> 💡 以上为估算值，实际需求因模型架构和配置不同而有差异。

---

> 💡 **提示**: 对于学习者，RTX 4090 (24GB) 是性价比最高的选择，配合量化技术可以运行 70B 参数模型。
