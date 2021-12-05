package main

import (
	"bytes"
)

func e3_10(s string) string{
	b := &bytes.Buffer{}
	pre := len(s) % 3
	// Write the first group of up to 3 digits.
	if pre == 0 {
		pre = 3
	}
	b.WriteString(s[:pre])
	// Deal with the rest.
	for i := pre; i < len(s); i += 3 {
		b.WriteByte(',')
		b.WriteString(s[i : i+3])
	}
	return b.String()
}

// 3.11  函数异位词
func isAnagram(s string, t string) bool {
	aFreq := make(map[rune]int)
	bFreq := make(map[rune]int)
	for _, c := range s{
		aFreq[c]++
	}

	for _, c := range t{
		bFreq[c]++
	}

	//for k, _ := range aFreq{
	//	if aFreq[k] != bFreq[k] {
	//		return false
	//	}
	//}

	for k, v := range aFreq{
		if bFreq[k] != v {
			return false
		}
	}


	for k, v := range bFreq{
		if aFreq[k] != v {
			return false
		}
	}

	return true
}

func ex3_13() {
	// 注意 这里不需要使用:=去定义以及初始化，也不需要使用var。 而且如果没有给出类型的话，是可以自动推导的
	const (
		KB = 1000
		MB = 1000 * KB
		GB = 1000 * MB
		PB = 1000 * GB
	)

}
