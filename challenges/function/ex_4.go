package main

import (
	"fmt"
)

func sum(i ...int) int {
	s := 0
	for _, v := range i {
		s += v
	}
	return s
}

func main() {
	fmt.Println(sum(1, 2, 3)) //return 6
}
