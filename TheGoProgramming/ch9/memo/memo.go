package memo

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}


type Memo struct {
	f Func
	cache map[string]result
}

type result struct {
	value interface{}
	err error
}

// 重命名
type Func func(key string) (interface{}, error)

func New(f Func) *Memo {
	return &Memo{
		cache: make(map[string]result),
		f : f,
	}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	res, ok := memo.cache[key]
	if !ok{
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err
}

func Run1() {
	m := New(httpGetBody)
	urls := []string{"https://www.cnblogs.com/hi3254014978/category/1763651.html?page=1", "https://leetcode-cn.com/study-plan/lcof/?progress=5rqgvle",
	"https://www.cnblogs.com/hi3254014978/category/1763651.html?page=1", "http://www.echangwang.com/pic/09/4069.html", "http://www.echangwang.com/pic/09/4069.html"}
	for _, url := range urls{
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}
}