package main

import (
	"crypto/sha256"
	"fmt"
)

func helper1() {
	var arr [5]int
	fmt.Println(arr[1])

	// ... 表示不确定
	arr2 := [...]int{12,13, 14}
	fmt.Println(arr2)

	// 根据索引去初始化
	type Currency int

	const (
		USD Currency = iota // 美元  // 注意这里的常量都是 named type 后的int，因此都是基本类型
		EUR                 // 欧元
		GBP                 // 英镑
		RMB                 // 人民币
	)

	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}

	fmt.Println(RMB, symbol[RMB]) // "3 ￥"

	// 数组的比较
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) // "true false false"
	//d := [3]int{1, 2}
	//fmt.Println(a == d) // compile error: cannot compare [2]int == [3]int

	// 更加复杂的数组比较
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X")) // X
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}

func helper2() {
	// slice
	nums := []int{10, 11, 12, 13, 14}
	reverse(nums[:])
	fmt.Println(nums)

	// equal
	x1, x2  := nums[1:3], nums[1:3]
	f := isEqualSlice(x1, x2)
	fmt.Println(f)
}
func reverse(arr []int)  {
	for i , j := 0,  len(arr) - 1; i < j; i,j = i + 1,j-1{
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func isEqualSlice(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	//遍历所有元素
	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func tryAppend() {
	var runes []rune
	for _,r := range "hello,中国"{
		runes = append(runes, r)
	}

	fmt.Printf("%q \n", runes)  // 单引号围绕的字符字面值，由Go语法安全地转义
	fmt.Println(runes)
}

func appendInt(x []int, y... int) []int {
	// 将y append到 x切片中
	var z []int
	zlen := len(x) +1
	if zlen <= cap(x){
		// 放得下
		z = x[:zlen]
	}else {
		// 需要扩容
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 *len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	// 正式插入
	//z[len(x)] = y

	// 这里改进了，变成可以插入多个值
	copy(z[len(x):], y)
	return z
}

func helper3() {
	//tryAppend()

	// run appendInt()
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}

	fmt.Println("=======")

	var q []int
	q = append(q, 1)
	q = append(q, 2,3)
	q = append(q, 4,5,6)
	q = append(q, q...)  // 注意这里的...
	fmt.Println(q)
}

func nonempty(strs []string)[]string {
	// 返回一个不包含空字符串的切片。 但是这样会更改到原来的切片
	i := 0
	for _, s := range strs{
		if s == "" {
			continue
		}
		strs[i] = s
		i++
	}
	return strs[:i]
}

func nonempty2(strs []string) []string{
	// 使用append来完成
	out := strs[:0]  // 长度为0的切片
	for _, s := range strs{
		if s == "" {
			continue
		}
		out = append(out, s)
	}
	return out
}

func helper4()  {
	data := []string{"one", "", "three"}
	//fmt.Printf("%q\n", nonempty(data)) // `["one" "three"]`
	//fmt.Printf("%q\n", data)           // `["one" "three" "three"]`

	fmt.Printf("%q\n", nonempty2(data)) // `["one" "three"]`
	fmt.Printf("%q\n", data)           // `["one" "three" "three"]`


	// 指定删除
	sliceTmp := []int{0, 1,2,3,4,5,6,7}
	//fmt.Println(removeTarget(sliceTmp, 5))  // 注意，会修改到原来的数组的
	fmt.Println(removeTargetv2(sliceTmp, 1))
}

func simuStack() {
	stack := make([]int, 10)
	val := 2
	// 入栈
	stack = append(stack, val)
	// 出栈
	topVal := stack[len(stack) - 1]
	stack = stack[:len(stack) - 1]
	fmt.Println(topVal)


}

func removeTarget(slice []int, idx int) []int{
	// 找出指定的idx
	//copy(slice[:idx], slice[idx + 1:])
	//copy(slice[idx:], slice[idx + 1:])
	//return slice[:len(slice) - 1]

	copy(slice[idx:], slice[idx+1:])
	return slice[:len(slice)-1]
}
func removeTargetv2(slice []int, idx int) []int{
	// 直接找出对应的值来覆盖掉，如果不在意值的位置的话
	slice[idx] = slice[len(slice) - 1]
	return slice[:len(slice) - 1]
}


func main() {
	//helper1()
	//helper2()
	//helper3()
	//helper4()

	//dedup()
	//tryMap()

	//runStruct()
	movie()
	//tryJSON()
}
