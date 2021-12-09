package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func errRun4() {
	err1 := errors.New("ConstomErr")
	fmt.Println(err1)

	err2 := fmt.Errorf("出现了错误，自己写 %s", "ErrCost")
	fmt.Println(err2)
}

func TypeAssert() {
	// 1 具体类型
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File)  // 后面是具体的
	fmt.Printf("%T\n", f)
	//c := w.(*bytes.Buffer)
	//fmt.Printf("%T", c)

	fmt.Println("============")
	// 2 抽象类型
	var w2 io.Writer
	w2 = os.Stdout
	rw := w2.(io.ReadWriter)  // 后面是抽象的
	fmt.Printf("rw  %T\n", rw)  // w2只可以调用write，而rw还可以在此基础上调用Read()。证明rw已经是io.ReadWriter类型了
	//rw.Read()

	w2 = rw
	w2 = rw.(io.Writer)

	// 3 多值返回
	var w3 io.Writer = os.Stdout
	if f3, ok := w3.(*os.File); ok {
		fmt.Println(f3)
	}

}

func m4() {
	_, err := os.Open("/no/such/file")
	fmt.Println(err) // "open /no/such/file: No such file or directory"
	fmt.Printf("%#v\n", err)

	_, err2 := os.Open("/no/such/file")
	fmt.Println(os.IsNotExist(err2)) // "true"  IsNotExist


}

func errMainRun()  {
	//TypeAssert()
	m4()
}
