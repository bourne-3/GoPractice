package main

import "fmt"

func f() *int {
	v := 1;
	return &v
}


func main() {
	p := f()  // 局部变量的地址
	fmt.Println(*p)
}
