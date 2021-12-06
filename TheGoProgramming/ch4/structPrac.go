package main

import (
	"fmt"
	"time"
)
type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}
func f1() {


	var dilbert Employee

	// 1 赋值
	dilbert.Salary -= 500

	// 使用指针接收字段
	position := &dilbert.Address
	*position += "潮阳区"

	// 使用指针接收一整个struct
	var otherEmployee *Employee = & dilbert
	otherEmployee.Name = "cindy"  // 注意这里等价于 (*otherEmployee).Name = "cindy"  它也会改变原来dilbert的值

	fmt.Println(dilbert)
	fmt.Println(otherEmployee)

	// 2 EmployeeByID 返回值的是指针的情况下
	fmt.Println(EmployeeByID(1).ID)
	//EmployeeByID(10).Salary = 1  // 如果返回值的是一个正常值，而不是指针的话，就会报错
}

//func EmployeeByID(id int)  Employee{
func EmployeeByID(id int)  *Employee{
	var kim Employee
	Jine := &kim
	//return kim
	return Jine
}

type Point struct {
	X, Y int
}
type Circle struct{
	Point  // 匿名内部成员
	Radius int
}
type Wheel struct {
	Circle
	Spokes int
}

func f2() {
	// 结构体匿名成员
	var w Wheel
	w.X = 10  // 可以直接访问X。相当于w.Circle.Point.X
	fmt.Println(w, w.Circle.Point.X, w.X)

	// 如果要字面量赋值，有两种方式
	// way1
	way1 := Wheel{Circle{Point{1, 2},10}, 10}
	// way2
	way2 := Wheel{Spokes: 10,
		Circle:Circle{Radius: 20, Point:Point{1, 2}}}
	fmt.Println(way1, way2)
	fmt.Printf("%#v\n", way1)
}



func runStruct() {
	//f1()
	f2()
}
