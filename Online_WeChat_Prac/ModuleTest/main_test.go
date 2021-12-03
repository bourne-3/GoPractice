package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

// 单元测试
func TestAdd(t *testing.T) {
	sum := Add(10, 5)
	if sum == 15 {
		t.Log("the result is right")
	}else {
		t.Fatal("wrong result")
	}
}

// 表组测试
func TestMultiAdd(t *testing.T) {
	sum := Add(10, 5)
	if sum == 15 {
		t.Log("the result is right")
	}else {
		t.Fatal("wrong result")
	}

	sum2 := Add(1, 3)
	if sum2 == 4 {
		t.Log("the result is right")
	}else {
		t.Fatal("wrong result")
	}
}

func init() {
	Routes()  // 配合下方使用
}

// 网络测试
func TestSendJSON(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/sendjson", nil)
	if err != nil {
		t.Fatal("创建request出错了")
	}

	rw := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rw, req)

	log.Println("code: ", rw.Code)
	log.Println("body: ", rw.Body.String())
}


// 测试覆盖率
func TestTag(t *testing.T) {
	Tag(1)
	Tag(2)
	Tag(3)
	Tag(100)

}