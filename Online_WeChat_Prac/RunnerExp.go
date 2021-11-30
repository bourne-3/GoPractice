package main

import (
	"os"
	"time"
)

type Runner struct {
	tasks []func(int)
	complete chan error
	timeout <-chan time.Time // 只能接收的单向通道
	interrupt chan os.Signal
}

// 工厂
func New(tm time.Duration) *Runner {
	return &Runner{
		complete: make(chan error),
		timeout: time.After(tm),
		interrupt: make(chan os.Signal, 1),  // 缓冲通道
	}
}

