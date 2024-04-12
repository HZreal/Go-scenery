package main

import (
	"fmt"
	// myself pkg from github.com
	"github.com/code-elastic/open-api-client-sdk-go"
)

func demo() {
	var json string
	json, _ = open_api_client_sdk_go.SerializeToJSON(map[string]int{"a": 1, "b": 2})
	fmt.Println("json -->  ", json)
}

func main() {
	demo()
}
