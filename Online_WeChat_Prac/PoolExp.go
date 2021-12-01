package main

import (
	"errors"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)
var ErrPollClosed = errors.New("资源池已经关闭")

type Pool struct {
	m sync.Mutex
	res chan io.Closer
	factory func() (io.Closer, error)
	closed bool
}

// NewPool 工厂函数
func NewPool(fn func() (io.Closer, error), size uint) (*Pool, error)  {
	if size <0 {
		 return nil, errors.New("给定的size太小了")
	}
	return &Pool{
		res: make(chan io.Closer, size),
		factory: fn,
	}, nil
}

// 获取资源
func (p *Pool) accquire() (io.Closer, error){
	select {
	case r, ok := <-p.res:  // 通道的多参返回 （判断通道是否已关闭
		log.Println("Acquire:共享资源")
		if !ok {
			return nil, ErrPollClosed
		}
		return r, nil
	default:
		log.Println("Acquire:新生成资源")
		return p.factory()
	}
}

// 关闭资源  p的方法close
func (p *Pool) close() {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}

	p.closed = true
	close(p.res)   // 内置函数close  关闭通道
	for r := range p.res{
		r.Close()  // 接口 Close
	}
}

func (p *Pool) release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		r.Close()
		return
	}

	select {
	case p.res <- r:
		log.Println("资源释放到池子里了")
	default:
		log.Println("资源池满了，将这个资源释放")
		r.Close()
	}
}

const(
	maxGoroutine = 5
	poolRes = 2
)

// 主要的main运行文件这里
func runExpPool() {
	var w sync.WaitGroup
	w.Add(maxGoroutine)

	p, err := NewPool(createConnection, poolRes)
	if err != nil {
		log.Println(err)
		return
	}

	// 模拟多个协程
	for query := 0; query < maxGoroutine; query++ {
		go func(q int) {
			dbQuery(q, p)
			w.Done()
		}(query)
	}

	w.Wait()
	log.Println("开始关闭资源池")
	p.close()
}

func dbQuery(q int, p *Pool) {
	conn, err := p.accquire()  // conn io.Close
	if err != nil{
		log.Println(err)
		return
	}
	defer p.release(conn)

	//模拟查询
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("第%d个查询，使用的是ID为%d的数据库连接", q, conn.(*dbConnection).ID)
}

var idCounter int32

type dbConnection struct {
	ID int32
}

// Close 实现io.Closer的接口
func (db *dbConnection) Close() error {
	log.Println("关闭连接", db.ID)
	return nil
}

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	return &dbConnection{id},nil
}



