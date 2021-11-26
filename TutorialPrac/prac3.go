package main

import "fmt"

// structs
type Vertex struct {
	X int
	y int
}

func tryStructs() {
	fmt.Println(Vertex{1, 2})
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