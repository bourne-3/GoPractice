package main

import "fmt"

func usrAppend() {
	// append；但切片容量足够的话，会是引用地址，如果容量不够，就会创建新的切片
	// 创建新切片的时候，最好要让新切片的长度和容量一样

	// 会导致原数组的改变
	slice1 := []int{11,12,13,14,15,16,17,18}
	slice2 := slice1[1:3]  // 13 14
	slice2 = append(slice2, 100)  // 将14修改为100
	fmt.Println(slice1, slice2)  // slice1和slice2两者底层都会修改

	// 这样使用append的话就导致原来数组的改变了
	slice3 := slice1[4:7:7]
	slice3 = append(slice3, 500)
	fmt.Println(slice1, slice3)  // 100 500

	// 切片作为参数传递
	slice4 := []int{21, 22, 23}
	fmt.Printf("函数外部的slice地址：%p \n", &slice4)
	modifySlice(slice4)
	fmt.Println(slice4)

}

func modifySlice(slice []int) {
	fmt.Printf("函数内部的slice地址：%p \n", &slice)
	slice[0] = 100
}


func run2() {
	slice1 := make([]int, 5)
	fmt.Println(slice1)  // 同样也是默认零值

	slice2 := []int{1,4,6}  // 字面量创建 literal
	fmt.Println(slice2)

	// nil切片和空切片
	var nilSlice []int  // 底层数组的指针为nil
	var emptySlice []int = []int{}  // 对应的指针是个地址。
	fmt.Println(nilSlice, emptySlice)

	// 共享内存
	slice3 := []int{5, 10 ,2}
	slice4 := slice3[:3]
	slice4[0] = 1
	fmt.Println(slice3) // slice4修改会同时也修改了slice3的内容

	// len和cap
	slice5 := []int{1,2,3,4,5,6,7,8,9,10}
	slice6 := slice5[:5]  // 注意这里slice6的cap
	fmt.Println(len(slice6), cap(slice6))  // 5 10?

	// 三个参数，切片同时指定新的cap
	slice7 := slice5[1:4:8]
	fmt.Println(len(slice7), cap(slice7)) // 3 7

	// 使用append
	usrAppend()
}