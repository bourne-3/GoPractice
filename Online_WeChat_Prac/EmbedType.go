package main

import "fmt"

type User struct {
	name string
	email string
}

type Admin struct {
	//user User  // 我一开始这里写错了， 是应该只写个User即可
	User
	level string
}

// 内部实现了接口，外部同样也会实现接口
type palyI interface {
	basketball()
}

func (u User) basketball()  {
	fmt.Println( "正在打篮球")
}

func Execice(p palyI) {
	p.basketball()
}


// 实现sayHello方法
func (u User) sayHello() {
	fmt.Println("user say Hello")
}

func (a Admin) sayHello() {
	fmt.Println("Admin say Hello")
}

func run7() {
	// 嵌套结构
	a := Admin{User{"bourne", "zebinli@gdufs.edu.cn"},"hight-Manager"}
	//ad:=Admin{User{"张三","zhangsan@flysnow.org"},"管理员"}
	fmt.Println(a)
	//fmt.Println(a.level, a.user, a.user.email, a.user.name)  // 不能够直接a.email
	fmt.Println(a.level, a.User, a.name, a.email)  // 不能够直接a.email

	a.sayHello()
	a.User.sayHello()

	// 内部类型实现了接口
	var u User
	u = a.User
	Execice(u)

	Execice(a)
}
