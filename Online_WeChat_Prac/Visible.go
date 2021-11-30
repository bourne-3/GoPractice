package main

import (
	"GoPractice/Online_WeChat_Prac/utils"
	"fmt"
)
func run8() {
	c := utils.Count(5)
	fmt.Println(c)  // 来在外部的包

	// 外部包的函数
	outVal := utils.NewVal()
	fmt.Println(outVal)

	// 使用别人定义好的接口
	loginer := utils.NewLoginer()
	loginer.Login()

	// 结构体中包含可以非export
	u := utils.User{Name: "bourne"}
	fmt.Println(u)  // 仅能访问u.Name 其他的访问不到


	// 嵌套可视访问
	var a utils.Admin
	a.Name = "Lucy"  // 可以访问到Admin中consumer中的Name，但是不能访问consumer
	fmt.Println(a)

	a2 := utils.Admin{}  // 第二种初始化方法
	a2.Name = "Kim"
	fmt.Println(a2)



}