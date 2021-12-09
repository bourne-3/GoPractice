package main

import "fmt"

func main() {
	v1 := count(10)
	v2 := count(20)
	path := []Path{{v1, v2}, {v2}}
	fmt.Println(path)
	for _, p := range path {
		p.Translate(3,true)
	}
	fmt.Println(path)
}


// 函数表达式
type count int

func (c count) name()  {
	fmt.Println(c)

}

func (c count) add(val int) count {
	c += count(val)  // 因为underlying的结构是一样的
	return c
}

func (c count) sub(val int) count{
	c -= count(val)
	return c
}

type Path []count

func (p Path) Translate(offset int, add bool) {
	var op func(c count, val int) count
	if add {
		op = count.add
	}else {
		op = count.sub
	}

	for i, _ := range p {
		// 计算后重新存回来
		p[i] = op(p[i], offset)
	}
}