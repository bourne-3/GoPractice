package main

import (
	"bytes"
	"fmt"
	"strings"
)

func basename1(s string) string{
	// without libraries
	// Discard last '/' and everything before.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// Preserve everything before last '.'.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func basename2(s string) string{
	slash := strings.LastIndex(s, "/")
	s = s[slash + 1:]
	if dot := strings.LastIndex(s, ".");dot>=0 {
		s = s[:dot]
	}
	return s

}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func printints(vals []int) string{
	var buf bytes.Buffer
	buf.WriteByte('[')

	for i, v := range vals{
		if i > 0 {
			buf.WriteByte(',')
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	// 收尾
	buf.WriteByte(']')
	return buf.String()
}

func h1() {
	//s := basename1("a/v/b/runCMD.go")
	//s := basename2("a/v/b/runCMD.go")
	//fmt.Println(s)

	//s := comma("121454681")
	//fmt.Print(s)

	//fmt.Println(printints([]int{1,2,5}))

	fmt.Println(e3_10("1523"))
}

func h2() {
	//convert()
	constantVal()
}

func main() {
	//h1()
	h2()
}