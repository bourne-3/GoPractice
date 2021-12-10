package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// 测试readALl
	t3()
}

// 测试os.Stdin.Read(make([]byte, 1)) // read a single byte
func t3() {
	c, _ := os.Stdin.Read(make([]byte, 1)) // read a single byte
	fmt.Println(c)
}


func t2() {
	// 检验Buffer的write
	var b bytes.Buffer
	b.Write([]byte("hello world"))
	fmt.Fprintf(&b, " 再加点东西")
	//fmt.Fprint(&b,",","http://www.flysnow.org")
	//b.WriteTo(os.Stdout)

	fmt.Println("==========")
	// 检验Buffer的Read    read 是读到切片中，写是将切片的内容写出去

	// 方法1
	//var p [100]byte
	//p := []byte("content read by go")
	//n,err:=b.Read(p[:])
	//fmt.Println(n,err,string(p[:n]))

	// 方法2
	data,err:=ioutil.ReadAll(&b)
	fmt.Println(string(data),err)
}

func t1() {
	reader := bufio.NewReader(os.Stdin)
	b, _, _ := reader.ReadLine()
	fmt.Printf(string(b))
}