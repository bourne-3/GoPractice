package main

import (
	"bytes"
	"fmt"
	"os"
)

func WRexp() {
	//定义零值Buffer类型变量b
	var b bytes.Buffer
	//使用Write方法为写入字符串
	b.Write([]byte("你好"))
	//这个是把一个字符串拼接到Buffer里
	fmt.Fprint(&b,",","http://www.flysnow.org")
	//把Buffer里的内容打印到终端控制台
	b.WriteTo(os.Stdout)

	var p [100]byte
	n,err:=b.Read(p[:])
	fmt.Println(n,err,string(p[:n]))

}

func run12()  {
	// 写入到Buffer
	var b bytes.Buffer
	b.Write([]byte("hello"))  // 把hello写到Buffer
	fmt.Fprint(&b, ",", "中国\n")  // 一个字符串拼接到Buffer里
	b.WriteTo(os.Stdout)

	// =========
	// 读取
	var p [100]byte  // 将数据读取到p中
	n, err := b.Read(p[:])
	fmt.Println(n, err, p[:])

	//data,err:=ioutil.ReadAll(&b)
	//fmt.Println(string(data),err)
	//WRexp()
}