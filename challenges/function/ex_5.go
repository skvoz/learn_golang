package main

import (
	"fmt"
)

func sum(i ...int) (s int) {

	for _, v := range i {
		s += v
	}
	return
}

func main() {
	fmt.Println(sum(1, 2, 3)) //return 6
}
