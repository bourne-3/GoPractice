package main

import (
	"fmt"
	"image/color"
	"math"
	"net/url"
	"sync"
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
	fmt.Println(m.Get("q"))  // ""
	fmt.Println(m.Get("item"))
	fmt.Println(m["item"])

	m = nil
	fmt.Println(m.Get("item"))
	m.Add("item", "3")
}

type ColoredPorint struct {
	Point  // 嵌入了Point
	Color color.RGBA
}

// 还可以嵌入指针
type ColoredPoint2 struct {
	*Point
	Color color.RGBA
}

// 当一个结构体有多个匿名內部struct
type ColoredPoint3 struct {
	Point
	color.RGBA
}

func run3() {
	// 1 嵌入结构体
	var cp ColoredPorint
	cp.X = 1  // 这个X是嵌入Point的内容
	fmt.Println(cp.Point.X)
	cp.Point.Y = 2
	fmt.Println(cp.Y)

	// 2 使用嵌入的结构体的方法
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPorint{Point{1, 1}, red}
	q := ColoredPorint{Point{5,4 }, blue}
	//fmt.Println(p.Distance(q.Point))  // 传入的需要是q.Point
	//fmt.Println(p.Distance(q))  // 单纯传入q是错的
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point))

	fmt.Println("===========")
	// 3 嵌入指针
	p2 := ColoredPoint2{&Point{1, 1},red}
	q2 := ColoredPoint2{&Point{5,4}, blue}
	fmt.Println(p2.Distance(*q2.Point))  // 这里主要使用 *q2.Point
	q2.Point = p2.Point
	fmt.Println(*p2.Point, *q2.Point)

	fmt.Println("===========")

}
// 4 多个匿名内部结构体
var (
	mu sync.Mutex
	mapping = make(map[string]string)
)

func Lookup1(key string) string {
	mu.Lock()
	v := mapping[key]
	mu.Unlock()
	return v
}

// 升级版本
var cache = struct {
	sync.Mutex  // 直接匿名嵌入
	mapping map[string]string
}{
	mapping: make(map[string]string),
}
	// 现在这个cache就可以直接当作一个整体来使用了
func Loopup2(key string) string {
	// 这个cache里面就有很多玩意
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}