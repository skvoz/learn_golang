package main

import "fmt"

func main() {
	birthYear, currentYear := 1980, 2020

	for i := birthYear; i <= currentYear; {
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Println()
}
