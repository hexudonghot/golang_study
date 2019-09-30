package main

import "fmt"

func main() {
	//var sl []int             // 声明一个切片
	//sl = append(sl, 1, 2, 3) // 往切片中追加值
	//fmt.Println(sl)          // 输出：[1 2 3]
	//
	//var arr = [5]int{1, 2, 3, 4, 5} // 初始化一个数组
	//var sl1 = arr[0:2]              // 冒号:左边为起始位（包含起始位数据），右边为结束位（不包含结束位数据）；不填则默认为头或尾
	//var sl2 = arr[3:]
	//var sl3 = arr[:5]
	//
	//fmt.Println(sl1) // 输出：[1 2]
	//fmt.Println(sl2) // 输出：[4 5]
	//fmt.Println(sl3) // 输出：[1 2 3 4 5]
	//
	//sl1 = append(sl1, 11, 22) // 追加元素
	//fmt.Println(sl1)          // 输出：[1 2 11 22]

	var sl1 = make([]int, 5)          // 定义元素个数为5的切片
	sl2 := make([]int, 5, 10)         // 定义元素个数5的切片，并预留10个元素的存储空间（预留空间不知道有什么用？）
	sl3 := []string{`aa`, `bb`, `cc`} // 直接创建并初始化包含3个元素的数组切片

	var intArr = [2]int{1, 2}
	strArr := [3]string{`aa`, `bb`, `cc`} // 定义数组 需要指定长度是和 slice的区别
	fmt.Println(intArr, strArr)

	fmt.Println(sl1, len(sl1)) // 输出：[0 0 0 0 0] 5
	fmt.Println(sl2, len(sl2)) // 输出：[0 0 0 0 0] 5
	fmt.Println(sl3, len(sl3)) // [aa bb cc] 3

}
