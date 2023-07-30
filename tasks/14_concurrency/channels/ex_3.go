package main

import (
	"fmt"
)

func power(p int, c chan int) {
	c <- p * p
}

func main() {
	ch := make(chan int)
	for i := 1; i <= 50; i++ {
		go power(i, ch)
	}

	for i := 1; i <= 50; i++ {
		data := <-ch
		fmt.Printf("обработчик получил %#v\n", data)

	}
}
