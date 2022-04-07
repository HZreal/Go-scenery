package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Book struct { // 自定义新的数据类型Book，具有struct的特性         struct是值类型
	title      string
	bookId     int
	author     string
	publicDate string
	Desc       string // 结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）
	//string                 // 匿名字段，默认采用类型名作为字段名，结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个

}

// 自定义构造函数
func newBook(title string, bookId int, author string, publicDate string, desc string) *Book {
	return &Book{
		title:      title,
		bookId:     bookId,
		author:     author,
		publicDate: publicDate,
		Desc:       desc,
	}
}

// 方法Method定义(方法与函数的区别是：函数不属于任何类型，方法属于特定的类型)
func (b Book) getPublicDate() string {
	fmt.Println(b.publicDate)
	return b.publicDate
}

// 方法Method定义 指针类型的接收者
func (b *Book) setBookId1(newAuthor string) {
	b.author = newAuthor
}

// 方法Method定义 值类型的接收者
func (b Book) setBookId2(newAuthor string) {
	// 传入的b是原Book结构体的值拷贝(副本)
	b.author = newAuthor
}

// 自定义类型，并基于这个类型定义类型方法
type myInt int // 定义类型
func (m myInt) SayHello() { //SayHello 为myInt添加一个SayHello的方法
	fmt.Println("Hello, 我是一个int。")
}

// 结构体嵌套
type Book2 struct {
	helle string
	book  Book
	Book  // 匿名结构体，变量名默认与类型名一致
}

// 结构体'继承'
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s can move\n", a.name)
}

type Dog struct {
	name    string
	feet    int
	*Animal //通过嵌套匿名结构体实现继承，默认字段名为Animal，类型为指针类型
}

func (d *Dog) wang() {
	fmt.Printf("%s can wangwangwang\n", d.name)
	fmt.Printf("%s can wangwangwang\n", d.Animal.name)
}

func structBasics() {
	//使用type关键字来定义自定义类型
	// 类型定义：type newType existType       // 自定义新的数据类型newType，具有已存在existType的特性
	type newInt int // 自定义新的数据类型newInt，具有int的特性

	//TypeAlias只是Type的别名，本质上TypeAlias与Type是同一个类型
	// 别名定义： type TypeAlias = Type
	type aliasInt = int
	//之前见过的rune和byte就是类型别名，他们的定义如下:
	//type byte = uint8
	//type rune = int32
	var a newInt
	var b aliasInt
	fmt.Printf("type of a:%T\n", a)
	fmt.Printf("type of a:%T\n", b)

	// 实例化
	var book1 = Book{"流浪记", 1, "huang", "2019-2-16", "a funny story"} // 值的列表(位置传参)初始化，不能少字段，顺序对应
	//var book2 = Book{title: "流浪记", bookId: 1, author: "huang", publicDate: "2019-2-16", desc: "a funny story"}  // 键值对初始化
	//var book3 = Book{title: "流浪记", bookId: 1, author: "huang"} // 关键字传参可以少值，忽略的默认为空
	//fmt.Println(book1, book2, book3)

	//var book4 Book
	//book4.bookId = 2
	//book4.author = "zzz"
	//fmt.Printf("book4 bookID is %d\n", book4.bookId)
	//fmt.Printf("book4 author is %s\n", book4.author)

	//匿名结构体
	//var user struct{name string; age int}
	//user.name = "hh"
	//user.age = 22

	//创建指针类型结构体
	//使用new关键字对结构体进行实例化，得到的是结构体的地址
	var p2 = new(Book)
	fmt.Printf("%T\n", p2)     // 类型为 *main.Book
	fmt.Printf("pr=%#v\n", p2) // p2=&main.Book{title:"", bookId:0, author:"", publicDate:"", desc:""}

	//取结构体的地址实例化
	p3 := &Book{}
	fmt.Printf("%T\n", p3)     //*main.Book
	fmt.Printf("p3=%#v\n", p3) //p3=&main.Book{title:"", bookId:0, author:"", publicDate:"", desc:""}
	p3.title = "Go"            // 语法糖写法，   其实在底层是(*p3).title = "Go"
	p3.bookId = 6

	//构造函数   Go语言的结构体没有构造函数，我们可以自己实现
	// 调用构造函数
	p4 := newBook("流浪记", 1, "huang", "2019-2-16", "story")
	fmt.Printf("p4 is %v\n", *p4)

	// 方法：作用于特定类型变量的函数，这种特定类型变量叫做接收者Receiver-------接收者的概念就类似于其他语言中的this或者 self
	//方法的定义格式如下：
	//func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
	//	函数体
	//}
	p5 := newBook("流浪记", 1, "huang", "2019-2-16", "story")
	p5.getPublicDate()
	//指针类型的接收者   调用方法时修改接收者指针的任意成员变量，在方法结束后，修改都是有效的
	p5.setBookId1("huuuu") // p5.author被修改
	fmt.Println("p5.author------", p5.author)
	//值类型的接收者    方法作用于值类型接收者时，Go语言会在代码运行时将接收者的值复制一份。在值类型接收者的方法中可以获取接收者的成员值，但修改操作只是针对副本，无法修改接收者变量本身
	p5.setBookId2("haaaa")
	fmt.Println("p5.author------", p5.author) // p5.author未被修改，修改的是其值拷贝的author

	//何时应该使用指针类型接收者？
	//1.需要修改接收者中的值
	//2.接收者是拷贝代价比较大的大对象
	//3.保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。

	//任意类型添加方法
	//接收者的类型可以是任何类型，不仅仅是结构体，任何类型都可以拥有方法
	//比如基于内置的int类型使用type关键字可以定义新的自定义类型，然后为我们的自定义类型添加方法

	// 结构体嵌套
	bb := Book2{helle: "hello", book: Book{"流浪记", 1, "huang", "2019-2-16", "story"}}
	fmt.Println(bb)

	// 结构体的“继承”
	dog := &Dog{name: "dog1", feet: 18, Animal: &Animal{name: "animal1"}}
	dog.move()
	dog.wang()

	// 结构体作为函数参数
	printBook(book1)

	// 定义的指针变量可以存储结构体变量的地址，查看结构体变量地址，可以将 & 符号放置于结构体变量前
	var structPointer *Book
	fmt.Println(structPointer)

	book1Pointer := &book1
	fmt.Println(book1Pointer.bookId)
}

