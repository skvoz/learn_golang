package main

import "fmt"

type emptyInterface interface {
}
type person struct {
	info interface{}
}

func main() {
	var empty interface{}

	empty = 5
	fmt.Println(empty) //5
	empty = "GO"
	fmt.Println(empty) //GO
	empty = []int{4, 5, 6}
	fmt.Println(empty)
	// fmt.Println(len(empty))//error len can't evalueate empty interface
	fmt.Println(len(empty.([]int))) //so we're using type assertion
	you := person{}
	you.info = "your name"
	fmt.Println(you.info) //your name
	you.info = 40
	fmt.Println(you.info) //40
	you.info = []int{1, 2, 3}
	fmt.Println(you.info) //[1 2 3]

}
