package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 测试readALl
	reader := bufio.NewReader(os.Stdin)
	b, _, _ := reader.ReadLine()
	fmt.Printf(string(b))
}
