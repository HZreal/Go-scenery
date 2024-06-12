package main

import (
	"fmt"
	"reflect"
)

func test1() {
	var num int64 = 100

	// reflect.Type.TypeOf()
	numT := reflect.TypeOf(num)
	fmt.Println(numT)

	// reflect.Type.ValueOf()
	numV := reflect.ValueOf(num)
	fmt.Println(numV)

	// reflect.Type.Kind()
	var i interface{} = 30.2
	iT := reflect.TypeOf(i)
	iV := reflect.ValueOf(i)
	iVT := iV.Type()
	fmt.Println(iT, iV, iVT)

}

type wrapInt int

func test2() {
	var num1 int = 10
	var num2 wrapInt = 20

	fmt.Println(reflect.TypeOf(num1).String())
	fmt.Println(reflect.TypeOf(num2).String())

	fmt.Println(reflect.Kind(num1))
	fmt.Println(reflect.Kind(num2))

}

func sliceReflect() {
	_slice := []interface{}{1, "a", true, 2.4}
	t := reflect.TypeOf(_slice)
	fmt.Printf("Slice Type: %s\nElem Type: %s\n", t, t.Elem())

	v := reflect.ValueOf(_slice)
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)
		realElem := reflect.ValueOf(elem.Interface())
		fmt.Printf("Index: %d, Elem Type: %v, Elem Kind: %v, Elem=%v\n", i, elem.Type(), realElem.Kind(), elem.Interface())
	}
}

func mapReflect() {
	_map := map[interface{}]interface{}{"name": "Bob", "age": 20, 1: "abc", 2.2: true}

	t := reflect.TypeOf(_map)
	fmt.Printf("Map Type: %s\nKey Type: %s\nValue Type: %s\n", t, t.Key(), t.Elem())

	v := reflect.ValueOf(_map)
	// fmt.Println(v)

	// v.MapKeys() 为 reflect.Value 的切片，i 为下标索引，key 为原 Map 的 key
	for i, key := range v.MapKeys() {
		fmt.Printf("%d: %v\n", i, key)
		realKey := reflect.ValueOf(key.Interface())

		value := v.MapIndex(key)
		valueField := reflect.ValueOf(value.Interface())
		fmt.Printf("Key Type: %s; Key Kind: %s; Key=%v; Value Type: %s; Value Kind: %s; Value=%v\n", key.Type(), realKey.Kind(), key.Interface(), value.Type(), valueField.Kind(), value.Interface())
	}
}

type Student struct {
	Name  string
	Age   int
	Score float64
	IsBoy bool
	I     interface{}
}

func structReflect() {
	st := Student{Name: "Alice", Age: 21, Score: 88.5, IsBoy: false, I: 50}
	t := reflect.TypeOf(st)
	v := reflect.ValueOf(st)

	fmt.Println("reflect.Type.Name:", t.Name())
	fmt.Println("reflect.Type.String:", t.String())

	fmt.Println("0 ---->  ", v.Field(0).Type(), v.Field(0).String())
	fmt.Println("1 ---->  ", v.Field(1).Type(), v.Field(1).Int())
	fmt.Println("2 ---->  ", v.Field(2).Type(), v.Field(2).Float())
	fmt.Println("3 ---->  ", v.Field(3).Type(), v.Field(3).Bool())
	fmt.Println("4 ---->  ", v.Field(4).Type(), v.Field(4).Interface(), reflect.ValueOf(v.Field(4).Interface()).Kind())

	for i := 0; i < t.NumField(); i++ {
		structField := t.Field(i)
		value := v.Field(i)
		fmt.Printf("%s: %v = %v\n", structField.Name, structField.Type.Name(), value.Interface())
	}
}
func main() {
	// test1()
	// test2()
	// sliceReflect()
	// mapReflect()
	// structReflect()
}
