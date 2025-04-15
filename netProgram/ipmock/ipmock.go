package main

/**
 * @Author nico
 * @Date 2025-04-14
 * @File: ipmock.go
 * @Description:
 */

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type IPEntry struct {
	Region string
	IP     string
}

var ipPool = []IPEntry{
	{"local", "127.0.0.1"},
	{"中国|江苏|扬州", "36.149.36.0"},
	{"武汉", "210.51.200.123"},
	{"黄石", "171.43.106.22"},
	{"美国|加利福尼亚州|费利蒙", "23.142.224.17"},
}

func main() {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjdXN0b21lclR5cGUiOjIsImN1c3RvbWVySWQiOjEsImNvbnRhY3RJZCI6MSwicHJpdmlsZWdlSUQiOjIsInJvbGVJZCI6MSwicm9sZVR5cGUiOjEsInVzZXJUeXBlIjowLCJleHRyYVR5cGUiOjAsImV4dHJhSUQiOjAsInNraXBQZXJtaXNzaW9uIjowLCJsYW5nIjoiIiwiZnJvbVdoaWNoVXNlciI6MCwiZnJvbVdoaWNoVXNlclR5cGUiOjAsImZyb21XaGljaFVzZXJSb2xlVHlwZSI6MCwib3JnYW5pemF0aW9uIjp7fSwiZXhwIjoxNzQ0Nzk3NjY4fQ.99f6Yw0s9AqfcW8dCum9mPelzBPiMQr66tYNLllr0rs" // 替换成真实 token
	url := "https://localhost:5173/v1/api/route/trunkGroup/get"

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: 8 * time.Second,
	}

	for _, entry := range ipPool {
		status := sendRequestWithFakeIP(client, url, token, entry)
		fmt.Printf("[%-6s] 使用 IP: %-15s -> 状态: %s\n", entry.Region, entry.IP, status)
	}
}

func sendRequestWithFakeIP(client *http.Client, url, token string, entry IPEntry) string {
	// 请求体数据
	data := map[string]interface{}{
		"id": 20,
	}
	body, _ := json.Marshal(data)

	// 构建请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Sprintf("请求构建失败: %v", err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Customer-Type", "customer_operation")
	req.Header.Set("Language", "zh-cn")
	req.Header.Set("Origin", "https://localhost:5173")
	req.Header.Set("Referer", "https://localhost:5173/content/service/trunkGroup")

	// 模拟来源 IP（伪造客户端 IP）
	req.Header.Set("X-Forwarded-For", entry.IP)
	req.Header.Set("Client-IP", entry.IP)
	req.Header.Set("X-Real-IP", entry.IP)

	// 发起请求
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Sprintf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	return resp.Status
}
