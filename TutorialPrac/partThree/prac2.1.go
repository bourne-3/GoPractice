package main

import (
	"fmt"
	"sync"
	"time"
)

// 解决脏读问题
type SafeCounter struct {
	mu sync.Mutex
	v map[string]int
}

// 增加
func (counter *SafeCounter) add(key string) {
	counter.mu.Lock()  // 加锁
	defer counter.mu.Unlock()
	counter.v[key]++
	//counter.mu.Unlock()
}

// 获取
func (counter *SafeCounter) getVal(key string) int{
	counter.mu.Lock()
	defer counter.mu.Unlock()
	return counter.v[key]
}

func run2_1()  {
	counter := SafeCounter{v: make(map[string]int)}
	// 多线程增加
	for i := 0; i < 1000; i++ {
		go counter.add("keyToUse")
	}
	time.Sleep(time.Second )
	fmt.Println(counter.getVal("keyToUse"))
}