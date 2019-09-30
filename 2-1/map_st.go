package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// 声明后赋值
	var m map[int]string
	fmt.Println(m) // 输出空的map：map[]
	//m[1] = `aa`    // 向未初始化的map中赋值报错：panic: assignment to entry in nil map

	// 声明并初始化，初始化使用{} 或 make 函数（创建类型并分配空间）
	var m1 = map[string]int{}
	var m2 = make(map[string]int)
	m1[`a`] = 11
	m2[`b`] = 22
	fmt.Println(m1) // 输出：map[a:11]
	fmt.Println(m2) // 输出：map[b:22]

	// 初始化多个值
	var m3 = map[string]string{"a": "aaa", "b": "bbb"}
	m3["c"] = "ccc"
	fmt.Println(m3) // 输出：map[a:aaa b:bbb c:ccc]

	// 删除 map 中的值
	delete(m3, "a") // 删除键 a 对应的值
	fmt.Println(m3) // 输出：map[b:bbb c:ccc]

	// 查找 map 中的元素
	v, ok := m3["b"]
	if ok {
		fmt.Println(ok)
		fmt.Println("m3中b的值为：", v) // 输出：m3中b的值为： bbb
	}
	// 或者
	if v, ok := m3["b"]; ok { // 流程处理后面讲
		fmt.Println("m3中b的值为：", v) // 输出：m3中b的值为： bbb
	}

	fmt.Println(m3["c"]) // 直接取值，输出：ccc

	// map 中的值可以是任意类型
	m4 := make(map[string][5]int)
	m4["a"] = [5]int{1, 2, 3, 4, 5}
	m4["b"] = [5]int{11, 22, 33}
	fmt.Println(m4)                // 输出：map[a:[1 2 3 4 5] b:[11 22 33 0 0]]
	fmt.Println(unsafe.Sizeof(m4)) // 输出：8，为8个字节，map其实是个指针，指向某个内存空间
}
