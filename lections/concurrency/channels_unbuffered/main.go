/////////////////////////////////
// Unbuffered Channels
// Go Playground: https://play.golang.org/p/_44csjQDJvM
/////////////////////////////////

package main

import (
	"fmt"
	"time"
)

func main() {
	// Declaring a buffered channel.
	c := make(chan int, 3)                     //buffered channel
	fmt.Println("Channel's capacity:", cap(c)) // => 3
	go func(c chan int) {
		for i := 1; i <= 5; i++ {
			fmt.Printf("func goroutine #%d starts sending data into the channel\n", i)
			c <- 10
			fmt.Printf("func goroutine #%d after sending data into the channel\n", i)
		}
		// closing the buffered channel.
		close(c)
	}(c) // calling the anonymous func and passing c1 as argument

	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	for v := range c { //v:=<-c
		fmt.Println("main goroutine receiving data from channel:", v)

	}
	// After running the program  we notice that the goroutines start sending data
	// into the channel BEFORE the main goroutine had a chance
	// to receive data from the channel.

	// The sender of this buffered channel will block only when there is no empty slot in the channel, in this
	// case after 3 writing attempts because the channel has a capacity of 3.
	// The receiver will block on the channel when it's empty.

	// A receive operation on a closed channel will proceed without blocking
	// and yield the zero-value for the type that is sent through the channel.
	fmt.Println(<-c)
	// c <- 10 //error panic: send on closed channel
}
