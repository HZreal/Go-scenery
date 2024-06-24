package main

import (
	// myself pkg from github.com
	"github.com/code-elastic/open-api-client-sdk-go"
)

func demo() {
	var client *open_api_client_sdk_go.OpenClient
	client = open_api_client_sdk_go.NewOpenClient("qwert123456", "asdfgqwertzxcvb")
	client.GetNameByPost("huang")
}

func main() {
	demo()
}
