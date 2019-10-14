package main

import "fmt"

type user struct {
	name string
	age  byte
}

func main() {
	s1 := make([]int, 5) // 定义元素个数为5的切片
	s1 = append(s1, 1)
	s1 = append(s1, 1)
	s1 = append(s1, 1)
	s1 = append(s1, 1)
	fmt.Println(s1)
	d := [...]user{{name: "xx", age: 10}, {name: "hxd1", age: 10}}
	fmt.Println(d)
}
