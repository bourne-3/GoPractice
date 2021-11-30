package utils

import "fmt"

type count int

type User struct {
	Name string
	salary float64  // 外部不可以访问
}

type consumer struct {
	Name string
	salary float64
}

type Admin struct {
	consumer  // consumer是未export的
}

// 一个可供外部使用的接口
type Loginer interface {
	Login()
}

type defaultLogin int
// 定义一个方法
func (d defaultLogin) Login() {
	fmt.Println("login in...")
}

func NewLoginer() defaultLogin {
	return defaultLogin(10)
}

func NewVal() count {
	return count(50)
}
