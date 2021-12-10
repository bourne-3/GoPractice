package reverb

import (
	"fmt"
	"os"
	"time"
)

func RocketRun() {
	fmt.Println("倒计时")
	tick := time.Tick(1 * time.Second)  // 返回只可接收的channel



	for count := 10; count > 0; count-- {
		fmt.Println(count)
		<- tick


	}
	fmt.Println("===============Launch==================")
}

func RocketRun2() {
	// 添加按键结束
	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	tick := time.Tick(10 * time.Second)  // 返回只可接收的channel
	fmt.Println("开始倒计时，按下enter也可以退出")
	select {
	case <-tick:
		// Do nothing
	case <-abort:
		fmt.Println("按下空格 退出")
		return
	}
	fmt.Println("===============Launch==================")
}

func Select2() {
	fmt.Println("Commencing countdown.  Press return to abort.")
	tick := time.Tick(1 * time.Second)
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	for c := 10; c > 0; c-- {
		fmt.Println(c)
		select {
		case <-tick:
			// DO NOTHING
		case <-abort:
			fmt.Println("Launch aborted!")
			return
			//default:
			//	fmt.Println("实现轮询")
			//}
		}
		fmt.Println("===============Launch==================")
	}
}
