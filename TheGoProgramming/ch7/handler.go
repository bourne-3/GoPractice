package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

// v1 固定格式
//func (d database) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
//	for item, price := range d {
//		fmt.Fprintf(w, "%s: %s\n", item, price)
//	}
//}

// v2 根据switch查找
//func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
//	switch req.URL.Path {
//	case "/list":
//		for item, price := range db {
//			fmt.Fprintf(w, "%s: %s\n", item, price)
//		}
//	case "/price":
//		item := req.URL.Query().Get("item")
//		price, ok := db[item]
//		if !ok {
//			w.WriteHeader(http.StatusNotFound) // 404
//			fmt.Fprintf(w, "no such item: %q\n", item)
//			return
//		}
//		fmt.Fprintf(w, "查找到的key为：%s\n", price)
//	default:
//		w.WriteHeader(http.StatusNotFound)
//		fmt.Fprintf(w, "no such page: %s\n", req.URL)
//	}
//}

// v3 每一个url对应一个函数，而不是耦合在一起成以为switch
func (db database) list(w http.ResponseWriter, req *http.Request) {
	for key, val := range db {
		fmt.Fprintf(w, "%s : %s \n", key, val)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")  // 取出请求的参数
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "找不到这个item %s", item)
		return  // 记住要退出
	}
	fmt.Fprintf(w, "key为 %s, val为 %s", item, price)

}


func handlerRun1() {
	//db := database{"shoes":50, "socks":5}
	//log.Fatal(http.ListenAndServe("localhost:8000", db))
}

func handlerRun2() {
	db := database{"shoes":50, "socks":5}
	mux := http.NewServeMux()
	// 注意 http.HandlerFunc() 是一个类型转换，充当适配器
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))  // 函数值

	log.Fatal(http.ListenAndServe("localhost:8001", mux))
}

func handlerRun3() {
	db := database{"shoes":50, "socks":5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}
