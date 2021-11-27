package main

import "fmt"





//==
type I interface {
	M()
}

type T struct {
	member string
}

// T去是实现I
func (t *T) M() {
	// 如果传进来的是初始化的
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.member)
}

// 接口
type Abser interface {
	Abs() float64
}
type MyFloat2 float64

type Vertex2 struct {
	x, y int
}

// 实现接口
func (myFloat MyFloat2) Abs() float64 {
	return float64(myFloat + 10)
}

// 再实现一个接口
func (v *Vertex2) Abs() float64 {
	return float64(v.x + 5 + v.y)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func run2() {
	var i Abser

	// 第一个实现
	n := MyFloat2(15)
	i = n
	fmt.Println(i.Abs())

	// 第二个实现
	v := Vertex2{1 , 1}
	i = &v  // 注意这里需要 &v
	fmt.Println(i.Abs())

	// 第三个 使用T去实现I
	var i2 I  // 这是一个接口
	t := T{"小白"}  // 这是一个实现了该接口的方法
	i2 = &t  // 接口可以接受实现了该接口下方法的值
	describe(i2)
	i2.M()  // 有点像多态的调用

	// =======
	// 第四个 接口接收了初始化的val
	var t2 *T
	i2 = t2
	describe(i2)
	i2.M()
}
