package main

import (
	"fmt"
	"reflect"
)

type user struct {
	name string
	age int
}

func run14() {
	u := user{name: "bourne", age: 19}
	t := reflect.TypeOf(u)

	fmt.Println(t)
}