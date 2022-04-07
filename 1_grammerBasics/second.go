package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// 数组
func arrBasics() {
	// 1.同种数据类型的固定长度的序列
	// 2.定义：var a [len]int  数组长度必须是常量，且是类型的组成部分。一旦定义，长度不能变
	var arr [10]float32 // 声明但未初始化的元素值为 0
	//var _arr1 [5]float32 = [5]float32{1.0, 2.1, 3.0, 4.0, 5.0}
	var arr1 = [5]float32{1.0, 2.1, 3.0, 4.0, 5.0}
	arr2 := [6]float32{1.0, 2.1, 3.0}  // 未初始化的元素值为 0
	arr3 := [5]float32{1: 2.0, 3: 7.0} // 使用索引号初始化元素，将索引为 1 和 3 的元素初始化
	arr4 := [...]int{1, 2, 4}          // 通过初始化值确定数组长度
	arr5 := [...]struct {              // 结构体数组
		name string
		age  uint8
	}{
		{"huang1", 22},
		{"huang2", 23},
	}
	twoDimensionArr := [...][2]int{{1, 1}, {2, 2}, {3, 3}} // 二维数组，第2维度不能用 "..."
	fmt.Println(arr, arr1, arr2, arr3, arr4, arr5, twoDimensionArr)
	// 3.长度是数组类型的一部分，因此，var a [5]int和var a [10]int是不同的类型
	// 4.数组可以通过下标进行访问，下标是从0开始，最后一个元素下标是：len-1
	//for i := 1; i <= len(arr1); i++ {
	//}
	//for index, value := range arr1 {
	//	fmt.Println(index, value)
	//}
	//for k1, v1 := range twoDimensionArr{
	//	for k2, v2 := range v1{
	//		fmt.Printf("(%d, %d) = %d\n", k1, k2, v2)
	//	}
	//}
	// 5. 访问越界，如果下标在数组合法范围之外，则触发访问越界，会panic
	// 6. 数组是值类型，赋值和传参会复制整个数组，而不是指针。因此改变副本的值，不会改变本身的值。
	// 7.支持 "=="、"!=" 操作符，因为内存总是被初始化过的。
	// 8.指针数组：[n]*T  即一个长度为n的数组，元素类型为指针(*int， *string)；   数组指针：*[n]T 即一个指针，指向某个类型为[n]T的数组

	// 值拷贝行为会造成性能问题，通常会建议使用 slice，或数组指针
	//testArr := [3]int{1, 2, 3}
	//fmt.Printf("testArr Address is %p\n", &testArr)
	//ret := arrTest(testArr)    // testArr以形参传入函数时进行了一次深拷贝
	//fmt.Printf("ret address is %p\n", &ret)

	// 内置函数 len 和 cap 都返回数组长度
	//println(len(arr1), cap(arr1))

	// 数组拷贝和传参
	testArr1 := [...]int{1, 2, 3}
	fmt.Println(testArr1)
	fmt.Printf("testArr1 address is %p\n", &testArr1)
	arrTest2(&testArr1)                               // 传入数组指针
	fmt.Println(testArr1)                             // testArr1被修改了， 但是testArr1的地址不变
	fmt.Printf("testArr1 address is %p\n", &testArr1) // 地址不变

	// 数组和
	//testArr := []int{1, 2, 3, 4, 5}
	//ret := sumArr(testArr)
	//fmt.Println(ret)

}

