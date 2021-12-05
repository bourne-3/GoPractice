package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

//重写reverse函数，使用数组指针代替slice。

func helper1() {
	s := []int{1, 2, 3, 4, 5}
	rotate_ints(s,3)
	fmt.Println(s)

	r1 := filteDup([]string{"a","c","c","h"})
	fmt.Println(r1)

	r2 := squashSpace([]byte("ni   hao    ya"))
	fmt.Println(string(r2))

	by2 := []byte("a b c")
	reverse(by2)
	fmt.Println(string(by2))
}


func reverse(in [] byte) {
	buf := make([]byte, 0, len(in))
	i := len(in)

	for i > 0 {
		_, s := utf8.DecodeLastRune(in[:i])
		buf = append(buf, in[i-s:i]...)
		i -= s
	}
	copy(in, buf)
}



func squashSpace(bytes []byte) []byte {
	out := bytes[:0]
	var last rune

	for i := 0; i < len(bytes); {
		r, s := utf8.DecodeRune(bytes[i:])

		if !unicode.IsSpace(r) {
			out = append(out, bytes[i:i+s]...)
		} else if unicode.IsSpace(r) && !unicode.IsSpace(last) {
			out = append(out, ' ')
		}
		last = r
		i += s
	}
	return out
}

func filteDup(s []string) []string {
	// 快慢指针的思路
	if len(s) == 0 {
		return s
	}

	i := 0
	for j := 1; j < len(s); j++ {
		if s[j] != s[i] {
			i++
			s[i] = s[j]
		}
	}
	return s[:i+1]
}

func rotate_ints(s []int, n int) []int{
	n = n % len(s)
	if n > 0 && n < len(s) {
		temp := make([]int, n)
		copy(temp, s[:n])

		copy(s, s[n:])
		copy(s[len(s)-n:], temp)
	}
	return s
}

func mainE2() {
	helper1()
	//testCopy()
}

func testCopy() {
	// 如果原本的切片就有值 怎么复制？
	a := []int{10,11,12}
	b := []int{22,23}
	copy(a, b)
	fmt.Println(a)
}