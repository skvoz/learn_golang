package main

import "fmt"

func main() {
	nums := [...]int{30, -1, -6, 90, -6}
	var counts int
	for _, v := range nums {
		if v > 0 {
			counts++
		}
	}

	fmt.Println(counts) //2
}
