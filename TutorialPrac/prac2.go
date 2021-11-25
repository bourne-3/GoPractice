package main

import (
	"fmt"
	"math"
)

func isEqual(num int) bool {
	if num == 15 {
		fmt.Println("if语句true")
		return true
	}
	return false
}

func pow(x , n , targetNum float64) float64 {
	// if可以先初始化一下然后再进行判断
	if k := math.Pow(x, n); k > targetNum{
		return k
	}
	return targetNum
}


func h1() {
	//for i:=1; i < 10; i++ {
	//	fmt.Println("The num is:", i)
	//}
	k := 15
	fmt.Println(isEqual(k))

	res := pow(2, 3, 7)
	fmt.Println(res)
}



var c, python, java bool = true , false, true

func prac2() {
	var i int
	// implicit type
	k := 12  // var + 初始化

	fmt.Println(i, k)
	fmt.Println(c, python, java)
	fmt.Printf(" k:%T and %v ", k, k)
	h1()
}
