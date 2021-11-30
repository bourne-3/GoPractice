package main

import (
	"fmt"
	"strconv"
)

type person struct {
	id int64
	name string
}

func run5() {
	// function
	sum := addTwoSum(1, 2)
	fmt.Println(sum)

	// method  (with a receiver)
	p := person{id:58, name: "bourne"}
	fmt.Println(p.getString())

	// point receiver
	p.pointModify()
	fmt.Println(p)

	// 多值返回
	subNum, e := subN(10,5)
	fmt.Println(subNum, e)

	// 可变参数
	printArgs(10, 12, 14)
}

// 使用...接受多个参数，类型是 interface{}
func printArgs(a ...interface{}) {
	fmt.Println(a)
}


func subN(a int, b int) (int, error) {
	return a - b, nil
}

func (p *person) pointModify() {
	p.name = "小龙"
}

// 值接收者
func (p person) getString() string {
	return "The person name is " + p.name + " and his age is " + strconv.FormatInt(p.id,10)   //+ " and his age is " + string(p.id) 这里好像有问题 在转换的时候
}

func addTwoSum(a int64, b int64) int64 {  //这里也可以返回一个空接口 interface{}
	return a + b
}