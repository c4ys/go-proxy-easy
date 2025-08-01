#!/bin/bash

# 构建脚本 for go-proxy-easy

set -e

echo "正在构建 go-proxy-easy..."

# 检查 Go 是否安装
if ! command -v go &> /dev/null; then
    echo "错误: 请先安装 Go"
    exit 1
fi

# 获取版本信息
VERSION=$(git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
GIT_COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")

# 构建信息
LDFLAGS="-X main.Version=${VERSION} -X main.BuildTime=${BUILD_TIME} -X main.GitCommit=${GIT_COMMIT}"

echo "版本: ${VERSION}"
echo "构建时间: ${BUILD_TIME}"
echo "Git 提交: ${GIT_COMMIT}"

# 构建不同平台的二进制文件
platforms=("linux/amd64" "linux/arm64" "windows/amd64" "darwin/amd64" "darwin/arm64")

mkdir -p dist

for platform in "${platforms[@]}"; do
    IFS='/' read -r os arch <<< "$platform"
    output_name="go-proxy-easy"
    
    if [ "$os" = "windows" ]; then
        output_name="${output_name}.exe"
    fi
    
    output_path="dist/${output_name}-${os}-${arch}"
    if [ "$os" = "windows" ]; then
        output_path="${output_path}.exe"
    fi
    
    echo "构建 ${os}/${arch}..."
    GOOS=$os GOARCH=$arch go build -ldflags "$LDFLAGS" -o "$output_path" main.go
done

echo "构建完成！输出目录: dist/"
ls -la dist/
