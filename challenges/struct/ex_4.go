package main

import "fmt"

type person struct {
	name    string
	age     int
	colors  []string
	mygrade grades
}

type grades struct {
	grade  string
	course int
}

func main() {

	me := person{name: "Petr", age: 10,
		colors:  []string{"blue", "red"},
		mygrade: grades{grade: "grade1", course: 1}}

	you := person{name: "Vasa", age: 11, colors: []string{"green", "yellow"},
		mygrade: grades{grade: "grade2", course: 2}}

	me.mygrade.grade = "Golang"
	me.mygrade.course = 3

	fmt.Printf("%+v\n%+v\n", me, you)
}
