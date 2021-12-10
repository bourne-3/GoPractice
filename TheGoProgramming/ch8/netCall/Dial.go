package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	// 主动去连接  这里模拟了一下telnet
	//conn, err := net.Dial("tcp", "localhost:8000")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer conn.Close()
	//mustCopy(os.Stdout, conn)

	//connetEcho()
	ChannelEcho()
}

func mustCopy(dst io.Writer, src io.Reader) {
	// 不断把内容写到 stdout
	if _, err := io.Copy(dst, src); err != nil{
		log.Fatal(err)
	}
}

func ChannelEcho() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}  // 仅仅是一个事件
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<- done  // 同步
}

func connetEcho() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)  // 参数需要用接口接收才是更好的
}
