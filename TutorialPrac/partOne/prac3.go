package main

import "fmt"

// structs
type Vertex struct {
	X int
	y int
}

func tryArray() {
	var a [2]int  // 先数量再类型
	a[0] = 21
	a[1] = 22
	fmt.Println(a)

	s1  := [3]string{"bourne", "lucy", "kim"}
	fmt.Println(s1)

	// 使用slice
	slice1 := s1[:2]  // [low:hight] default 0 or len()
	fmt.Println(slice1)

	// 重要：Slices are like references to arrays
}

func tryStructs() {
	fmt.Println(Vertex{1, 2})
	// 获得struct的内容
	v2 := Vertex{12, 23}
	fmt.Println(v2.X)

	// 指针结合struct
	v3 := Vertex{13, 32}
	p := &v3
	p.X = 31  // 使用指针修改struct的内容  without the explicit dereference.
	fmt.Println(*p, v3, p)


}


// 指针
func pointer1() {
	var p *int  // 指针
	i := 15
	p = &i
	fmt.Println("这是一个地址", p)  // 这是一个地址
	fmt.Println("具体的值", *p)  // 具体的值
	*p = 31
	fmt.Println("*p = 31之后：", *p)  // 具体的值
}