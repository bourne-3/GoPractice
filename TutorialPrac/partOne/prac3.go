package main

import (
	"fmt"
	"strings"
)

// structs
type Vertex struct {
	X int
	y int
}

type Record struct {
	Salary, cost float64
}

func tryMap() {
	m1 := make(map[string]int)
	m1["lucy"] = 100
	m1["bourne"] = 400
	m1["kim"] = 200
	fmt.Println(m1)

	m2 := make(map[string]Record)
	m2["member1"] = Record{12500, 400}
	fmt.Println(m2["member1"])

	// map literals
	m3 := map[string]Record{
		"member1" : {40000, 100},
		"member2" : {30000, 200},
	}
	fmt.Println(m3["member1"])

	// 判断是否存在， 这里和python里面的拆包还不大一样
	content, isExist := m3["member2"]
	fmt.Println(content, isExist)

	c2, ex := m3["random"]
	fmt.Println(c2, ex)
}

func tryRange() {
	// 有三种方法可以创建slice
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}  // way1
	//var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}  // way2

	// way3
	//var pow []int
	//pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	printSlice(pow)

	// 使用range  // 类似于python中的enumerate 以及增强for循环
	for i, val := range pow{
		fmt.Print(i, val, "  ||  ")
	}
	fmt.Println()
	// ==============
	// 可以省略前半部分以及后半部分
	fmt.Println("================")
	nums := make([]int, 10)
	printSlice(nums)

	for idx, _ := range nums{
		fmt.Print(idx, " ")
		nums[idx] = 1 << idx  // 2 ^ idx
	}
	fmt.Println()

	for _, val := range  nums{
		fmt.Print(val, " ")
	}
	fmt.Println()

}

func tryMake() {
	//nums := make([]int, 3)
	//printSlice(nums)

	nums2 := make([]int, 0, 5) // first is len, second is cap
	printSlice(nums2)

	nums3 := nums2[:2]
	printSlice(nums3)

	//nums4 := nums3[2:]
	nums4 := nums3[2:5]
	printSlice(nums4)

	matrix1 := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	matrix1[0][0] = "X"
	matrix1[len(matrix1)-1][len(matrix1[0])-1] = "O"
	//fmt.Println(len(matrix1), len(matrix1[0]))

	for i := 0; i < len(matrix1); i++ {
		fmt.Printf("%s\n", strings.Join(matrix1[i], " "))
	}

	//==================
	// append
	var t2 []int
	fmt.Println(t2)
	t2 = append(t2, 1)
	printSlice(t2)
	t2 = append(t2, 2, 3, 4, 5)
	fmt.Println(t2)
	printSlice(t2)
}

func trySlice() {
	// 重要：Slices are like references to arrays
	name := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	n1 := name[:3]
	n2 := name[1:]
	fmt.Println(n1, n2)

	n1[2] = "bourne" // 切片是一个引用 不存储值
	fmt.Println(n1, n2)

	// slice literals  ||  其实literals就是分配内存加上初始化而已
	q := []int{11, 12, 13} // creates the  array , then builds a slice that references it:
	fmt.Println(q)         // 注意这里的q仍然是一个slice

	// 必须有两个{}
	s1 := struct {
		id   int
		name string
	}{20201050008, "borune"}
	fmt.Println(s1)

	// struct切片
	s2 := []struct {
		pid   int
		pName string
	}{
		{1, "iphone"},
		{2, "ipad"},
		{3, "macBook"},
	}
	fmt.Println(s2)

	// ==========
	// 不断在新的切片上操作
	fmt.Println("=================")
	nums := []int{2, 3, 5, 7, 11, 13}
	nums = nums[:0] // 3个
	fmt.Println(nums)

	nums = nums[:2]
	fmt.Println(nums) // 3 5

	nums = nums[2:]
	fmt.Println(nums) // []

	// ================
	fmt.Println("==============")
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)
	s = s[:1]
	printSlice(s)
	s = s[:4]
	printSlice(s)
	s = s[2:] // drop the first two values
	printSlice(s)

	// ================= nil
	var t1 []int
	fmt.Println(t1)
	if t1 == nil {
		fmt.Println("nil ")
	}

}

func tryArray() {
	var a [2]int // 先数量再类型
	a[0] = 21
	a[1] = 22
	fmt.Println(a)

	s1 := [3]string{"bourne", "lucy", "kim"}
	fmt.Println(s1)

	// 使用slice
	slice1 := s1[:2] // [low:hight] default 0 or len()
	fmt.Println(slice1)

}

func tryStructs() {
	fmt.Println(Vertex{1, 2})
	// 获得struct的内容
	v2 := Vertex{12, 23}
	fmt.Println(v2.X)

	// 指针结合struct
	v3 := Vertex{13, 32}
	p := &v3
	p.X = 31 // 使用指针修改struct的内容  without the explicit dereference.
	fmt.Println(*p, v3, p)

}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// 指针
func pointer1() {
	var p *int // 指针
	i := 15
	p = &i
	fmt.Println("这是一个地址", p) // 这是一个地址
	fmt.Println("具体的值", *p)  // 具体的值
	*p = 31
	fmt.Println("*p = 31之后：", *p) // 具体的值
}
