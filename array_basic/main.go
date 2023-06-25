package main

import "fmt"

func main() {
	var numbers [4]int

	fmt.Printf("%#v\n", numbers)

	var a2 = [3]int{-10, 1, 100}

	fmt.Printf("%#v\n", a2)

	var a4 = [4]string{"x", "y"}

	fmt.Printf("%#v\n", a4)

	var a5 = [...]int{1, 1, 1, 1, 1, 1}
	// var a5 = []int{1, 1, 1, 1, 1, 1} //thie save as [...] ?
	fmt.Printf("%#v\n", a5)
	fmt.Printf("the length of a5 %d\n", len(a5))
}
