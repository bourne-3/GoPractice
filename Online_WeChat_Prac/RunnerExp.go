package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"time"
)
var (
	ErrTimeOut = errors.New("执行者执行超时")
	ErrInterrupt = errors.New("执行者中断")
)

type Runner struct {
	tasks []func(int)
	complete chan error
	timeout <-chan time.Time // 只能接收的单向通道
	interrupt chan os.Signal
}

// New 工厂
func New(tm time.Duration) *Runner {
	return &Runner{
		complete: make(chan error),
		timeout: time.After(tm),
		interrupt: make(chan os.Signal, 1),  // 缓冲通道
	}
}

// Add 方法 添加tasks
func (r *Runner) Add(taks... func(int)) {
	r.tasks = append(r.tasks, taks...)  // 这里的...表示0个或者多个
}

// 方法 执行任务
func (r *Runner) run() error{
	// 取出任务
	for id,task := range r.tasks{
		if r.isInterrupt(){
			return ErrInterrupt
		}
		task(id)  // 执行该函数
	}
	return nil
}

// 判断是否有中断信号
func (r *Runner) isInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}


// Start 开始运行
func (r *Runner) Start() error{
	signal.Notify(r.interrupt, os.Interrupt)
	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeOut
	}
}

func expRun() {
	log.Println("..开始执行任务..")
	timeout := 4 * time.Second
	r := New(timeout)

	// 添加任务
	r.Add(createTask(), createTask(), createTask())

	if err := r.Start(); err != nil{
		switch err {
		case ErrTimeOut:
			log.Println(err)
			os.Exit(1)
		case ErrInterrupt:
			log.Println(err)
			os.Exit(2)
		}
	}

	log.Println("...任务执行结束")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("正在执行任务%d", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
