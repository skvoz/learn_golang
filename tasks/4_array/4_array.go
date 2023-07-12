package main

import "fmt"

func main() {
	fmt.Println("Ex1#######################")
	var cities [2]string
	fmt.Printf("%#v\n", cities)

	grades := [3]float64{1.1, 2.2}
	fmt.Printf("%#v\n", grades)

	tasdDone := [...]bool{}
	fmt.Printf("%#v\n", tasdDone)

	cities = [2]string{"city1", "city2"}
	for i := 0; i < len(cities); i++ {
		fmt.Print(cities[i], " ")
	}
	fmt.Println("")

	for _, grade := range grades {
		fmt.Print(grade, " ")
	}
	fmt.Println("")
	/**
	  [2]string{"", ""}
	  [3]float64{1.1, 2.2, 0}
	  [0]bool{}
	  city1 city2
	  1.1 2.2 0
	  **/
	fmt.Println("Ex2#######################")

	nums := [...]int{30, -1, -6, 90, -6}

	var count int
	for _, v := range nums {
		if v%2 == 0 && v > 0 {
			count++
		}
	}
	fmt.Println(count)
}
