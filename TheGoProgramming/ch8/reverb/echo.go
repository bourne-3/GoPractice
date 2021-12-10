package reverb

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func HandleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1 * time.Second)  // 这里再进行并发 才可以保证是不阻塞的 交替出现
	}
	c.Close()
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func Pipeline2() {
	// 单向通道
	naturals := make(chan int)
	squares := make(chan int)

	go generate(naturals)
	go square(squares, naturals)
	printer(squares)

}
func generate(out chan<- int) {
	for x := 0; x < 10; x++{
		out <- x  // 注意 这里是给一个通道发送消息，因此不需要用 =
	}
	close(out)
}

func square(out chan<- int, in <-chan int) {
	for n := range in {
		out <- n * n
	}
	close(out)
}

func printer(in <-chan int)  {
	for v := range in{
		fmt.Println(v)
	}

}

func Pipeline() {
	naturals := make(chan int)
	squares := make(chan int)
	// 三个部分
	go func() {
		for x := 0; x < 10; x++{
			naturals <- x  // 注意 这里是给一个通道发送消息，因此不需要用 =
		}
		close(naturals)
	}()

	// 平方
	go func() {
		// 用for range来遍历一个channel
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// 输出
	for x := range squares{
		fmt.Println(x)
	}

}