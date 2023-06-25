package main

import "fmt"

type foo uint
type bar foo

func main() {
	//aliases (type a1 = a2 ) like underline types but we don't need convert data
	var a uint8 = 10
	var b byte

	b = a

	_ = b

	type second = uint

	var hour second = 3600

	fmt.Printf("hour type: %T\n", hour)             // => hour type: uint
	fmt.Printf("minutes in a hour: %d \n", hour/60) //no error, wow

	var d bar
	fmt.Printf("%T\n", d)
}
