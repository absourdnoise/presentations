package main

import (
	"fmt"
)

var c = make(chan int)
var a string

func f() {
	a = "hello, world"
	<-c
}
func main() {
	go f()
	c <- 0
	fmt.Println(a)
}
