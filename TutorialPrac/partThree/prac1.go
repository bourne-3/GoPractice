package main

import (
	"fmt"
	"time"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x + y
	}
	close(c)
}

func tryBuffer() {
	ch := make(chan int, 2)
	ch <- 10
	ch <- 21
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}

func tryChannel(arr []int, c chan int) {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	c <- sum
}


func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func run() {
	//go say("bourne is handsome")
	//say("hello")

	//============
	// 创建channel
	//c := make(chan int)
	//s := []int{7, 2, 8, -9, 4, 0}
	//// 使用两个协程完成
	//go tryChannel(s[:len(s) / 2], c)
	//go tryChannel(s[len(s) / 2 :], c)
	//n1, n2 := <-c, <-c
	//fmt.Println(n1, n2)

	//==============
	//tryBuffer()

	//==============
	//c := make(chan int, 10)  // 使用管道去里面写
	//fibonacci(10, c)
	//for i := range c{
	//	fmt.Println(i)
	//}

	//================== select
	//c := make(chan int)
	//quit := make(chan int)
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		fmt.Println(<-c)
	//	}
	//	quit <- 0
	//}()  // 类似于匿名函数，在这里就直接定义一个函数，后面的括号对应前面的括号
	//fibonacci2(c, quit)

	//===================default
	withDefault()
}

func withDefault() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("Boom!")
			return
		default:
			fmt.Println("....")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func fibonacci2(c chan int, quit chan int) {
	x, y := 0, 1
	for  {
		select {
		case c <- x:
			x, y = y, x + y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}

}