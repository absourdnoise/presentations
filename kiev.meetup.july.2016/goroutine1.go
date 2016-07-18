package main

import "fmt"

var a string

func f() {
	a = "hello"
}

func main() {
	go f()
	fmt.Print(a)
}
