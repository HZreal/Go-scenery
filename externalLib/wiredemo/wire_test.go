package wiredemo

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"
)

/*
*
测试启动服务
*/
func TestStartServer(t *testing.T) {
	controller, err := InitializeUserController()
	if err != nil {
		t.Fatalf("failed to initialize user controller: %v", err)
	}

	http.HandleFunc("/user", controller.GetUserByID)
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		t.Fatalf("failed to start server: %v", err)
	}
}

/*
*
测试访问服务
*/
func TestRequest(t *testing.T) {
	// 发送请求到由 Test_Main 启动的服务器
	resp, err := http.Get("http://localhost:8080/user")
	if err != nil {
		t.Fatalf("failed to send GET request: %v", err)
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status OK; got %v", resp.Status)
	}

	// 读取响应主体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	// 在这里检查响应主体内容是否符合预期
	log.Printf("Response body: %s", string(body))
}

func TestStartServerAndRequest(t *testing.T) {
	// Initialize the controller using the wire-generated function
	controller, err := InitializeUserController()
	if err != nil {
		t.Fatalf("failed to initialize user controller: %v", err)
	}

	// Start the HTTP server in a goroutine
	http.HandleFunc("/user", controller.GetUserByID)
	log.Println("Starting server on :8080")
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalf("failed to start server: %v", err)
		}
	}()
	// Give the server some time to start
	time.Sleep(1 * time.Second)

	// Send a request to the server
	resp, err := http.Get("http://localhost:8080/user")
	if err != nil {
		t.Fatalf("failed to send GET request: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status OK; got %v", resp.Status)
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	// Log the response body for demonstration purposes
	log.Printf("Response body: %s", string(body))
}
