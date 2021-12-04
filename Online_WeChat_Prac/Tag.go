package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func run15() {
	// json转换为对象
	var a Animal
	h:=`{"name":"pig","age":15}`
	err := json.Unmarshal([]byte(h), &a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)

	// 对象转换为json
	newJson, err := json.Marshal(&a)  // 这里输入a和&a都可以
	fmt.Println(string(newJson))

	// 反射获取字段的Tag
	t := reflect.TypeOf(a)
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag
		fmt.Println(tag)
		//fmt.Println(t.Kind())
	}

	// 多个Tag的情况
	for i := 0; i < t.NumField(); i++ {
		tag2 := t.Field(i).Tag.Get("bson")  // 多个tag，根据tag来取出具体的值
		fmt.Println(tag2)
	}

}

type Animal struct {
	//Name string `name`  // 这里后面就是Tag例如
	//Age int `age`

	Name string `json:"name" bson:"b_name"`  // 这里后面就是Tag例如
	Age int `json:"age" bson:"b_age"`
}