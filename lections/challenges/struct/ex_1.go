package main

import "fmt"

//cource solution
// type person struct {
// 	name   string
// 	age    int
// 	colors []string
// }

func main() {
	//in course colution out in header
	type person struct {
		name   string
		age    int
		colors []string
	}

	me := person{name: "Petr", age: 10, colors: []string{"blue", "red"}}
	you := person{name: "Vasa", age: 11, colors: []string{"green", "yellow"}}

	fmt.Printf("%+v\n%+v\n", me, you)
}
