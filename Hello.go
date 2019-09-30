package main

import "fmt"

func main() {
	var sl1 = make([]int, 5)          // 定义元素个数为5的切片
	sl2 := make([]int, 5, 10)         // 定义元素个数5的切片，并预留10个元素的存储空间（预留空间不知道有什么用？）
	sl3 := []string{`aa`, `bb`, `cc`} // 直接创建并初始化包含3个元素的数组切片
	fmt.Println(sl1, len(sl1))        // 输出：[0 0 0 0 0] 5
	fmt.Println(sl2, len(sl2))        // 输出：[0 0 0 0 0] 5
	fmt.Println(sl3, len(sl3))        // [aa bb cc] 3
	sl1[1] = 1                        // 声明或初始化大小中的数据，可以指定赋值
	sl1[4] = 4
	//sl1[5] = 5 // 编译报错，超出定义大小
	sl1 = append(sl1, 5)       // 可以追加元素
	fmt.Println(sl1, len(sl1)) // 输出：[0 1 0 0 4 5] 6
	sl2[1] = 1
	sl2 = append(sl2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	fmt.Println(sl2, len(sl2)) // 输出：[0 1 0 0 0 1 2 3 4 5 6 7 8 9 10 11] 16
	// 遍历切片
	for i := 0; i < len(sl2); i++ {
		v := sl2[i]
		fmt.Printf("i: %d, value:%d \n", i, v)
	}
}
