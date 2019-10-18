package main

import "fmt"

func main() {
	//s:="abc"
	//for i:=range s{
	//	fmt.Println(s[i])
	//}
	//for _,c:=range  s{
	//	fmt.Println(c)
	//}
	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
