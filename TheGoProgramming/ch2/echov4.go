package main

import (
	"flag"
	"fmt"
	"strings"
)

func echo() {
	var n = flag.Bool("n", false, "omit trailing newline")  // 忽略尾部的换行
	var sep = flag.String("s", " ", "separator")  // 分隔符
	flag.Parse()  // 需要解析 不然全部都是默认值
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()  // default下是换行的，将n设为true就不会换行了
	}
}

func h2() {
	// new函数
	val1 := new(int)  // 无名字 初始化零值 返回地址 三个要素
	//val1在这里是一个地址
	fmt.Println(val1)
	fmt.Println(*val1)
	*val1 = 10
	fmt.Println(*val1)
}

// tempConv0
type Celsius float64
type Fahrenheit float64
const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)

func C2F(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func F2C(f Fahrenheit)  Celsius{
	return Celsius((f - 32) * 5 / 9)
}

var a = b + c
var b = f1()
var c = 1

func f1() int{
	return c + 1
}


func pkgInit() {
	fmt.Println(a , b, c)
}

func runCompare() {
	// 运算操作
	fmt.Printf("%g\n", BoilingC - FreezingC)
	boilingF := C2F(BoilingC)  // 转换
	fmt.Printf("%g\n", boilingF - C2F(FreezingC))
	//fmt.Printf("%g\n", boilingF - FreezingC)  // error 不同的underlying

	// 比较操作
	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0)
	fmt.Println(f >= 0)

	//fmt.Println(c == f)  // error 不同的underlying
	fmt.Println(c == Celsius(f))  // 这是类型转换，不是函数调用
}

func runHelper() {
	//echo()
	//h2()

	// way2
	runCompare()


}