package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var (
	count int
	m sync.Mutex
)

func h1() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	count++
	fmt.Println("handler被触发")
	m.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func h2() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}


func h3() {
	http.HandleFunc("/", hander3)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// 结合gif图片，嘿嘿成功
func h4() {
	http.HandleFunc("/gif", func(writer http.ResponseWriter, request *http.Request) {
		lissajous(writer)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func counter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "目前一共点击的次数是 = %v\n", count)
}

func hander3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k,v := range r.Header{
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err!=nil{
		log.Print(err)
	}
	// Form是表单的意思、

	for k, v := range r.Form{
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}

}