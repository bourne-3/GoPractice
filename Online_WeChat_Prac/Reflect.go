package main

import (
	"fmt"
	"reflect"
)

type user struct {
	name string
	age int
}

func (u user) Say(s string) {
	fmt.Printf("%v 说了声 %s", u.name, s)
}

func run14() {
	u := user{name: "bourne", age: 19}

	// reflect中的 TypeOf 和 ValueOf
	t := reflect.TypeOf(u)
	fmt.Println(t)
	v := reflect.ValueOf(u)  // reflect.ValueOf函数把任意类型的对象转为一个reflect.Value
	fmt.Println(v)

	// tips，使用fmt包中的Printf函数也可以满足我们的要求
	fmt.Printf("u的类型为 %T , 值为 %v \n", u, u)


	// 2 反转为原始类型
	u1 := v.Interface().(user)  // 这个u1就变成原来的user对象了
	fmt.Println(u1)

	// 根据reflect.Value获得reflect.Type
	t1 := v.Type()
	fmt.Println(t1)

	// 提取底层类型  Kind()
	fmt.Println(t.Kind())

	fmt.Println("==================")
	//// 反射的使用
	// 遍历字段和方法
	// 字段
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Name)
	}

	// 方法
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Println(t.Method(i).Name)
	}

	// 反射修改字段的值
	x := 2
	valOfx := reflect.ValueOf(&x)
	valOfx.Elem().SetInt(100)
	fmt.Println(x)

	u2 := user{name: "lucy", age: 32}  // 结构体的修改
	valOfuser :=  reflect.ValueOf(&u2.name)
	fmt.Println(valOfuser.Elem().CanSet())
	valOfuser.Elem().SetString("Jordan")
	fmt.Println(u2)

	// 反射调用方法
	u3 := user{name: "Camber", age: 45}
	valOfu3 := reflect.ValueOf(u3)
	mSay := valOfu3.MethodByName("Say")
	fmt.Println(mSay.IsValid())
	args := []reflect.Value{reflect.ValueOf("快要过年啦")}
	fmt.Println(mSay.Call(args))

}

func he() {
	u:=User2{"张三",20}
	v:=reflect.ValueOf(u)

	mPrint:=v.MethodByName("Print")
	args:=[]reflect.Value{reflect.ValueOf("前缀")}
	fmt.Println(mPrint.Call(args))

}
type User2 struct{
	Name string
	Age int
}
func (u User2) Print(prfix string){
	fmt.Printf("%s:Name is %s,Age is %d",prfix,u.Name,u.Age)
}