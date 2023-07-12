package main

import "fmt"

func main() {
	var n []int
	fmt.Println(n == nil)

	a, b := []int{1, 2, 3}, []int{1, 2, 3}
	// fmt.Println(a == b) //error only compare with nil

	var eq bool = true
	a = nil
	for i, valueA := range a {
		if valueA != b[i] {
			eq = false
			break
		}
	}

	if len(a) != len(b) {
		eq = false
	}
	if eq {
		fmt.Println("a and b slice are equal")
	} else {
		fmt.Println("a and b slice are not equal")

	}
}
