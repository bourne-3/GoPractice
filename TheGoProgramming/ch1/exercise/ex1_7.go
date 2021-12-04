package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func run1_7() {
	// 使用io.Copy 替换 ioutil.ReadAll
	for _, url := range os.Args[1:] {

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprint(os.Stderr, "fetch: %v \n", err)
			os.Exit(1)
		}
		// 读取
		//b, err := ioutil.ReadAll(resp.Body)
		//if err != nil {
		//	fmt.Fprint(os.Stderr, "fetch reading: %s %v \n", url, err)
		//	os.Exit(1)
		//}
		//fmt.Println(string(b))

		_, err2 := io.Copy(os.Stdout, resp.Body)
		if err2 != nil {
			fmt.Fprint(os.Stderr, "fetch: %v \n", err2)
			os.Exit(1)
		}
	}
}

func run1_8() {
	// 使用io.Copy 替换 ioutil.ReadAll
	for _, url := range os.Args[1:] {
		pre := "https://"
		if !strings.HasPrefix(url, pre) {
			url = pre + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprint(os.Stderr, "fetch: %v \n", err)
			os.Exit(1)
		}


		_, err2 := io.Copy(os.Stdout, resp.Body)
		if err2 != nil {
			fmt.Fprint(os.Stderr, "fetch: %v \n", err2)
			os.Exit(1)
		}
		fmt.Printf("\n The StatusCode is:%v", resp.StatusCode)
	}
}