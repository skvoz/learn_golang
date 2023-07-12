package main

import "fmt"

func main() {
	var cities [2]string

	fmt.Printf("%#v\n", cities)

	grades := [3]float64{0.1, 0.2}
	fmt.Printf("%#v\n", grades)

	taskDone := [...]bool{true, true, false}
	fmt.Printf("%#v\n", taskDone)

	for i := 0; i < len(cities); i++ {
		fmt.Printf("%s\n", cities[i])
	}

	for _, v := range grades {
		fmt.Printf("%#v\n", v)
	}
}
