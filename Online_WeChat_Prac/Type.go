package main

import "fmt"
type orders struct {
	age int
	name string
}

func run4() {
	// 基本类型  基本都是传递副本
	name := "张三"
	fmt.Println(name)
	newName := modifyVal(name)
	fmt.Println(newName)
	fmt.Println(name)

	// 引用类型  传递的是个指向底层数据的指针
	person := map[string]int{"age": 12}
	modifyReferece(person)
	fmt.Println(person)

	// 结构类型


	var o1 orders
	fmt.Println(o1)

	o2 := orders{age: 19}
	fmt.Println(o2)

	// 结构体的引用类型
	modifyConstruct(o2)  // 没有变化
	fmt.Println(o2)
	modifyConstructP(&o2)  // 有变化
	fmt.Println(o2)

	// 自定义类型  （基于已有的类型创建现有的类型
	type Duration int64
	var d Duration  // way1
	d = 13
	fmt.Println(d)

	d2 := Duration(31)  // way2
	fmt.Println(d2)
}

func modifyConstructP(o *orders) {
	// 使用指针修改了
	o.age = 100

}

func modifyConstruct(o1 orders) {
	o1.age = 0
}

func modifyReferece(person map[string]int) {
	person["age"] = 100

}

func modifyVal(name string) string{
	name = name + " modified"
	return name
}
