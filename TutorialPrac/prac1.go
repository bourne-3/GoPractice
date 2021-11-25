package main

import "fmt"

// naked return
func spilit(sum int) (x int, y int) {
	x = sum + 1
	y = sum - 1
	return // naked
}

func twoSum(x int, y int) int {
	return x + y
}

func prac1() {
	fmt.Println("ok")
	fmt.Println(twoSum(1, 2))
	fmt.Println(spilit(10))
}
