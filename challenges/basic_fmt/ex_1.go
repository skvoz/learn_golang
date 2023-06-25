package main

import "fmt"

func main() {
	//this ex difficult for understanding
	x, y, z := 10, 15.5, "Gophers"
	score := []int{10, 20, 30}

	fmt.Printf("x is %d, y is %f, z is %s\n", x, y, z)
	fmt.Printf("score is %#v\n", score)

	fmt.Printf("\"Gophers\"\n")
	fmt.Printf("x is %v, y is %v, z is %v, score is  %v\n", x, y, z, score)
	fmt.Printf("type y is %T, type score is %T\n", y, score)
}
