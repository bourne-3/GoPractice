package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)


// 这个函数有问题，一直是处于死循环，不会结束
func dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	// 输出值
	for k, v := range counts{
		if v > 1 {
			fmt.Printf("%d\t%s\n", v, k)
		}
	}
}

// 输入有多个：可以是用命令行输出，也可以是从文件输入
func dup2() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0{
		countLines(os.Stdin, counts)
	}else {
		for _, args := range files{
			f, err := os.Open(args)
			if err != nil{
				fmt.Fprintf(os.Stderr, "dup2 : %v \n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	// 输出值
	for k, v := range counts{
		if v > 1 {
			fmt.Printf("%d\t%s\n", v, k)
		}
	}

}

func countLines(f *os.File, count map[string]int) {
	in := bufio.NewScanner(f)
	for in.Scan(){
		t := in.Text()
		count[t]++
	}
}

// 一次性将文件中的内容加到内容中
func dup3() {
	args := os.Args[3:]
	fmt.Println(args)
	counts := make(map[string]int)
	for _, a := range args{
		data, err := ioutil.ReadFile(a)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3 : %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n"){
			counts[line]++
		}
	}

	// 输出值
	for k, v := range counts{
		if v >= 1 {
			fmt.Printf("%d\t%s\n", v, k)
		}
	}
}

func readFiletest() {
	// 原本是使用绝对路径，这里更改为使用flag包来相助
	path := flag.String("fpath", "defalutValue", "path that access to the real file")
	flag.Parse()  // 必须的操作
	//data, err := ioutil.ReadFile("F:\\Go_dir\\CodePrac\\GoPractice\\TheGoProgramming\\ch1\\Script.txt")
	data, err := ioutil.ReadFile(*path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
	}
	fmt.Println(string(data))
	fmt.Println(os.Args[1:])
}

func testFlag() {
	path := flag.String("fpath", "Script.txt", "Path to access to file")
	flag.Parse()
	fmt.Println(*path)
}


