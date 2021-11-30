package main

import "fmt"

func run3() {
	// 字面量创建  当然 也可以使用make来创建
	namesMap := map[string]int{"age":19}
	namesMap["date"] = 15
	fmt.Println(namesMap)

	//var nilMap map[string]int  // 需要使用make创建先

	// 获得两个val，后面是bool的
	age, exist := namesMap["age"]
	fmt.Println(age, exist)
	delete(namesMap, "age")
	fmt.Println(namesMap)
}