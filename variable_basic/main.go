package main

import "fmt"

func main() {
	var x int = 1 //first approach to declaration
	y := 1        //second approach to declaration, use when we now value

	fmt.Printf("first approach %d, second approache %d", x, y)
	y = 2 //set new variables value
	fmt.Printf("first approach %d, second approache %d", x, y)

	foo, bar := "foo", 100 //multiple declaration
	_, _ = foo, bar

	car, cost := "audi", 5000
	car, year := "bmw", 1000 //haha, car redeclaration

	fmt.Println(car, cost, year)

	var opened = false

	opened, file := true, "a.txt" // one more example redeclaration

	_, _ = opened, file

	var (
		one   float64
		two   string
		three bool
	)
	fmt.Println(one, two, three) //will output 0, "", false

	var i, j int
	i, j = 5, 8
	i, j = j, i
	fmt.Println(i, j) //return 8,5

	var x1 = 10

	fmt.Println(x1)
}
