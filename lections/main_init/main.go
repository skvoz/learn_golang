package main

import "fmt"

//declare array

var numbers [10]int
//magic )
func init()  {
	fmt.Println("this is init 1")
}

//magic )
//use this function like decorator
func init()  {
	fmt.Println("this is init 2")
	for i := 0; i < len(numbers); i++ {
		numbers[i]  = i*2
	}
}

func main()  {
	fmt.Println("this is main")
	// init() //error undefinded init
	fmt.Printf("%#v\n", numbers)
}

/**
this is init 1
this is init 2
this is main
[10]int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18}
**/
