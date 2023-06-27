package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	for i := 1; i <= 50; i++ {
		go func(x int) {
			c <- x * x
		}(i)
	}

	for i := 1; i <= 50; i++ {
		fmt.Println(<-c)
	}
}
