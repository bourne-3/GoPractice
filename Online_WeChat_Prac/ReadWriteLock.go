package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
)

var (
	val int
	lock sync.RWMutex
	wg sync.WaitGroup
)


// 自己写一个SynchronizedMap
type synchronizedMap struct {
	m *sync.RWMutex
	data map[interface{}]interface{}
}

// CRUD and Each
// add
func (s *synchronizedMap) add(k interface{}, v interface{}) {
	s.m.Lock()
	defer s.m.Unlock()
	s.data[k] = v
}

// delete
func (s *synchronizedMap) delete(k interface{}) {
	s.m.Lock()
	defer s.m.Unlock()
	delete(s.data, k)
}

// modify
func (s *synchronizedMap) modify(k interface{}, v interface{}) {
	s.m.Lock()
	defer s.m.Unlock()
	s.data[k] = v
}

// get
func (s *synchronizedMap) get(k interface{})  interface{}{
	s.m.RLock()
	defer s.m.RUnlock()
	v := s.data[k]
	return v
}

// each
func (s *synchronizedMap) each(fn func(interface{}, interface{}))  {
	s.m.RLock()
	defer s.m.RUnlock()
	for k, v := range s.data {
		fn(k, v)
	}
}

// 初始化一个synchronizedMap
func newSynchronizedMap() *synchronizedMap {
	return &synchronizedMap{
		m:new(sync.RWMutex),
		data:make(map[interface{}]interface{}),
	}
}

func test(data map[interface{}]interface{}) {
	fmt.Println(data)
}

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

	//==============
	//实验synchronized
	syMap := newSynchronizedMap()
	syMap.add(10, 20)
	// 这里使用了回调函数  重要
	syMap.each(func(k interface{}, v interface{}) {
		fmt.Println(k, " is ", v)
	})
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