// 切片   -----> 数组的一个引用
func _slice() {
	// 1. 切片：切片是`数组的一个引用`，因此`切片是引用类型`。但`自身是结构体`，值拷贝传递。
	// 2. 切片的长度可以改变，因此，切片是一个可变的数组。
	// 3. 切片遍历方式和数组一样，可以用len()求长度。表示可用元素数量，读写操作不能超过该限制。
	// 4. cap可以求出slice最大扩张容量，不能超出数组限制。0 <= len(slice) <= len(array)，其中array是slice引用的数组。
	// 5. 切片的定义：var 变量名 []类型，比如 var str []string  var arr []int。
	// 6. 如果 slice == nil，那么 len、cap 结果都等于 0。

	// 1.1.1. 创建切片的各种方式
	//1.声明切片
	var s1 []int
	if s1 == nil {
		fmt.Println("空")
	} else {
		fmt.Println("非空")
	}
	// 2. :=
	s2 := []int{}
	// 3. make()内建函数  make用来为 slice，map 或 chan 类型分配内存和初始化一个对象(注意：只能用在这三种类型上) 第一个参数也是一个类型而不是一个值，跟 new 不同的是，make 返回类型的本身而不是指针，而返回值也依赖于具体传入的类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了
	//补充：new()内建函数  new用来分配内存，第一个参数是一个类型，不是一个值，返回一个指向类型为T、值为零的指针，即new(t)分配了零值填充的类型为T内存空间，并且返回其地址，即一个*t类型的值。返回的永远是类型的指针，指向分配类型的内存地址，适用于值类型如数组和结构体
	var s3 []int = make([]int, 2, 4)
	fmt.Println(s1, s2, s3)
	// 4. 初始化赋值
	var s4 []int = make([]int, 1, 5)
	fmt.Println(s4)
	s5 := []int{1, 2, 3}
	fmt.Println(s5)
	// 5. 从数组切片
	arr0 := [5]int{1, 2, 3, 4, 5}
	var s6 []int
	// 前包后不包
	s6 = arr0[1:4]
	fmt.Println(s6)

	// 1.1.2. 切片初始化
	start := 0
	end := 5
	// 全局：
	var arr1 = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var slice0 []int = arr1[start:end]
	var slice1 []int = arr1[:end]
	var slice2 []int = arr1[start:]
	var slice3 []int = arr1[:]
	var slice4 = arr1[:len(arr1)-1] //去掉切片的最后一个元素
	fmt.Println(slice0, slice1, slice2, slice3, slice4)
	// 局部：
	arr2 := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	slice5 := arr2[start:end]
	slice6 := arr2[:end]
	slice7 := arr2[start:]
	slice8 := arr2[:]
	slice9 := arr2[:len(arr2)-1] //去掉切片的最后一个元素
	fmt.Println(slice5, slice6, slice7, slice8, slice9)

	// 1.1.3. 通过make来创建切片
	//slice  := make([]type, len, cap)    // 参数len设置该slice的长度，参数cap设置该slice的容量（即底层数组的长度），省略 cap，相当于 cap = len
	//读写操作实际目标是底层数组，只需注意索引号的差别
	//可直接创建 slice 对象，自动分配底层数组

	// 1.1.4. 用append内置函数操作切片（切片追加） 向 slice 尾部添加数据，返回新的 slice 对象
	var a = []int{1, 2, 3}
	fmt.Printf("slice a : %v\n", a)
	var b = []int{4, 5, 6}
	fmt.Printf("slice b : %v\n", b)
	c := append(a, b...)
	fmt.Printf("slice c : %v\n", c)
	d := append(c, 7)
	fmt.Printf("slice d : %v\n", d)
	e := append(d, 8, 9, 10)
	fmt.Printf("slice e : %v\n", e)

	// 1.1.5. 超出原 slice.cap 限制，就会重新分配底层数组，即便原数组并未填满
	data1 := [...]int{0, 1, 2, 3, 4, 10: 0}
	s := data1[:2:3]
	s = append(s, 100, 200)       // 一次 append 两个值，超出 s.cap 限制。
	fmt.Println(s, data1)         // 重新分配底层数组，与原数组无关。
	fmt.Println(&s[0], &data1[0]) // 比对底层数组起始指针。

	// 1.1.6. slice中cap重新分配规律

	// 1.1.7. 切片拷贝

	// 1.1.8. slice遍历

	// 1.1.9. 切片resize（调整大小）

	// 1.1.10. 数组和切片的内存布局

	// 1.1.11. 字符串和切片（string and slice

	// 1.1.12. 含有中文字符串

}

