package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func dedup() {
	seen := make(map[string]bool) // a set of strings
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}

func tryMap() {
	// 1 map的值是不可以取地址的
	ages := make(map[string]int)
	ages["cidy"] = 52
	ages["bourne"] = 11
	ages["lucy"] = 12
	//_ := &ages["bourne"]  // error

	// 2 在指定顺序下遍历map
	agesSlice := make([]string, 0, len(ages)) // len应该设置为0，我写错了一开始
	for n := range ages {
		// 将所有k放到一个切片中
		agesSlice = append(agesSlice, n)
	}
	// 排序
	sort.Strings(agesSlice)
	//// 注意，这里agesSlice是一个切片
	//for value, key := range agesSlice{
	//	fmt.Printf("rang取出来的key为：%s, val为：%d \n", key ,value)
	//}

	fmt.Printf("经过排序后的map如下\n")
	for _, k := range agesSlice {
		fmt.Printf("名字：%s  年龄：%d \n", k, ages[k])
	}

	// 3 需要声明之后才能使用
	//var team map[string]int
	//team["bourne"] = 15
	//fmt.Println(team)  // panic  没有声明make(map[string]int)

	// 4 多值返回
	age, ok := ages["bourne"]
	fmt.Println(age, ok)

	// 5
	var mapB map[string]int
	mapB = ages
	f := isEqual(mapB, ages)
	fmt.Println(f)

	// 6 使用一个map实现一个set
	set := make(map[rune]bool)
	inVal := []rune("中过中国")
	for _, v := range inVal{
		// 这里的if是核心
		if !set[v] {
			set[v] = true
		}
	}
	fmt.Println(set)

	// 7 当要求map的key是一个切片的时候，需要如何处理

}

func mapConvert() {

}
var m = make(map[string]int)

// 利用k来转换
func k(list []string) string {
	return fmt.Sprintf("%q", list)
}

func Add(list []string) {
	m[k(list)]++
}
func Count(list []string) int{
	return m[k(list)]
}




// 5 两个map之间的比较
func isEqual(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	// 注意这里用到了 多值返回。以及for之前进行一个初始化定义
	for k, xv := range a{
		if yv, ok := b[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
