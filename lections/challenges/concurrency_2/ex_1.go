package main

import "fmt"

func main() {
	var c1 chan float64

	c2 := make(chan<- rune)
	c3 := make(<-chan rune)
	c4 := make(chan rune, 10)
	fmt.Println(c1, c2, c3, c4)
}
