package main

import (
	"fmt"
)

func main() {
	s := "你好 Go!"

	b := []rune(s)
	fmt.Printf("%#v\n", b)
	for i, v := range b {
		fmt.Printf("%d -> %c\n", i, v)
	}
}
