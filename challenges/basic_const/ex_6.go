package main

import "fmt"

func main() {
	const (
		jun = iota + 6
		jul
		aug
	)

	fmt.Print(jun, jul, aug)
}
