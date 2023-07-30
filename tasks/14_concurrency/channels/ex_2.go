package main

import (
	"fmt"
)

func main() {
	c := make(chan string)
	go func(s string) {
		c <- s
	}("foo")

	v := <-c
	fmt.Println("main goroutine receiving data from channel:", v)
}
