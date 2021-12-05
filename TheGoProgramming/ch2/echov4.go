package main

import (
	"flag"
	"fmt"
	"math"
	"strings"
	"unicode/utf8"
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

func f2() {
	o := 0666
	fmt.Printf("%d %[1]o", o)

	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)   // "97 a 'a'"
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 国 '国'"
	fmt.Printf("%d %[1]q\n", newline)  // %q是作用？
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))  // 8.3。 8表示缩进 3表示保留小数位
	}
}

func f3String() {
	s := "Hello world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])  // 77 111

	s2 := `这是一个原生的字面量
还可以换行处理 。 \n等转义字符都不可以用了`
	fmt.Println(s2)

	s3 := "hello,世界"  // 8个字符 但是len()为12  因为包含了中文
	s4 := "hello,world"  // 11个字符
	fmt.Println(len(s3), len(s4))
	fmt.Println(utf8.RuneCountInString(s3))  // 这样就可以获得真正包含的字符数了
	fmt.Println("=============")

	for i := 0; i < len(s3); {
		r, size := utf8.DecodeRuneInString(s3[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	fmt.Println("===========")

	// "program" in Japanese katakana
	//s = "プログラム"
	s = "你好中国"
	fmt.Printf("% x\n", s) // "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0"
	r := []rune(s)
	fmt.Printf("%x\n", r)  // "[30d7 30ed 30b0 30e9 30e0]"
}

func pkgInit() {
	fmt.Println(a , b, c)
	//f2()
	f3String()
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

	fmt.Println("=============")
}

func runHelper() {
	//echo()
	//h2()

	// way2
	runCompare()


}