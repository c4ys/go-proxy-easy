# go-proxy-easy

一个简单的 Go HTTP 代理服务器，支持通过命令行设置监听地址、端口、用户名和密码。

## 功能特性

- 支持 HTTP 和 HTTPS 代理
- 支持 Basic 认证
- 支持自定义监听地址和端口
- 轻量级，无外部依赖

## 安装

```bash
git clone https://github.com/c4ys/go-proxy-easy.git
cd go-proxy-easy
go mod tidy
```

## 使用方法

### 基本用法

```bash
go run main.go
```

### 指定监听地址和端口

```bash
go run main.go -listen 127.0.0.1:8888
```

### 启用认证

```bash
go run main.go -listen 127.0.0.1:8888 -user admin -pass secret
```

### 命令行参数

- `-listen` 监听地址和端口，默认 `0.0.0.0:8080`
- `-user`   代理认证用户名（可选）
- `-pass`   代理认证密码（可选）

## 编译

编译生成可执行文件：

```bash
go build -o go-proxy-easy main.go
```

运行编译后的程序：

```bash
./go-proxy-easy -listen 0.0.0.0:8080 -user admin -pass 123456
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
- 代理服务器：127.0.0.1
- 端口：8080
- 如果启用了认证，需要输入用户名和密码

## 许可证

MIT License
