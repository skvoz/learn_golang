package main

import "fmt"

func main() {
	// 1.
	var empty interface{}
	fmt.Printf("%T\n", empty)

	// 2.
	empty = 5
	fmt.Printf("%T\n", empty)

	// 3.
	empty = 10.10
	fmt.Printf("%T\n", empty)

	// 4.
	empty = []int{1, 2, 3}
	fmt.Printf("%T\n", empty)

	// 5. Type Assertion
	empty = append(empty.([]int), 10)

	// 6.
	fmt.Printf("%v\n", empty)

}
