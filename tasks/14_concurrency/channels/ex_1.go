package main

import "fmt"

func main() {
	var c1 chan float64
	c2 := make(<-chan rune)
	c3 := make(chan<- rune)
	c4 := make(chan int, 10)

	fmt.Printf("c1 type is :%#v\nc2 type is :%#v\nc3 type is :%#v\nc4 type is :%#v\n",
		c1, c2, c3, c4)
}
