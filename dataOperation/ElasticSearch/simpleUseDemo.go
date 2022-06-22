package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

// Elasticsearch demo
func demo() {
	Client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	// Client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"), elastic.SetBasicAuth("user", "password"))
	if err != nil {
		panic(err)
	}
	fmt.Println("connect to es success")

	person := Person{Name: "huang", Age: 18, Married: false}
	indexResponse, err := Client.Index().Index("user").BodyJson(person).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n", indexResponse.Id, indexResponse.Index, indexResponse.Type)

}
func main() {
	demo()
}
