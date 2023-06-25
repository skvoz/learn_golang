package main

import "fmt"

func main() {
	grades := [3]int{
		1: 10,
		0: 5,
		2: 7,
	}

	fmt.Println(grades)

	accounts := [3]int{2: 50}
	fmt.Println(accounts)

	names := [...]string{5: "Dan"}
	fmt.Println(names) //[     Dan], first 5 el is space

	cities := [...]string{
		5:        "Paris",
		"London", // this is at index 6
		1:        "NYC",
	}
	fmt.Printf("%#v\n", cities) // -> [7]string{"", "NYC", "", "", "", "Paris", "London"}

	weekend := [7]bool{5: true, 6: true}
	fmt.Println(weekend) // => [false false false false false true true]
}
