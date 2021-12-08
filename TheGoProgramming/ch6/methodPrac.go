package main

import (
	"fmt"
	"math"
	"net/url"
)

type Point struct {
	X, Y float64
}

// function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

// method
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

// 定义一个命名类型并且声明同样的方法名称
// A Path is a journey connecting the points with straight lines.
type Path []Point
// Distance returns the distance traveled along the path.
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

// 指针接收器
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func run1() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))

	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance()) // "12"

	// 指针作为receiver

	//1
	r := &Point{1, 2}  // 这里r是一个地址
	r.ScaleBy(2)
	fmt.Println(*r)

	// 2 通过值也是可以直接调用的
	r2 := Point{2, 3}  // 这里r2是一个值
	r2.ScaleBy(2)  // 这里做了一个自动类型转换 (&r2).ScaleBy()
	fmt.Println(r2)

	// 3 通过指针也可以调用值的方法
	pptr := &Point{3, 4}
	pptr.Distance(Point{1,2})  // 这里的Distance的receiver是一个值，但是可以用指针调用
}

func run2() {
	m := url.Values{"lang":{"en"}}
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang"))

}
