package main

import (
"encoding/base64"
"flag"
"fmt"
"io"
"log"
"net"
"net/http"
"strings"
)

// 检查代理认证
func basicAuth(r *http.Request, username, password string) bool {
	auth := r.Header.Get("Proxy-Authorization")
	if auth == "" {
		return false
	}
	const prefix = "Basic "
	if !strings.HasPrefix(auth, prefix) {
		return false
	}
	decoded, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
	if err != nil {
		return false
	}
	parts := strings.SplitN(string(decoded), ":", 2)
	if len(parts) != 2 {
		return false
	}
	return parts[0] == username && parts[1] == password
}

// 处理代理请求
func handleProxy(username, password string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 如果设置了用户名和密码，进行认证
		if username != "" || password != "" {
			if !basicAuth(r, username, password) {
				w.Header().Set("Proxy-Authenticate", "Basic realm=\"Proxy\"")
				w.WriteHeader(http.StatusProxyAuthRequired)
				w.Write([]byte("407 Proxy Authentication Required"))
				return
			}
		}

		// 根据不同请求方式处理
		if r.Method == http.MethodConnect {
			// 处理HTTPS请求
			handleTunneling(w, r)
		} else {
			// 处理HTTP请求
			handleHTTP(w, r)
		}
	}
}

// 处理HTTPS隧道
func handleTunneling(w http.ResponseWriter, r *http.Request) {
	destConn, err := net.Dial("tcp", r.Host)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)
	
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}
	
	clientConn, _, err := hijacker.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	// 双向传输数据
	go transfer(destConn, clientConn)
	go transfer(clientConn, destConn)
}

// 传输数据
func transfer(destination io.WriteCloser, source io.ReadCloser) {
	defer destination.Close()
	defer source.Close()
	io.Copy(destination, source)
}

// 处理HTTP请求
func handleHTTP(w http.ResponseWriter, r *http.Request) {
	resp, err := http.DefaultTransport.RoundTrip(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	// 复制HTTP头
	for k, vv := range resp.Header {
		for _, v := range vv {
			w.Header().Add(k, v)
		}
	}

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func main() {
	// 定义命令行参数
	listen := flag.String("listen", "0.0.0.0:8080", "监听地址和端口")
	user := flag.String("user", "", "用户名")
	pass := flag.String("pass", "", "密码")
	flag.Parse()

	// 输出启动信息
	fmt.Printf("HTTP代理服务启动: %s\n", *listen)
	if *user != "" {
		fmt.Printf("需要认证，用户名: %s\n", *user)
	}

	// 启动服务器
	log.Fatal(http.ListenAndServe(*listen, handleProxy(*user, *pass)))
}
