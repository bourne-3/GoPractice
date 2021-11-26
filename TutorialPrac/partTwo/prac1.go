package main

import "fmt"

type Vertex struct {
	x, y int
}

type MyFloat float64

// (v Vertex) 被称之为 receiver
func (v Vertex) tryMethod() float64{
	return float64(v.x + v.y)
}



// pointer receiver
func (v *Vertex) scale()  {
	// 动态变化这个v里面的值
	v.x += 10
	v.y += 10
}

//
func (customF MyFloat) tryMethod() float64 {
	return float64(customF + 10)
}



func Method2(v Vertex) float64 {
	return float64(v.x + v.y)
}

func run1() {
	v := Vertex{1, 2}
	fmt.Println(v.tryMethod())  // 类似与对象调用具体的函数  obj.method

	// Remember: a method is just a function with a receiver argument.
	v2 := Vertex{3, 4}
	fmt.Println(Method2(v2))

	// 不仅可以使用struct
	in := MyFloat(21)
	val_in := in.tryMethod()
	fmt.Println(val_in)

	fmt.Println("========")
	// ======== pointer reciver
	v3 := Vertex{5, 5}
	v3.scale()  // 各自再增加10
	val_v3 := v3.tryMethod()
	fmt.Println(val_v3)
}