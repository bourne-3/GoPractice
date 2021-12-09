package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// write方法
type ByteCounter int

func (b *ByteCounter) Write(p []byte) (int, error){
	*b += ByteCounter(len(p))
	return len(p), nil
}
func run1()  {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	// reset
	c = 0
	var name = "bourne"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)

}
type IntSet struct {

}

func (i *IntSet) String() string{
	return ""
}

func m2() {
	var any interface{}
	any = true
	any = 12.34
	any = "hello"
	any = make(map[string]int)
	any = new(bytes.Buffer)
	fmt.Println(any)
}

func m3() {
	// 接口值
	var w io.Writer  // nil nil
	w = os.Stdout
	w.Write([]byte("hello"))
	fmt.Printf("\nw type:%T \n", w)

	w = new(bytes.Buffer)
	w.Write([]byte("hello"))
	fmt.Println(w)
	w = nil

}

func run2() {
	// 1
	var w io.Writer
	w = os.Stdout
	w = new(bytes.Buffer)
	//w = time.Second

	var rwc io.ReadWriteCloser
	rwc = os.Stdout
	//rwc = new(bytes.Buffer)  // 右边没有全部实现左边的接口

	w = rwc
	//rwc = w  // error
	fmt.Println(w)

	// 2
	var s IntSet
	s.String()  // 可以 会自动转换为 (*s).string()

	// 由于是*IntSet实现了String()，因此也只能是指针才能赋值给该接口
	var _ fmt.Stringer = &s // 是一个接口  ok
	//var _ fmt.Stringer = s //error

	// 3
	os.Stdout.Write([]byte("hello")) // OK: *os.File has Write method
	os.Stdout.Close()                // OK: *os.File has Close method

	var w2 io.Writer
	w2 = os.Stdout  // os.Stdout.Write 好像只能出现一次？
	w2.Write([]byte(" word"))
	//w2.Close()

}


func main() {
	//run1()
	//run2()
	m3()
}
