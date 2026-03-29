# Makefile for deployment

# 目标目录
TARGET_DIR = /var/www/html

# 源目录
SOURCE_DIR = $(PWD)/repository/nginx

# 默认目标
.PHONY: all
all: help

# 显示帮助信息
.PHONY: help
help:
	@echo "可用的目标:"
	@echo "  install    - 复制文件到 $(TARGET_DIR)"
	@echo "  uninstall  - 从 $(TARGET_DIR) 删除文件"
	@echo "  clean      - 清理临时文件"
	@echo "  help       - 显示此帮助信息"

# 安装文件
.PHONY: install
install:
	@echo "正在检查目标目录..."
	@mkdir -p $(TARGET_DIR)
	@echo "正在复制文件..."
	@cp -r $(SOURCE_DIR)/* $(TARGET_DIR)/
	@echo "复制完成! 文件已复制到 $(TARGET_DIR)"

# 卸载文件
.PHONY: uninstall
uninstall:
	@echo "正在从 $(TARGET_DIR) 删除文件..."
	@rm -rf $(TARGET_DIR)/*
	@echo "删除完成!"

# 清理
.PHONY: clean
clean:
	@echo "清理临时文件..."
	@find . -name "*~" -delete
	@echo "清理完成!"
