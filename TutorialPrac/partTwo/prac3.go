package main
// the rest

import (
	"fmt"
	"io"
	"strings"
	"time"
)



// io.Reader
func tryReader() {
	r := strings.NewReader("Hello,world")
	b := make([]byte, 4)
	fmt.Println(b)
	for  {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

}



// error接口
type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func getError() error{  // 接口作为返回值
	return &MyError{time.Now(), "Error happen"}
}

type Person struct {
	id int
	name string
}

// Stringer接口
func (p Person) String() string {
	return fmt.Sprintf("自定义输出, 我的id是 %v, name 是 %v", p.id, p.name)
}

func runRest() {
	//p := Person{id:18, name:"kobe"}
	//fmt.Println(p)
	if error := getError(); error != nil {
		fmt.Println(error)
	}

}

func do(i interface{})  {
	switch v := i.(type) {
	case int:
		fmt.Println("is int ", v)
	case string:
		fmt.Println("is string ", v)
	default:
		fmt.Println("may be is bool ", v)
	}


}

func trySwitchAssert() {
	do(15)
	do("hello")
	do(true)
}

func tryAssert() {
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)

	// 可以用两个值接受
	s2, flag := i.(string)
	fmt.Println(s2, flag)

	// 如果是不存在的
	s3, flag2 := i.(int)
	fmt.Println(s3, flag2)

	//s4 := i.(int) // 只能一次判断错误，第二次就不可以了
	//fmt.Println(s4)
}

func run3() {

}