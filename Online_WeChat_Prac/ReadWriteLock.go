package main

import (
	"log"
	"math/rand"
	"sync"
)

var (
	val int
	lock sync.RWMutex
	wg sync.WaitGroup
)

func run10() {
	// 读写锁
	// 使用10个协程，分别5个去写5个去读
	wg.Add(10)
	for i := 0; i < 5; i++ {
		go read(i)
	}

	for i := 0; i < 5; i++ {
		go write(i)
	}
	wg.Wait()
}

func read(i int)  {
	lock.RLock()
	log.Printf("读goroutine %d 正在读取 \n", i)
	v := val
	log.Printf("读goroutine %d 读取到的数据为：%d \n", i, v)
	wg.Done()
	lock.RUnlock()
}

func write(i int) {
	lock.Lock()
	log.Printf("写goroutine %d 正在写入 \n", i)
	v := rand.Intn(1000)
	val = v
	log.Printf("写goroutine %d 写入的数据为：%d \n", i, v)
	wg.Done()
	lock.Unlock()
}