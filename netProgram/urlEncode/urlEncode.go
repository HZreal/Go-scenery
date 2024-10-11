package main

/**
 * @Author elastic·H
 * @Date 2024-10-11
 * @File: urlEncode.go
 * @Description:
 */

import (
	"fmt"
	"net/url"
)

func queryEscapeUnescape() {
	// 原始 URL 和查询参数
	baseURL := "https://example.com/search"
	queryParams := "q=Golang URL编码&lang=zh-CN"

	// 手动构造原始 URL
	originalURL := baseURL + "?" + queryParams
	fmt.Println("Original URL:", originalURL)

	// 1. URL 编码
	// 将查询参数部分进行编码
	encodedQuery := url.QueryEscape("Golang URL编码")
	encodedLang := url.QueryEscape("zh-CN")

	// 构造完整的编码后的 URL
	encodedURL := fmt.Sprintf("%s?q=%s&lang=%s", baseURL, encodedQuery, encodedLang)
	fmt.Println("Encoded URL:", encodedURL)

	// 2. URL 解码
	// 将编码后的查询参数解码回来
	decodedQuery, _ := url.QueryUnescape(encodedQuery)
	decodedLang, _ := url.QueryUnescape(encodedLang)

	// 构造完整的解码后的 URL
	decodedURL := fmt.Sprintf("%s?q=%s&lang=%s", baseURL, decodedQuery, decodedLang)
	fmt.Println("Decoded URL:", decodedURL)
}

func pathEscapeUnescape() {
	path := "/foo/bar 你好"
	encodedPath := url.PathEscape(path)
	fmt.Println("Encoded Path:", encodedPath) // 输出: /foo/bar%20%E4%BD%A0%E5%A5%BD

	decodedPath, _ := url.PathUnescape(encodedPath)
	fmt.Println("Decoded Path:", decodedPath) // 输出: /foo/bar 你好
}

func urlParse() {
	rawURL := "https://example.com/search?q=golang&lang=zh-CN"
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	// 打印 URL 的各个部分
	fmt.Println("Scheme:", parsedURL.Scheme)  // 输出: Scheme: https
	fmt.Println("Host:", parsedURL.Host)      // 输出: Host: example.com
	fmt.Println("Path:", parsedURL.Path)      // 输出: Path: /search
	fmt.Println("Query:", parsedURL.RawQuery) // 输出: Query: q=golang&lang=zh-CN

	// 解析查询参数
	queryParams := parsedURL.Query()
	fmt.Println("q param:", queryParams.Get("q"))       // 输出: q param: golang
	fmt.Println("lang param:", queryParams.Get("lang")) // 输出: lang param: zh-CN
}

func urlEscapedPath() {
	rawURL := "https://example.com/foo bar/你好"
	parsedURL, _ := url.Parse(rawURL)

	// 获取已编码的路径部分
	fmt.Println("Escaped Path:", parsedURL.EscapedPath()) // 输出: /foo%20bar/%E4%BD%A0%E5%A5%BD
}

func urlValues() {
	// 使用 url.Values 构造查询参数
	queryParams := url.Values{}
	queryParams.Add("q", "golang")
	queryParams.Add("lang", "zh-CN")
	queryParams.Add("tag", "go")
	queryParams.Add("tag", "programming")

	// 编码查询参数
	encodedQuery := queryParams.Encode()
	fmt.Println("Encoded Query:", encodedQuery) // 输出: q=golang&lang=zh-CN&tag=go&tag=programming
}

func urlUrl() {
	parsedURL := &url.URL{
		Scheme:   "https",
		Host:     "example.com",
		Path:     "/search",
		RawQuery: "q=golang&lang=zh-CN",
	}

	fmt.Println("Complete URL:", parsedURL.String()) // 输出: https://example.com/search?q=golang&lang=zh-CN

}

func main() {
	queryEscapeUnescape()
	pathEscapeUnescape()
	urlParse()
	urlEscapedPath()
	urlValues()
	urlUrl()
}
