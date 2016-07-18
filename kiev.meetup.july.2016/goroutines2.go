package main

import "fmt"

var a string
var c = make(chan struct{})

func f() {
	a = "hello"
	c <- struct{}{}
}

func main() {
	go f()
	<-c
	fmt.Print(a)
}
