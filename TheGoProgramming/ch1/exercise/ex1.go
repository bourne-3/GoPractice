package main

import (
	"fmt"
	"os"
)

func ex1() {
	// 同样输出 os.Args[0]
	var s, sep string
	for _, c := range os.Args{
		s += sep + c
		sep = " "
	}
	fmt.Println(s)
}

func ex2() {
	// idx与参数成对出现  one per line.
	for idx, c := range os.Args[1:]{
		fmt.Printf("%v : %v \n", idx, c)
	}
}

func ex3() {
	//TODO 验证join和拼接方式在runtime下的效率差异

}
