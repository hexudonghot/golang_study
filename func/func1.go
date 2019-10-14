package main

import "fmt"

func main() {
	//s := square
	//普通函数
	//fmt.Println(s(1))
	t := incr()
	//闭包
	//fmt.Println(t)
	fmt.Println(t())
	fmt.Println(t())
	fmt.Println(t())
}
func square(x int) int {
	return x
}

func incr() func() int {
	x := 0
	fmt.Println("bbbbb")
	return func() int {
		fmt.Println("aaa")
		x++
		return x
	}
}
