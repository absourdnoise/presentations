package main

import "fmt"

func main() {
	var a string
	c := make(chan struct{})
	go func() {
		a = "hello"
		c <- struct{}{}
	}()
	<-c
	fmt.Print(a)
}
