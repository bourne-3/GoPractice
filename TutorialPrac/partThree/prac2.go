package main

import "fmt"
// 互斥锁


func addMap(myMap map[string]int, key string) {
	myMap["money"]++
}

// 先写一个会出现脏读的
func run2() {
	myMap := make(map[string]int)
	myMap["money"] = 0
	for i := 0; i < 100; i++ {
		go addMap(myMap, "money")
	}
	fmt.Println(myMap["money"])
}


