package main

import (
	"fmt"
	"strconv"
	"time"
)

func convert() {

	// 数字转换为str
	i := 123
	iStr2 := fmt.Sprintf("%d", i)
	h := fmt.Sprint(i)
	iStr := strconv.Itoa(i)
	fmt.Println(h)
	fmt.Println(iStr2, iStr)

	// str转换为数字
	fmt.Println(strconv.Atoi(h))
}


// 常量只能是基础类型
//type Temporature struct {
//	name string
//}
//
//const t = Temporature{name: "ok"}


func constantVal() {
	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay)     // "time.Duration 0"
	fmt.Printf("%T %[1]v\n", timeout)     // "time.Duration 5m0s"
	fmt.Printf("%T %[1]v\n", time.Minute) // "time.Duration 1m0s"

	// 这个很重要，如果右边的忽略了，就按照上面的来
	const (
		a = 1
		b
		c = 2
		d
	)

	fmt.Println(a, b, c, d) // "1 1 2 2"

	// iota
	type Weekday int

	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Sunday, Monday, Friday)

}