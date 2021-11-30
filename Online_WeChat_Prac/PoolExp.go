package main

import (
	"errors"
	"io"
	"log"
	"sync"
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
	case r, ok := <-p.res:
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





