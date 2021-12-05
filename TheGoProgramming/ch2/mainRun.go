package main

func f() *int {
	v := 1;
	return &v
}


func main() {
	// w1
	//p := f()  // 局部变量的地址
	//fmt.Println(*p)

	//runHelper()
	pkgInit()
}
