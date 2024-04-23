package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
)

type Book struct { // 自定义新的数据类型Book，具有struct的特性         struct是值类型
	title      string
	bookId     int
	author     string
	publicDate string
	Desc       string // 结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）
	// string                 // 匿名字段，默认采用类型名作为字段名，结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个

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
func (m myInt) SayHello() { // SayHello 为myInt添加一个SayHello的方法
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
	*Animal // 通过嵌套匿名结构体实现继承，默认字段名为Animal，类型为指针类型
}

func (d *Dog) wang() {
	fmt.Printf("%s can wangwangwang\n", d.name)
	fmt.Printf("%s can wangwangwang\n", d.Animal.name)
}

// //////////////////////// 值接收器 & 指针接收器 /////////////////////////////
// 如果一个接口方法是用指针接收器定义的，那么只有指针类型的变量才能实现该接口。
// 如果一个接口方法是用值接收器定义的，那么值类型和指针类型的变量都可以实现该接口。

type Describer interface {
	Describe() string
}
type Person struct {
	Name string
	Age  int
}

type Person2 struct {
	Name string
	Age  int
}

// value receiver
func (p Person) Describe() string {
	return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
}

// point receiver
func (p *Person2) Describe() string {
	return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
}

func testValueReceiver() {
	p := Person{"Alice", 30}
	var d Describer = p // Person 值类型
	fmt.Println(d.Describe())

	pd := &p
	d = pd // Person 指针类型
	fmt.Println(d.Describe())
}

func testPointReceiver() {
	p := Person2{"Bob", 25}
	// var d Describer = p  // 这会产生编译错误，因为 p 是一个值
	var d Describer = &p // 这是正确的，因为 &p 是一个指针
	fmt.Println(d.Describe())
}

func structBasics() {
	// 使用type关键字来定义自定义类型
	// 类型定义：type newType existType       // 自定义新的数据类型newType，具有已存在existType的特性
	type newInt int // 自定义新的数据类型newInt，具有int的特性

	// TypeAlias只是Type的别名，本质上TypeAlias与Type是同一个类型
	// 别名定义： type TypeAlias = Type
	type aliasInt = int
	// 之前见过的rune和byte就是类型别名，他们的定义如下:
	// type byte = uint8
	// type rune = int32
	var a newInt
	var b aliasInt
	fmt.Printf("type of a:%T\n", a)
	fmt.Printf("type of a:%T\n", b)

	// 实例化
	var book1 = Book{"流浪记", 1, "huang", "2019-2-16", "a funny story"} // 值的列表(位置传参)初始化，不能少字段，顺序对应
	// var book2 = Book{title: "流浪记", bookId: 1, author: "huang", publicDate: "2019-2-16", desc: "a funny story"}  // 键值对初始化
	// var book3 = Book{title: "流浪记", bookId: 1, author: "huang"} // 关键字传参可以少值，忽略的默认为空
	// fmt.Println(book1, book2, book3)

	// var book4 Book
	// book4.bookId = 2
	// book4.author = "zzz"
	// fmt.Printf("book4 bookID is %d\n", book4.bookId)
	// fmt.Printf("book4 author is %s\n", book4.author)

	// 匿名结构体
	// var user struct{name string; age int}
	// user.name = "hh"
	// user.age = 22

	// 创建指针类型结构体
	// 使用new关键字对结构体进行实例化，得到的是结构体的地址
	var p2 = new(Book)
	fmt.Printf("%T\n", p2)     // 类型为 *main.Book
	fmt.Printf("pr=%#v\n", p2) // p2=&main.Book{title:"", bookId:0, author:"", publicDate:"", desc:""}

	// 取结构体的地址实例化
	p3 := &Book{}
	fmt.Printf("%T\n", p3)     // *main.Book
	fmt.Printf("p3=%#v\n", p3) // p3=&main.Book{title:"", bookId:0, author:"", publicDate:"", desc:""}
	p3.title = "Go"            // 语法糖写法，   其实在底层是(*p3).title = "Go"
	p3.bookId = 6

	// 构造函数   Go语言的结构体没有构造函数，我们可以自己实现
	// 调用构造函数
	p4 := newBook("流浪记", 1, "huang", "2019-2-16", "story")
	fmt.Printf("p4 is %v\n", *p4)

	// 方法：作用于特定类型变量的函数，这种特定类型变量叫做接收者Receiver-------接收者的概念就类似于其他语言中的this或者 self
	// 方法的定义格式如下：
	// func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
	//	函数体
	// }
	p5 := newBook("流浪记", 1, "huang", "2019-2-16", "story")
	p5.getPublicDate()
	// 指针类型的接收者   调用方法时修改接收者指针的任意成员变量，在方法结束后，修改都是有效的
	p5.setBookId1("huuuu") // p5.author被修改
	fmt.Println("p5.author------", p5.author)
	// 值类型的接收者    方法作用于值类型接收者时，Go语言会在代码运行时将接收者的值复制一份。在值类型接收者的方法中可以获取接收者的成员值，但修改操作只是针对副本，无法修改接收者变量本身
	p5.setBookId2("haaaa")
	fmt.Println("p5.author------", p5.author) // p5.author未被修改，修改的是其值拷贝的author

	// 何时应该使用指针类型接收者？
	// 1.需要修改接收者中的值
	// 2.接收者是拷贝代价比较大的大对象
	// 3.保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。

	// 任意类型添加方法
	// 接收者的类型可以是任何类型，不仅仅是结构体，任何类型都可以拥有方法
	// 比如基于内置的int类型使用type关键字可以定义新的自定义类型，然后为我们的自定义类型添加方法

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

// Student 学生
type Student struct {
	ID     int    `json:"id" bson:"id" xml:"id"` // Tag是结构体的元信息，可以在运行的时候通过反射的机制读取出来，通过指定tag实现json序列化该字段时的key
	Gender string `json:"gender"`                // json序列化是默认使用字段名作为key
	Name   string
	desc   string // 私有不能被json包访问
}

// Class 班级
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

// 方法
func methodBasics() {
	// 1. 方法定义
	// Golang 方法总是绑定对象实例，并隐式将实例作为第一实参 (receiver)。
	//		• 只能为当前包内命名类型定义方法。
	//		• 参数 receiver 可任意命名。如方法中未曾使用 ，可省略参数名。
	//		• 参数 receiver 类型可以是 T 或 *T。基类型 T 不能是接口或指针。
	//		• 不支持方法重载，receiver 只是参数签名的组成部分。
	//		• 可用实例 value 或 pointer 调用全部方法，编译器自动转换。
	// 一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。
	// 所有给定类型的方法属于该类型的方法集
	// 1.1.1. 方法定义： func (recevier type) methodName(参数列表)(返回值列表){}

	// 当接受者是指针时，即使用值类型调用那么函数内部也是对指针的操作。
	// 1.1.2. 普通函数与方法的区别
	// 1.对于普通函数，接收者为值类型时，不能将指针类型的数据直接传递，反之亦然。
	// 2.对于方法（如struct的方法），接收者为值类型时，可以直接用指针类型的变量调用方法，反过来同样也可以。

}

func (s *Student) tTest() {
	fmt.Println(s.Name)
}

// 表达式
func valueAndExpression() {
	u := Student{1, "1", "hu", "hhhhhhh"}
	u.tTest()

	mValue := u.tTest
	mValue() // 隐式传递 receiver   会复制 receiver

	mExpression := (*Student).tTest
	mExpression(&u) // 显式传递 receiver
}

// Sayer 接口
type Sayer interface {
	say()
}

// 定义dog和cat两个结构体：
type dog struct{}
type cat struct{}

// dog实现了Sayer接口
func (d dog) say() {
	fmt.Println("汪汪汪")
}

// cat实现了Sayer接口
func (c cat) say() {
	fmt.Println("喵喵喵")
}
func interfaceTypeVariable() {
	var x Sayer // 声明一个Sayer类型的变量x   能够存储dog和cat类型的变量
	a := cat{}  // 实例化一个cat
	b := dog{}  // 实例化一个dog
	x = a       // 可以把cat实例直接赋值给x
	x.say()     // 喵喵喵
	x = b       // 可以把dog实例直接赋值给x
	x.say()     // 汪汪汪
}

type Mover interface {
	move()
}
type fish struct{}

//	func (d fish) move() {
//		fmt.Println("鱼会游")
//	}
func (d fish) move() {
	fmt.Println("鱼会游")
}

// WashingMachine 洗衣机
type WashingMachine interface {
	wash()
	dry()
}

// 甩干器
type dryer struct{}

// 实现WashingMachine接口的dry()方法
func (d dryer) dry() {
	fmt.Println("甩一甩")
}

// 海尔洗衣机
type haier struct {
	dryer // 嵌入甩干器，即可调用甩干器的dry方法
}

// 实现WashingMachine接口的wash()方法
func (h haier) wash() {
	fmt.Println("洗刷刷")
}

// 接口嵌套：创造出新的接口
type animal interface {
	Sayer
	Mover
}

// 空接口作为函数参数
func funcWithEmptyInterfaceAsParams(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}

type i interface {
	Say(s string) string
}
type People struct {
	Name string
}

func (p *People) Say(s string) string {
	fmt.Println(p.Name)
	return s
}

// 接口
func interfaceBasics() {
	// 接口（interface）定义了一个对象的行为规范，只定义规范不实现，由具体的对象来实现规范的细节
	// 1.1.1. 接口类型
	// 在Go语言中接口（interface）是一种类型，一种抽象的类型   牢记接口（interface）是一种类型
	// interface是一组method的集合，是duck-type programming的一种体现。接口做的事情就像是定义一个协议（规则），只要一台机器有洗衣服和甩干的功能，我就称它为洗衣机。不关心属性（数据），只关心行为（方法）

	// 1.1.2. 为什么要使用接口
	// 比如三角形，四边形，圆形都能计算周长和面积，我们能不能把它们当成“图形”来处理呢？
	// 比如销售、行政、程序员都能计算月薪，我们能不能把他们当成“员工”来处理呢？
	// Go语言中为了解决类似上面的问题，就设计了接口这个概念。接口区别于我们之前所有的具体类型，接口是一种抽象的类型。当你看到一个接口类型的值时，你不知道它是什么，唯一知道的是通过它的方法能做什么

	// 1.1.3. 接口的定义
	//		接口是一个或多个方法签名的集合。
	//		任何类型的方法集中只要拥有该接口'对应的全部方法'签名。就表示它 "实现" 了该接口，无须在该类型上显式声明实现了哪个接口。这称为Structural Typing。
	//		所谓对应方法，是指有相同名称、参数列表 (不包括参数名) 以及返回值。
	//		当然，该类型还可以有其他方法。
	//
	//		接口只有方法声明，没有实现，没有数据字段。
	//		接口可以匿名嵌入其他接口，或嵌入到结构中。
	//		对象赋值给接口时，会发生拷贝，而接口内部存储的是指向这个复制品的指针，既无法修改复制品的状态，也无法获取指针。
	//		只有当接口存储的类型和对象都为nil时，接口才等于nil。
	//		接口调用不会做receiver的自动转换。
	//		接口同样支持匿名字段方法。
	//		接口也可实现类似OOP中的多态。
	//		空接口可以作为任何类型数据的容器。
	//		一个类型可实现多个接口。
	//		接口命名习惯以 er 结尾。

	// 每个接口由数个方法组成，接口的定义格式如下：
	//		type 接口类型名 interface{
	//			方法名1( 参数列表1 ) 返回值列表1
	//			方法名2( 参数列表2 ) 返回值列表2
	//			…
	//		}
	// 其中：
	// 1.接口名：使用type将接口定义为自定义的类型名。Go语言的接口在命名时，一般会在单词后面添加er，如有写操作的接口叫Writer，有字符串功能的接口叫Stringer等。接口名最好要能突出该接口的类型含义。
	// 2.方法名：当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
	// 3.参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以省略。

	// 注意：只有当有两个或两个以上的具体类型必须以相同的方式进行处理时才需要定义接口。不要为了接口而写接口，那样只会增加不必要的抽象，导致不必要的运行时损耗

	// 1.1.4. 实现接口的条件
	// 一个对象只要全部实现了接口中的方法，那么就实现了这个接口。换句话说，接口就是一个需要实现的方法列表

	// 1.1.5. 接口类型变量
	// 接口类型变量能够存储所有实现了该接口的实例
	// interfaceTypeVariable()

	// 1.1.6. 值接收者实现接口
	// 使用值接收者实现接口之后，不管是struct还是*struct类型的变量都可以赋值给该接口变量。因为Go语言中有对指针类型变量求值的语法糖，dog指针fugui内部会自动求值*fugui
	// var x1 Mover
	// var wangcai1 = fish{} // 旺财是fish类型
	// x1 = wangcai1         // x可以接收fish类型
	// var fugui1 = &fish{}  // 富贵是*fish类型
	// x1 = fugui1           // x可以接收*fish类型
	// x1.move()

	// 1.1.7. 指针接收者实现接口
	// 使用值接收者实现接口时，只能将*struct类型的变量赋值给该接口变量，传struct类型给接口变量编译不通过
	// var x2 Mover
	// var wangcai2 = fish{} // 旺财是fish类型
	// x2 = wangcai2         // 当实现Mover接口的是*fish类型时，   x不可以接收fish类型
	// var fugui2 = &fish{}  // 富贵是*fish类型
	// x2 = fugui2           // x可以接收*fish类型
	// x2.move()

	// 1.1.8. 值接收者和指针接收者实现接口的区别
	// 实现接口的是struct类型时，接口类型变量可以接收struct类型、*struct类型(语法糖)
	// 实现接口的是*struct类型时，接口类型变量不可以接收struct类型

	// 1.2. 类型与接口的关系
	// 1.2.1. 一个类型实现多个接口
	// 一个类型可以同时实现多个接口，而接口间彼此独立，不知道对方的实现。 例如，狗可以叫，也可以动
	// 1.2.2. 多个类型实现同一接口。例如狗可以动，汽车也可以动
	// 一个接口的方法，不一定需要由一个类型完全实现
	// 接口的方法可以通过在类型中嵌入其他类型或者结构体来实现，当前类型可以调用被嵌入类型的方法
	// h := haier{}
	// h.wash()
	// h.dry()

	// 1.2.3. 接口嵌套

	// 1.3. 空接口
	// 1.3.1. 空接口的定义
	// 空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口
	// 空接口类型的变量可以存储任意类型的变量
	// 定义一个空接口x
	// var x interface{}
	// s := "pprof.cn"
	// x = s
	// fmt.Printf("type:%T value:%v\n", x, x)
	// i := 100
	// x = i
	// fmt.Printf("type:%T value:%v\n", x, x)
	// b := true
	// x = b
	// fmt.Printf("type:%T value:%v\n", x, x)

	// 1.3.2. 空接口的应用
	// a. 空接口作为函数的参数，可以接收任意类型的函数参数
	// funcWithEmptyInterfaceAsParams(1)

	// b. 空接口作为map的值，可以保存任意值的字典
	// var studentInfo = make(map[string]interface{})
	// studentInfo["name"] = "李白"
	// studentInfo["age"] = 18
	// studentInfo["married"] = false
	// fmt.Println(studentInfo)

	// 1.3.3. 类型断言
	// 一个接口的值（简称接口值）是由一个具体类型和具体类型的值两部分组成的。这两部分分别称为接口的动态类型和动态值
	var w io.Writer       // 声明一个接口，此时动态类型和动态值均为nil
	w = os.Stdout         // 动态类型为*os.File，动态值为此类型下的零值
	w = new(bytes.Buffer) // 动态类型为*bytes.Buffer，动态值为此类型下的零值
	w = nil
	n, _ := w.Write([]byte("hello"))
	fmt.Println(n)

	// 判断空接口中的值这个时候就可以使用类型断言，其语法格式：
	// x.(T)     x：表示类型为interface{}的变量   T：表示断言x可能是的类型。
	// 该语法返回两个参数，第一个参数是x转化为T类型后的变量，第二个值是一个布尔值，若为true则表示断言成功，为false则表示断言失败
	var x interface{}
	x = "baidu"
	v, ok := x.(string)
	if ok {
		fmt.Println("类型断言成功")
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}

	// 当某种类型 T 实现了某接口 i，则可以初始化该类型并赋给该接口
	var ss i
	ss = new(People)
	ss.(*People).Name = "fff"
	ss.Say("ddd")

}

func printBook(book Book) {
	fmt.Println(book.bookId)
}

// //////////////////// 接口作为参数 //////////////////////

type Reader interface {
	Read() int
}

type MyReader struct {
	a, b int
}

func (m *MyReader) Read() int {
	return m.a + m.b
}

func DoJob(r Reader) {
	fmt.Printf("myReader is %d\n", r.Read())
}

func interfaceAsParameter() {
	myReader := &MyReader{2, 5}
	DoJob(myReader)
}

// //////////////////// 接口嵌套 //////////////////////

type A interface {
	run1()
}

type B interface {
	run2()
}

// 定义继承接口C
type C interface {
	A
	B
	run3()
}

type Runner struct{}

// 实现接口A的方法
func (r Runner) run1() {
	fmt.Println("run1!!!!")
}

// 实现接口B的方法
func (r Runner) run2() {
	fmt.Println("run2!!!!")
}

func (r Runner) run3() {
	fmt.Println("run3!!!!")
}

func interfaceEmbed() {
	r := Runner{}
	r.run1()
	r.run2()
	r.run3()
}

func main() {
	// structBasics()
	// testValueReceiver()
	// testPointReceiver()
	// jsonSerializer()
	// structTag()
	// valueAndExpression()
	interfaceBasics()
	interfaceAsParameter()
	interfaceEmbed()
}
