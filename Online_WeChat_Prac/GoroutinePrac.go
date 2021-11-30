package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func exp1() {
	//runtime.GOMAXPROCS(1)
	var w sync.WaitGroup
	w.Add(2) // 计数的信号量
	go func() {
		defer w.Done()
		for i := 0; i < 100; i++ {
			fmt.Println("A")
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		defer w.Done()
		for i := 0; i < 100; i++ {
			fmt.Println("B")
			time.Sleep(100 * time.Millisecond)

		}
	}()

	w.Wait() // 一直等待2个goroutine完成后，再结束。
}

var (
	count int32
	w     sync.WaitGroup
	m sync.Mutex
)

func exp2() {
	// 并发处理临界内存
	w.Add(2)
	go countInt()
	go countInt()
	w.Wait()
	fmt.Println(count)
}

func countInt() {
	defer w.Done()
	for i := 0; i < 2; i++ {
		// 使用atomic包来解决并发问题
		/**
		time.Sleep(100 * time.Millisecond)
		//val := count
		val := atomic.LoadInt32(&count)
		runtime.Gosched() // yields the processor
		val++
		//count = val
		atomic.StoreInt32(&count, val)
		 */

		// 使用锁
		m.Lock()
		time.Sleep(100 * time.Millisecond)
		val := count
		runtime.Gosched() // yields the processor
		val++
		count = val
		m.Unlock()
	}
}

func exp3() {
	// 无缓冲通道
	ch := make(chan int)
	go func() {
		sum := 0
		for i := 0; i < 10; i++ {
			sum += i
		}
		ch <- sum  // 在计算sum和的goroutine没有执行完，把值赋给ch通道之前，fmt.Println(<-ch)会一直等待。所以main主goroutine就不会终止
	}()
	fmt.Println(<-ch)  // 由于是无缓冲通道，需要阻塞到写入都写完
}

func exp4() {
	// 类比Linux中 | 的作用
	c1 := make(chan int)
	c2 := make(chan int)

	// 在c1里面写入，结果写到c2中
	go func() {
		c1 <- 100
		fmt.Println("通道c1写完了")
	}()

	go func() {
		var val int
		val = <- c1
		c2 <- val
		fmt.Println("通道c2写完了")
	}()

	// 阻塞等到写完
	fmt.Println(<-c2)
}

func run9() {
	//exp1()
	//exp2()
	//exp3()
	exp4()



}
