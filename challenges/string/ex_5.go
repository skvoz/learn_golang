package main

import (
	"fmt"
)

func main() {
	s := "你好 Go!"

	b := []byte(s)
	fmt.Printf("%#v\n", b)
	for i, v := range b {
		fmt.Printf("%d -> %d\n", i, v)
	}
}
