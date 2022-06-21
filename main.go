package main

import (
	"fmt"
	split "goBasics/split"
	"log"
)

func main() {
	res := split.Split("a:b:c", ":")
	fmt.Println("res ---->  ", res)

	log.Fatal("some logs")

}
