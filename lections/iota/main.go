package main

import "fmt"

func main() {
	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)

	fmt.Println(c1, c2, c3) //0,1,2

	const (
		c11 = iota
		c22
		c33
	)

	fmt.Println(c11, c22, c33) //0,1,2
	const (
		a = iota * 2 //0
		b            //2
		c            //4
	)

	fmt.Println(a, b, c) //0,2,4

	const foo int = 1

}
