package main

import (
	"fmt"
	"time"
)

func run4() {
	// 1 方法值
	p := Point{1, 2}
	q := Point{4, 6}
	methodVal := p.Distance  // 方法值
	fmt.Println(methodVal(q))

	var ori Point
	fmt.Println(methodVal(ori))

	scaleP := p.ScaleBy
	scaleP(2)
	scaleP(3)
	fmt.Println(p)
	scaleP(10)
	fmt.Println(p)

	r := new(Rocket)  // 返回的是一个指针
	time.AfterFunc(3 * time.Second, r.Launch)

	fmt.Println("=========")

	// 2 方法表达式
	distance := Point.Distance
	fmt.Println(distance(q, p))
	fmt.Printf("%T \n", distance)

	scale := (*Point).ScaleBy
	scale(&p, 2)  // 第一个参数是receiver
	fmt.Println(p)
	fmt.Printf("%T \n", scale)


}

// 例子
type PathSlice []Point

func (p Point) Add(q Point) Point {
	// 返回一个Point
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

func (p PathSlice) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add  // 函数表达式
	}else {
		op = Point.Sub
	}

	// 这个op已经被赋值了，可以直接使用了
	for i, _ := range p {
		p[i] = op(p[i], offset)  // 第一个参数是函数receiver
	}
}

// 火箭发射
type Rocket struct {

}

func (r *Rocket) Launch() {
	fmt.Println("火箭发射！")
}