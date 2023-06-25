package main

import "fmt"

func main() {
	numbers := []int{2, 3}
	numbers = append(numbers, 10, 20, 30)
	fmt.Println(numbers)

	n := []int{100, 200}
	// numbers = append(numbers, n) //error
	numbers = append(numbers, n...) //ellipce oeraptor add all elements from slice to anohter slice
	fmt.Println(numbers)

	src := []int{10, 20, 30}
	dst := make([]int, len(src))
	nn := copy(dst, src)
	fmt.Printf("is value dst after copy %#v , nn is %#v", dst, nn) //return 3 , copy return number copied elements
}
