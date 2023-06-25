package main

import "fmt"

type person struct {
	name   string
	age    int
	colors []string
}

func main() {

	me := person{name: "Petr", age: 10, colors: []string{"blue", "red"}}
	you := person{name: "Vasa", age: 11, colors: []string{"green", "yellow"}}

	fmt.Printf("%+v\n%+v\n", me, you)

	me.name = "Andrei"
	youFavourColor := you.colors
	fmt.Printf("%#v\n", youFavourColor)

	you.colors = append(you.colors, "black")
	fmt.Printf("%#v\n", you)
}
