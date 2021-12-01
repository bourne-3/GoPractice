package main

import (
	"log"
)



func init() {
	log.SetPrefix("【LogFile】")  // 标志是哪一些业务
	log.SetFlags(log.Ldate|log.Ltime|log.Lshortfile)
}

func run11() {
	log.Println("飞雪无情的博客:\",\"http://www.flysnow.org")
	log.Printf("www.hupu.com")
}