package main

import "fmt"

type user struct {
	name string
	email string
}
type admin struct {
	user
	level string
}
func main() {
	ad:=admin{user{"张三","zhangsan@flysnow.org"},"管理员"}
	fmt.Println("可以直接调用,名字为：",ad.name)
	fmt.Println("也可以通过内部类型调用,名字为：",ad.user.name)
	fmt.Println("但是新增加的属性只能直接调用，级别为：",ad.level)


	ad =admin{user{"张三","zhangsan@flysnow.org"},"管理员"}
	sayHello(ad.user)//使用user作为参数
	sayHello(ad)//使用admin作为参数
}
type Hello interface {
	hello()
}
func (u user) hello(){
	fmt.Println("Hello，i am a user")
}
func sayHello(h Hello){
	h.hello()
}
