# go-proxy-easy

一个**极简**的 Go HTTP 代理服务器，一行命令即可启动！

## 🚀 快速开始

下载编译好的程序，直接运行：

```bash
# 无需认证，启动代理服务器
./go-proxy-easy

# 自定义端口和认证
./go-proxy-easy -listen 0.0.0.0:8080 -user admin -pass 123456
```

**就这么简单！** 🎉

## 功能特性

- ✨ **使用简单** - 一行命令即可启动
- 🔐 **安全可靠** - 支持 Basic 认证
- 🌐 **全面支持** - HTTP 和 HTTPS 代理
- ⚙️ **灵活配置** - 自定义监听地址和端口
- 🪶 **轻量级** - 无外部依赖，单文件运行

## 安装

### 方式一：下载预编译版本（推荐）
前往 [Releases](https://github.com/c4ys/go-proxy-easy/releases) 页面下载适合您系统的版本。

### 方式二：从源码编译
```bash
git clone https://github.com/c4ys/go-proxy-easy.git
cd go-proxy-easy
go build -o go-proxy-easy main.go
```

## 使用方法

### 命令行参数

- `-listen` 监听地址和端口，默认 `0.0.0.0:8080`
- `-user`   代理认证用户名（可选）
- `-pass`   代理认证密码（可选）

### 使用示例

**基本用法（无认证）：**
```bash
./go-proxy-easy
```

**指定端口：**
```bash
./go-proxy-easy -listen 127.0.0.1:8888
```

**启用认证：**
```bash
./go-proxy-easy -listen 127.0.0.1:8888 -user admin -pass secret
```

**开发调试：**
```bash
go run main.go -listen 127.0.0.1:8080 -user test -pass 123456
```

## 客户端配置

### curl 使用代理

无认证：
```bash
curl -x http://127.0.0.1:8080 http://example.com
```

需要认证：
```bash
curl -x http://admin:secret@127.0.0.1:8080 http://example.com
```

### 浏览器配置

在浏览器中设置 HTTP 代理：
- 代理服务器：`127.0.0.1`
- 端口：`8080`
- 如果启用了认证，需要输入用户名和密码

## 多平台构建

使用内置的构建脚本一键构建所有平台版本：

```bash
./build.sh
```

支持的平台：
- Linux (amd64, arm64)
- Windows (amd64)
- macOS (amd64, arm64)

## 许可证

MIT License