//Student 学生
type Student struct {
	ID     int    `json:"id" bson:"id" xml:"id"` // Tag是结构体的元信息，可以在运行的时候通过反射的机制读取出来，通过指定tag实现json序列化该字段时的key
	Gender string `json:"gender"`                // json序列化是默认使用字段名作为key
	Name   string
	desc   string // 私有不能被json包访问
}

//Class 班级
type Class struct {
	Title    string
	Students []*Student
}

// 结构体与JSON序列化
func jsonSerializer() {
	c := &Class{
		Title:    "101",
		Students: make([]*Student, 0, 200),
	}
	for i := 0; i < 10; i++ {
		stu := &Student{
			Name:   fmt.Sprintf("stu%02d", i),
			Gender: "男",
			ID:     i,
		}
		c.Students = append(c.Students, stu)
	}
	// JSON序列化：结构体-->JSON格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data)
	// JSON反序列化：JSON格式的字符串-->结构体
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu0"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)

}

// 结构体标签（Tag）
func structTag() {
	// Tag在结构体字段的后方定义，由一对反引号包裹起来，具体的格式如下：
	// `key1:"value1" key2:"value2"`
	// 结构体标签由一个或多个键值对组成。键与值使用冒号分隔，值用双引号括起来。键值对之间使用一个空格分隔。
	// 注意事项： 为结构体编写Tag时，必须严格遵守键值对的规则。结构体标签的解析代码的容错能力很差，一旦格式写错，编译和运行时都不会提示任何错误，通过反射也无法正确取值。例如不要在key和value之间添加空格
	// 示例如上 type Student struct

	stu := &Student{ID: 2, Gender: "boy"}
	t := reflect.TypeOf(stu)
	// 获取第一个字段的Struct Tag的值
	f0 := t.Elem().Field(0)
	println(f0.Tag.Get("json"))
	println(f0.Tag.Get("bson"))
	println(f0.Tag.Get("xml"))

}

func main() {
	//structBasics()
	//jsonSerializer()
	structTag()

}

func printBook(book Book) {
	fmt.Println(book.bookId)
}
