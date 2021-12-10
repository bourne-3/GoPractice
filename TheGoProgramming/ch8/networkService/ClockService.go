package networkService

import (
	"GoPractice/TheGoProgramming/ch8/reverb"
	"io"
	"log"
	"net"
	"time"
)





func Clock() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for{
		// 会一直阻塞到有人连接
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		// 处理连接过来的conn
		//go handleConn(conn)
		go reverb.HandleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()  // 保证网络关闭
	for{
		// conn有实现了write接口
		_, err := io.WriteString(conn, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}

}