// 指针
//Go语言中的函数传参都是值拷贝，当我们想要修改某个变量的时候，我们可以创建一个指向该变量地址的指针变量。传递数据使用指针，而无须拷贝数据。
//类型指针不能进行偏移和运算。Go语言中的指针操作非常简单，只需要记住两个符号：&（取地址）和*（根据地址取值）
func pointBasics() {

	//1.1.1. 指针地址和指针类型
	//v := 5
	//ptr := &v    // v的类型为T
	//fmt.Println(ptr)
	//v:代表被取地址的变量，类型为T
	//ptr:用于接收地址的变量，ptr的类型就为*T，称做T的指针类型。*代表指针。
	//a := 10
	//b  := &a     // 取变量a的地址(即10的存储地址)，将地址(指针)保存到变量b中，所以变量b的类型为指针，值为地址
	//fmt.Printf("a:%d  ptr:%p\n", a, &a) // %p 指针16进制输出                 输出：a:10 ptr:0xc00001a078
	//fmt.Printf("b:%p  type:%T\n", b, b) // %T 相应值的类型                   输出：b:0xc00001a078 type:*int
	//fmt.Println(&b)    // 存储地址(数据10的地址)的地址

	//1.1.2. 指针取值
	//1.对变量进行取地址（&）操作，可以获得这个变量的指针变量。
	//2.指针变量的值是指针地址。
	//3.对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值。
	//aa := 10
	//bb := &aa
	//fmt.Printf("type of bb:%T\n", bb)     // %T 输出类型
	//cc := *bb // 指针取值（根据指针去内存取值）
	//fmt.Printf("type of cc:%T\n", cc)
	//fmt.Printf("value of cc:%v\n", cc)    // %v 以该数据的默认格式输出值

	// 1.1.3.空指针
	//当一个指针被定义后没有分配到任何变量时，它的值为 nil
	//空指针的判断
	//var p *string
	//fmt.Println(p)
	//fmt.Printf("p的值是%v\n", p)
	//if p != nil {
	//	fmt.Println("非空")
	//} else {
	//	fmt.Println("空值")
	//}

	//1.1.4. new和make
	//在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，否则我们的值就没办法存储。
	//而对于值类型的声明不需要分配内存空间，是因为它们在声明的时候已经默认分配好了内存空间。
	//要分配内存，就引出来今天的new和make。 Go语言中new和make是内建的两个函数，主要用来分配内存
	// new函数
	//函数定义 func new(Type) *Type {}    Type表示类型，new函数只接受一个参数，这个参数是一个类型，*Type表示类型指针，new函数返回一个指向该类型内存地址的指针
	//使用new函数得到的是一个类型的指针，并且该指针对应的值为该类型的零值，如：
	//aaa := new(int)
	//bbb := new(bool)
	//fmt.Printf("%T\n", aaa) // *int
	//fmt.Printf("%T\n", bbb) // *bool
	//fmt.Println(*aaa)       // 0
	//fmt.Println(*bbb)       // false

	//var abc *int     // 仅声明了一个指针变量a但没有初始化，不能赋值
	//abc = new(int)   // 指针作为引用类型需要初始化后才会拥有内存空间，才可以给它赋值，用内置的new函数对a进行初始化之后就可以正常对其赋值了
	//*abc = 10
	//fmt.Println(*abc)

	// make函数
	//make也是用于内存分配的，区别于new，它只用于slice、map以及chan的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了
	//函数定义 func make(t Type, size ...IntegerType) Type {}
	//var _map map[string]int               // 只是声明变量b是一个map类型的变量，没有初始化，不能赋值
	//_map = make(map[string]int, 10)       // 用make函数进行初始化操作之后，才能对其进行键值对赋值
	//_map["测试"] = 100
	//fmt.Println(_map)

	//1.1.7. new与make的区别
	//1.二者都是用来做内存分配的。
	//2.make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
	//3.而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。

}

