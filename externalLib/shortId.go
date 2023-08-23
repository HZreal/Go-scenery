package main

import (
	"fmt"
	"github.com/teris-io/shortid"
	"time"
)

func testShortid1() {
	value, err := shortid.Generate()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(value) // 例如：
}

func testShortid2() {
	shortId := shortid.MustNew(2, shortid.DefaultABC, uint64(time.Now().UnixNano()))
	value, err := shortId.Generate()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(value)
}

func main() {
	// testShortid1()
	testShortid2()
}
