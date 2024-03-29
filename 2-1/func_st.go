package main

import "fmt"

func main() {
	/* 定义局部变量 */
	var a = 100
	var b = 200
	var ret int

	/* 调用函数并返回最大值 */
	ret = max(a, b)

	fmt.Printf("最大值是 : %d\n", ret)

	//函数返回多个值
	a1, b11 := swap("Google", "Runoob")
	fmt.Println(a1, b11)
}

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	/* 定义局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func swap(x, y string) (string, string) {
	return y, x
}
