package main

import "fmt"

type person struct {
	name   string
	age    int
	colors []string
}

func main() {

	me := person{name: "Petr", age: 10, colors: []string{"blue", "red"}}

	for _, v := range me.colors {
		fmt.Println(v)
	}
}