// map
//map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用。
func mapBasics() {
	//1.1.1. map定义
	//map[KeyType]ValueType    KeyType:表示键的类型。  ValueType:表示键对应的值的类型
	//map类型的变量默认初始值为nil，需要使用make()函数来分配内存
	//make(map[KeyType]ValueType, [cap])        cap表示map的容量，非必须，但应该合适指定

	//1.1.2. map基本使用
	scoreMap := make(map[string]int, 8)
	scoreMap["huang"] = 11
	scoreMap["zzz"] = 22
	fmt.Println(scoreMap, scoreMap["huang"])
	// 在声明的时候初始化元素
	userInfo := map[string]string{
		"username": "huang",
		"password": "123456",
	}
	fmt.Println(userInfo)

	var userList = make([]map[string]interface{}, 3)
	userList[0] = map[string]interface{}{"username": "hh", "age": 22}
	userList[1] = map[string]interface{}{"username": "zz", "age": 22}

	userList2 := []map[string]interface{}{
		{"username": "hh", "age": 22},
		{"username": "zz", "age": 23},
	}
	fmt.Println(userList2)

	//1.1.3. 判断某个键是否存在
	//判断map中键是否存在的特殊写法，格式如下:
	//v, ok := userInfo["username"]   // 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	//if ok {
	//	fmt.Println(v)
	//}else {
	//	fmt.Println("没有这个key")
	//}

	//1.1.4. map的遍历
	//注意： 遍历map时的元素顺序与添加键值对的顺序无关
	//for k, v := range userInfo {   // 使用for range遍历map
	//	fmt.Println(k, v)
	//}
	//for k := range userInfo {   // 只遍历key
	//	fmt.Println(k)
	//}

	//1.1.5. 使用delete()函数删除键值对
	//使用delete()内建函数从map中删除一组键值对，delete()函数的格式如下：
	//delete(map, key)     map:表示要删除键值对的map   key:表示要删除的键值对的键
	//delete(scoreMap, "zzz")

	//1.1.6. 按照指定顺序遍历map
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子
	sMap := make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		sMap[key] = value
	}
	keys := make([]string, 0, 200) // 定义切片
	for k := range sMap {
		keys = append(keys, k) // 取出map中的所有key存入切片keys
	}
	sort.Strings(keys) //对切片进行排序
	for _, v := range keys {
		fmt.Println(v, sMap[v])
	}

	//1.1.7. 元素为map类型的切片
	mapSlice := make([]map[string]string, 5)
	fmt.Println("mapSlice----------", mapSlice)
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["username"] = "huang"
	mapSlice[0]["password"] = "123456"

	//1.1.8. 值为切片类型的map
	sliceMap := make(map[string][]string, 3)
	fmt.Println("sliceMap----------", sliceMap)
	sliceMap["userList"] = make([]string, 4, 10)
	sliceMap["userList"][0] = "huang"
	sliceMap["userList"][1] = "zzzz"
	sliceMap["score"] = append(sliceMap["score"], "99")
	fmt.Println("sliceMap------------", sliceMap)

}

// 指针数组
func pointArr() {
	// 指针数组(存指针的数组，数组元素为指针)
	const MAX int = 3
	b := []int{10, 100, 200}
	var i int
	var ptr [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr[i] = &b[i] /* 整数地址赋值给指针数组 */
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}
}

func main() {

	//arrBasics()
	//_slice()
	//pointBasics()
	mapBasics()
	//pointArr()

}

func swap(x *int, y *int) {
	var temp int
	temp = *x /* 保存 x 地址的值 */
	*x = *y   /* 将 y 赋值给 x */
	*y = temp /* 将 temp 赋值给 y */
}

func arrTest(x [3]int) [3]int {
	fmt.Printf("x address is %p\n", &x)
	x[1] = 111 // 修改时进行了一次深拷贝
	return x
}

func arrTest2(x *[3]int) {
	fmt.Printf("x address is %p\n", &x)
	x[0] = 99
}

func sumArr(arr []int) int {
	var numSum int
	for i := 0; i < len(arr); i++ {
		numSum += arr[i]
	}
	return numSum
}

func findIndex(arr []int, target int) [2]int {

	return [2]int{1, 2}
}
