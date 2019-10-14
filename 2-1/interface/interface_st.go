package _interface

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokia NokiaPhone) call() {
	fmt.Println("I am Phone, I can call you !")
}
