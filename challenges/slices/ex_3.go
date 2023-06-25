package main

import "fmt"

func main() {
	nums := []float64{0.1, 0.2, 0.3}

	nums = append(nums, 10.1)
	nums = append(nums, 4.1, 5.5, 6.6)
	fmt.Println(nums)

	n := []float64{7.7, 8.8}

	nums = append(nums, n...)
	fmt.Println(nums)
}
