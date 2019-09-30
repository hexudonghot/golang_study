package main

import (
	"fmt"
	"strconv"
)

// 定义一个 Person 接口
type Person interface {
	Say(s string) string
	Walk(s string) string
}

// 定义一个 Man 结构体
type Man struct {
	Name string
	Age  int
}

// Man 实现 Say 方法
func (m Man) Say(s string) string {
	return s + ", my name is " + m.Name
}

// Man 实现 Walk 方法，strconv.Itoa() 数字转字符串
func (m Man) Walk(s string) string {
	return "Age: " + strconv.Itoa(m.Age) + " and " + s
}

func main() {
	var m Man       // 声明一个类型为 Man 的变量
	m.Name = "Mike" // 赋值
	m.Age = 30
	fmt.Println(m.Say("hello"))    // 输出：hello, my name is Mike
	fmt.Println(m.Walk("go work")) // 输出：Age: 30 and go work

	jack := Man{Name: "jack", Age: 25} // 初始化一个 Man 类型数据
	fmt.Println(jack.Age)
	fmt.Println(jack.Say("hi")) // 输出：hi, my name is jack
}
