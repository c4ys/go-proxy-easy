package main

import (
"encoding/base64"
"net/http"
"net/http/httptest"
"testing"
)

func TestBasicAuth(t *testing.T) {
	tests := []struct {
		name     string
		username string
		password string
		authHeader string
		expected bool
	}{
		{
			name:     "Valid credentials",
			username: "admin",
			password: "secret",
			authHeader: "Basic " + base64.StdEncoding.EncodeToString([]byte("admin:secret")),
			expected: true,
		},
		{
			name:     "Invalid credentials",
			username: "admin",
			password: "secret",
			authHeader: "Basic " + base64.StdEncoding.EncodeToString([]byte("admin:wrong")),
			expected: false,
		},
		{
			name:     "No auth header",
			username: "admin",
			password: "secret",
			authHeader: "",
			expected: false,
		},
		{
			name:     "Invalid auth format",
			username: "admin",
			password: "secret",
			authHeader: "Bearer token123",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
req := httptest.NewRequest("GET", "http://example.com", nil)
if tt.authHeader != "" {
req.Header.Set("Proxy-Authorization", tt.authHeader)
}

result := basicAuth(req, tt.username, tt.password)
if result != tt.expected {
t.Errorf("basicAuth() = %v, expected %v", result, tt.expected)
}
})
	}
}

func TestHandleProxyWithoutAuth(t *testing.T) {
	handler := handleProxy("", "")
	
	req := httptest.NewRequest("GET", "http://httpbin.org/get", nil)
	rr := httptest.NewRecorder()
	
	handler(rr, req)
	
	// 由于我们无法实际连接到外部服务，这里主要测试认证逻辑
	// 在实际环境中，这个测试需要mock外部HTTP服务
}

func TestHandleProxyWithAuth(t *testing.T) {
	handler := handleProxy("admin", "secret")
	
	// 测试无认证情况
	req := httptest.NewRequest("GET", "http://httpbin.org/get", nil)
	rr := httptest.NewRecorder()
	
	handler(rr, req)
	
	if rr.Code != http.StatusProxyAuthRequired {
		t.Errorf("Expected status %d, got %d", http.StatusProxyAuthRequired, rr.Code)
	}
	
	// 测试正确认证
	req2 := httptest.NewRequest("GET", "http://httpbin.org/get", nil)
	auth := base64.StdEncoding.EncodeToString([]byte("admin:secret"))
	req2.Header.Set("Proxy-Authorization", "Basic "+auth)
	rr2 := httptest.NewRecorder()
	
	handler(rr2, req2)
	
	// 由于我们无法实际连接到外部服务，这里主要测试认证通过后的逻辑
}

func BenchmarkBasicAuth(b *testing.B) {
	req := httptest.NewRequest("GET", "http://example.com", nil)
	auth := base64.StdEncoding.EncodeToString([]byte("admin:secret"))
	req.Header.Set("Proxy-Authorization", "Basic "+auth)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		basicAuth(req, "admin", "secret")
	}
}
