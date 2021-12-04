package main

import (
	"fmt"
	"os"
	"strings"
)

func echo() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " " // 用空格隔开
	}
	fmt.Println(s)

	// 2 i++是一个statements 而不是expression
	/**
	i := 1
	fmt.Println(i)
	i++  // 可以
	++i  // 不可以
	 */

}

// 使用了range
func echo2() {
	var s, sep string
	for _, c := range os.Args[1:]{  // 从1开始
		// 因为os.Args[0]是这个.go的全路径名字
		s += sep + c
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	// 使用join来避免字符串的拼接
	s := strings.Join(os.Args[1:], " ")
	fmt.Println(s)
}