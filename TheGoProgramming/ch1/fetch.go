package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func feachRun() {
	//url := flag.String("url", "defalutUrl", "url that needed to be fetch")

	for _, url := range os.Args[4:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprint(os.Stderr, "fetch: %v \n", err)
			os.Exit(1)
		}
		// 读取
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprint(os.Stderr, "fetch reading: %s %v \n", url, err)
			os.Exit(1)
		}
		fmt.Println(string(b))
	}

}
