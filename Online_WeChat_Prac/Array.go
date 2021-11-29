package main

import "fmt"
// 只写那些没有接触过的

func run1() {
	// 初始化零值
	var arr [3]int  // 可以直接使用，默认是对应的0值
	var strArr [3]string
	var boolArr [3]bool
	fmt.Println(arr, strArr, boolArr)

	// 长度推导
	lenArr := [...]int{10,2,3}
	fmt.Println(lenArr)

	// 指定初始化
	pointArr := [...]int{1:5,6:2}
	fmt.Println(pointArr)

	// 数组相互赋值
	var copyPoint [len(pointArr)]int = pointArr  // 长度必须一致才可以copy
	fmt.Println(copyPoint)

	// 指针数组
	arr2 := [5]*int{1:new(int), 3:new(int)}  // 指定分配了空间
	*arr2[1] = 10
	fmt.Println(arr2, arr2[1], *arr2[1])
		// 要给下标为0的位置赋值，需要先分配内存
	arr2[0] = new(int)
	*arr2[0] = 400
	fmt.Printf("%v 的地址为 %v", *arr2[0], arr2[0])

	// 数组的修改
	arr3 := [...]int{10,12,13,15}
		// 使值传递修改
	modify(arr3)
	fmt.Println(arr3)

		// 使用指针传递修改
	modifyPoint(&arr3)
	fmt.Println(arr3)
}

func modifyPoint(arr *[4]int) {
	arr[0] = 1
}

func modify(arr3 [4]int) {
	arr3[0] = 1
}

