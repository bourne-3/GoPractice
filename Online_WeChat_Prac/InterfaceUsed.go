package main

import (
	"bytes"
	"fmt"
	"io"
)

func run6() {
	var b bytes.Buffer
	fmt.Fprintf(&b, "Hello world")  // 将后面的内容写到&b中 // 看源码，这里是将&b传递给io.Writer这个接口
	fmt.Println(b.String())

	var w io.Writer  // 这是一个接口
	w = &b
	fmt.Println(w)  // 可以把用户定义的类型称之为实体类型  实际上调用的是 实体类型

	// 多态的使用
	var a animal
	a = Cat(1)
	a.printInfo()
	a = Dog(2)
	a.printInfo()

	// receiver是用值还是用指针的
	var c2 = Cat(2)
	invoke(c2)
	invoke(&c2)  // 两者都可以传递给animal接口

	var d1 = Duck(1)
	invoke(&d1)
	//invoke(d1)  // 这样就不行了，必须使用%d1传递

}

func invoke(a animal) {
	a.printInfo()
}

type animal interface {
	printInfo()
}
type Cat int
type Dog int
type Duck int

// 使用值接收者实现接口
func (d *Duck) printInfo() {
	fmt.Println("This is a duck. Implement by point receiver")
}

// 是方法实现接口
func (c Cat) printInfo() {
	fmt.Println("这是一只猫 cat")
}

func (d Dog) printInfo()  {
	fmt.Println("这是一直狗 dog")
}