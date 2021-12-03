package main

import (
	"fmt"
	"sync"
	"time"
)

// waitGroup
func wGroup() {
	var w sync.WaitGroup
	w.Add(2)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("1号goroutine完成")
		w.Done()
	}()

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("2号goroutine完成")
		w.Done()
	}()

	// 等待执行完毕
	w.Wait()
	fmt.Println("下班！")
}

// 通过向chan发送消息的方式
// 当数量小并且不存在依赖的情况下，这是一个不错的选择
func way2() {
	// 创建一个通道
	stop := make(chan int)

	// 协程
	go func() {
		for{
			select {
			case <- stop:
				// stop通道里面有信息了 （由于创建的是无缓冲通道，只有阻塞等待的feature
				fmt.Println("监控退出，停止了 （使用通道的方式）")
				return
			default:
				fmt.Println("goroutine正在监控中....")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("可以了，我要通知监控停止了")
	stop <- 1

	// 检测是否真正停止了，看看还会不会输出消息
	time.Sleep(2*time.Second)
}




func run13() {
	//wGroup()
	way2()
}
