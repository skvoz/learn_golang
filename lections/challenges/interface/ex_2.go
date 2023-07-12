package main

import "fmt"

func main() {
	var empty interface{}
	empty = 5
	fmt.Println("int value", empty)

	empty = 11.2
	fmt.Println("int value", empty)

	empty = []int{1, 2, 3}
	fmt.Println("int value", empty)

	empty = append(empty.([]int), 4, 5, 6)
	fmt.Println("int value", empty)
}
