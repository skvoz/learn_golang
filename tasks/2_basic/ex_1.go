package main

import "fmt"

func main() {
	//long declaration
	var a int
	var b float64
	var c bool
	var d string

	//blank identifier
	// _, _, _, _ = a, b, c, d
	//short
	// var (
	// 	a int
	// 	b float64
	// 	c bool
	// 	d string
	// )

	fmt.Printf("%T, %T, %T, %T, \n", a, b, c, d)
	//short
	z := -20
	fmt.Printf("%T %#v\n", z, z)

}
