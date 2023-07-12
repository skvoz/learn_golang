package main

import "fmt"

func main() {
	var i int = 3
	var f float64 = 3.2

	fmt.Printf("type i is %T, type f is %T\n", float64(i), int(f))
}
