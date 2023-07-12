package main

import (
	"fmt"
	"time"
)

func main() {
	// var c chan string
	c := make(chan string)
	go func(s string) {
		c <- s
		time.Sleep(time.Second * 2)
	}("foo bar")
	fmt.Println(<-c)
}
