package main

import "fmt"

type Book struct {
	title      string
	bookId     int
	author     string
	publicDate string
	desc       string
}

type inter interface {
}

func main() {
	var book1 = Book{"流浪记", 1, "huang", "2019-2-16", "a funny story"} // 位置传参不能少值
	var book2 = Book{title: "流浪记", bookId: 1, author: "huang", publicDate: "2019-2-16", desc: "a funny story"}
	var book3 = Book{title: "流浪记", bookId: 1, author: "huang"} // 关键字传参可以少值，忽略的默认为空
	fmt.Println(book1, book2, book3)

	var book4 Book
	book4.bookId = 2
	book4.author = "zzz"
	fmt.Printf("book4 bookID is %d\n", book4.bookId)
	fmt.Printf("book4 author is %s\n", book4.author)

	// 结构体作为函数参数
	printBook(book1)

	// 定义的指针变量可以存储结构体变量的地址，查看结构体变量地址，可以将 & 符号放置于结构体变量前
	var structPointer *Book
	fmt.Println(structPointer)

	book1Pointer := &book1
	fmt.Println(book1Pointer.bookId)

}

func printBook(book Book) {
	fmt.Println(book.bookId)
}
