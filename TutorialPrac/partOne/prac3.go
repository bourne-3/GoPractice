package main

import "fmt"

// structs
type Vertex struct {
	X int
	y int
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

	n1[2] = "bourne"  // 切片是一个引用 不存储值
	fmt.Println(n1, n2)


	// slice literals  ||  其实literals就是分配内存加上初始化而已
	q := []int {11, 12, 13}  // creates the  array , then builds a slice that references it:
	fmt.Println(q)  // 注意这里的q仍然是一个slice

	// 必须有两个{}
	s1 := struct {
		id int
		name string
	}{20201050008, "borune"}
	fmt.Println(s1)

	// struct切片
	s2 := []struct {
		pid int
		pName string
	}{
		{1, "iphone"},
		{2, "ipad"},
		{3, "macBook"},
	}
	fmt.Println(s2)

	// ==========
	fmt.Println("=================")
	nums := []int{2, 3, 5, 7, 11, 13}
	nums = nums[1:4]
	fmt.Println(nums)


}

func tryArray() {
	var a [2]int  // 先数量再类型
	a[0] = 21
	a[1] = 22
	fmt.Println(a)

	s1  := [3]string{"bourne", "lucy", "kim"}
	fmt.Println(s1)

	// 使用slice
	slice1 := s1[:2]  // [low:hight] default 0 or len()
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
	p.X = 31  // 使用指针修改struct的内容  without the explicit dereference.
	fmt.Println(*p, v3, p)


}


// 指针
func pointer1() {
	var p *int  // 指针
	i := 15
	p = &i
	fmt.Println("这是一个地址", p)  // 这是一个地址
	fmt.Println("具体的值", *p)  // 具体的值
	*p = 31
	fmt.Println("*p = 31之后：", *p)  // 具体的值
}