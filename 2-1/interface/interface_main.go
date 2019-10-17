package main

import "fmt"

//  https://juejin.im/post/5c9d5631f265da60d82dde9c
type Shape interface {
	Area() float32
}

type Rect struct {
	width  float32
	height float32
}

func (r Rect) Area() float32 {
	return r.width * r.height
}

func main() {
	var s Shape
	s = Rect{5.0, 4.0}
	//r := Rect{5.0, 4.0}
	//fmt.Printf("type of s is %T\n", s)
	//fmt.Printf("value of s is %v\n", s)
	fmt.Println("area of rectange s", s.Area())
	//fmt.Println("s == r is", s == r)
}
