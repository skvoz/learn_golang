package main

import (
	"fmt"
)

func power(i int, c chan int) {
	x := i * i
	c <- x
}

func main() {
	c := make(chan int)
	/**
	  my solution
	  	for i := 1; i <= 50; i++ {
	  		go power(i, c)
	  		fmt.Println(<-c)
	  	}
	  **/
	//source  solution
	for i := 1; i <= 50; i++ {
		go power(i, c)
	}

	for i := 1; i <= 50; i++ {
		fmt.Println(<-c)
	}
}
