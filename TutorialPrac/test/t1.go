package main

import "fmt"
// 这是一个闭包
func getSeq() func() int{
	i := 0
	return func() int {  // 匿名函数
		// 内部使用外部的变量
		i += 1
		return i
	}
}

func main() {
	// 使用闭包
	nextNum := getSeq()  // 这里的返回值是一个函数
	fmt.Println(nextNum())
	fmt.Println(nextNum())

	// 新的闭包
	nextNum2 := getSeq()
	fmt.Println(nextNum2())
	fmt.Println(nextNum2())
	fmt.Println(nextNum2())
}