package main

import "fmt"

func main() {
	var a string
	go func() { a = "hello" }()
	fmt.Print(a)
}
