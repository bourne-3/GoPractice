package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"time"
)

type Runner struct {
	tasks []func(int)
	timeout <-chan time.Time
	complete chan error
	interrupt chan os.Signal
}

// 创建一个Runner
func new(t time.Duration) *Runner {
	return &Runner{
		complete: make(chan error),
		interrupt: make(chan os.Signal, 1),
		timeout: time.After(t),
	}
}

// 添加任务
func (r *Runner) addFunc(fn... func(int)) {
	// tasks是一个slice
	r.tasks = append(r.tasks, fn...)
}

// 自定义错误
var (
	ErrTimeOut = errors.New("ErrTimeOut..")
	ErrOsInterrupt = errors.New("ErrOsInterrupt...")
)

// 运行 run
func (r *Runner) run() error{
	for id, task := range r.tasks{
		// 判断是否系统的interrupt
		if r.isInterrupt() {
			return ErrOsInterrupt
		}
		task(id)
	}
	return nil
}

func (r *Runner) isInterrupt() bool{
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}

// 开始 start。并且监控所有事件
func (r *Runner) start() error{
	signal.Notify(r.interrupt,os.Interrupt)
	go func() {
		r.complete <- r.run()  // 管道写入
	}()
	
	// 监控
	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeOut
	}
}

func RunnerExcute() {
	log.Println("开始执行任务")
	// 创建runner
	r := new(1 * time.Second)

	// 添加任务
	r.addFunc(createTask(), createTask(), createTask())

	// 开始执行并且监控
	if err := r.start(); err != nil {
		switch err {
		case ErrTimeOut:
			log.Println("出现了超时异常", err)
			os.Exit(1)
		case ErrOsInterrupt:
			log.Println("出现的系统异常 ", err)
			os.Exit(2)
		}
	}

	log.Println("任务执行结束Ending")
}

func createTask() func(int){
	return func(id int) {
		log.Println("正在执行任务 " , id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}