package main

import (
	"GoPractice/TheGoProgramming/ch8/reverb"
	"flag"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	//fib.Cal()

	// 开启网络服务  等待链接
	//networkService.Clock()

	//reverb.Pipeline()
	//reverb.Pipeline2()

	//reverb.RocketRun()
	//reverb.RocketRun2()
	//reverb.Select2()
	//reverb.RunSelect()
	reverb.RunChat()
}