package main

import "fmt"

func main() {

	//predeclared built-in types
	//int float, byte, run bool
	//RUNE TYPE
	var r rune = 'f'
	fmt.Printf("%T\n", r) //return int32, go don't have datatype character
	//so rune alias for int32

	fmt.Println(r) //102

	//STRING
	//sequences of bytes
	var s string = "Hello GO!"

	fmt.Printf("%T\n", s)

	//composite types
	//array and slice

	//array type - array with fixed length
	var numbers = [4]int{4, 5, -9, 101}

	fmt.Printf("%T\n", numbers) //[4]int

	//slice - array with dynamic length
	var cities = []string{"London", "Tokyo", "new York"}

	fmt.Printf("%T\n", cities) //[]string

	//map

	balances := map[string]float64{
		"USD": 2332.2,
		"EUR": 511.1,
	}

	fmt.Printf("%T\n", balances) // map[string]float64

	//struct
	type Person struct {
		name string
		age  int
	}

	var you Person

	fmt.Printf("%T\n", you) //main.Person

	//pointer
	var x int = 2

	ptr := &x

	fmt.Printf("%T %v \n", ptr, ptr) //*int 0xc0000180a8

	//function type
	fmt.Printf("%T \n", f) //func()

}

func f() {

}
