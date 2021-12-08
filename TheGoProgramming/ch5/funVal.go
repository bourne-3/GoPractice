package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }
func add1(r rune) rune     { return r + 1 }

func double(x int)(result int) {
	// 使用defer和匿名函数可以拿到返回值
	defer func() {fmt.Printf("double(%d) = %d\n", x, result)}()
	return x + x
}

func defer1(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	defer1(x - 1)  // 递归调用
}

func helperRun() {
	double(3)
	defer1(3)
}

func m2(vals ...int) int {
	// 1 可变参数
	total := 0
	for _, n := range vals {
		total += n
	}
	return total
}

func m4(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html") {
		resp.Body.Close()
		return fmt.Errorf("%s  type %s", url, ct)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}
	// TODO 写一个forEachNode
	fmt.Println(doc, visitNode)
	return nil
}

func m3() {
	vals := []int{1, 2, 3, 4, 5}
	fmt.Println(m2(vals...))

	linenum, name := 12, "count"
	errorf(linenum, "undefined: %s %s", name, "bourne") // "Line 12: undefined: count"
}

func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Line %d : ", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}

func s2() func() int {
	// 闭包函数
	var x int
	// 返回的是一个匿名函数
	return func() int {
		x++
		return x * x
	}
}

func funV2() {
	f := s2()
	//the inner function can refer to var iables fro m the enclosing function
	fmt.Println(f()) // 具有记忆的特性
	fmt.Println(f())
	fmt.Println(f())
}

func funV1() {
	// 1 func val
	f := square
	fmt.Println(f(10))

	f = negative
	fmt.Println(f(10))

	//f = product  // err cannot use product (type func(int, int) int) as type func(int) int in assignment

	// 2 nil的函数值
	//var f2 func(int) int
	//f2(10)  // panic

	// 3 函数值作为参数传递
	fmt.Println(strings.Map(add1, "HAL-9000")) // "IBM.:111"
	fmt.Println(strings.Map(add1, "VMS"))      // "WNT"
	fmt.Println(strings.Map(add1, "Admix"))    // "Benjy"
	fmt.Println(strings.Map(add1, "ab"))       //

}
