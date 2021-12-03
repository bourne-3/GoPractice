package main

import (
	"context"
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


// 使用context  （控制单个goroutine
func way3() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for  {
			select {
			case <- ctx.Done():
				fmt.Println("监控停止，退出！")
				return
			default:
				fmt.Println("goroutine正在监控中")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}(ctx)
	time.Sleep(3 * time.Second) // 让业务监控个几秒
	fmt.Println("可以了，我要通知监控停止了")
	cancel()
	// 检测是否真正停止了，看看还会不会输出消息
	time.Sleep(2*time.Second)
}

// 控制多个goroutine
func way4() {
	ctx, cancel := context.WithCancel(context.Background())
	// 创建三个协程
	go watch(ctx, 1)
	go watch(ctx, 2)
	go watch(ctx, 3)

	time.Sleep(3 * time.Second)
	fmt.Println("可以了，监控停止，停止多个协程的工作")
	cancel()
	time.Sleep(2 * time.Second)  // 检测是否还有在跑
}

func watch(ctx context.Context, num int) {
	for {
		select {
		case <- ctx.Done():
			fmt.Printf("监控结束，协程 %v 结束\n", num)
			//fmt.Println(ctx.Value("name"))
			return  // 这里要return
		default:
			fmt.Printf("协程 %v 正在监控着程序 \n", num)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func tryWithVal() {
	key := "name"
	ctx, cancel := context.WithCancel(context.Background())
	valCtx := context.WithValue(ctx, key,"【main监控】")
	go watch(valCtx, 10)
	time.Sleep(3 * time.Second)
	fmt.Println("停止监控")
	cancel()
	time.Sleep(3 * time.Second)

}


func run13() {
	//wGroup()
	//way2()
	//way3()
	//way4()
	tryWithVal()
}
