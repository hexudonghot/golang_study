package main

import "fmt"

func main() {
	var p *int // 定义指向int型的指针，默认值为空：nil

	// nil指针不指向任何有效存储地址，操作系统默认不能访问
	//fmt.Printf("%x\n", *p) // 编译报错

	var a int = 10
	p = &a        // 取地址
	add := a + *p // 取值

	fmt.Println(a)   // 输出：10
	fmt.Println(p)   // 输出：0xc0420080b8
	fmt.Println(add) // 输出：20
}
