package main

import "fmt"

type person struct {
	name   string
	age    int
	colors []string
	gr     grades
}
type grades struct {
	grade  int
	course string
}

func main() {
	fmt.Println("Ex1,2############################")
	me := person{name: "foo", age: 10, colors: []string{"red", "blue"}, gr: grades{grade: 10, course: "foo"}}
	you := person{name: "bar", age: 20, colors: []string{"green", "red"}, gr: grades{grade: 12, course: "language"}}

	me.name = "Andei"
	you.colors = []string{"one", "two"}
	fmt.Printf("me: %#v\nyou: %#v", me, you)

	// 2.
	var colors []string = you.colors
	fmt.Printf("Type: %T, Value: %v\n", colors, colors)

	colors = append(colors, "green")
	you.colors = colors

	fmt.Printf("%v\n", me)
	fmt.Printf("%+v\n", you)

	fmt.Println("Ex3############################")
	for _, v := range me.colors {
		fmt.Printf("%#v", v)
	}
	fmt.Println()
	fmt.Println("Ex4############################")
	me.gr.course = "Golang"
	me.gr.grade = 20
	fmt.Printf("%#v", me)
	/**
	me: main.person{name:"Andei", age:10, colors:[]string{"red", "blue"}, gr:main.grades{grade:10, course:"foo"}}
	you: main.person{name:"bar", age:20, colors:[]string{"one", "two"}, gr:main.grades{grade:12, course:"language"}}Type: []string, Value: [one two]
	{Andei 10 [red blue] {10 foo}}
	{name:bar age:20 colors:[one two green] gr:{grade:12 course:language}}
	Ex3############################
	"red""blue"
	Ex4############################
	main.person{name:"Andei", age:10, colors:[]string{"red", "blue"}, gr:main.grades{grade:20, course:"Golang"}}
	**/

}
