package main

import (
	"fmt"
)

//purpouse from course author
func main() {

	nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}
	sum := 0
	for _, v := range nums[2 : len(nums)-2] {
		sum += v
	}

	fmt.Println(sum)
}
