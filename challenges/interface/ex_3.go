package main

import "fmt"

type cube struct {
	edge float64
}

func volume(c cube) float64 {
	return c.edge * c.edge * c.edge
}

func main() {
	//:) my solution
	// x := cube{edge: 5}
	// v := volume(x)
	// fmt.Printf("Cube Volume: %v\n", v) //return 125

	//correct solution
	var x interface{}
	x = cube{edge: 5}

	// Type Assertion
	v := volume(x.(cube))

	fmt.Printf("Cube Volume: %v\n", v) //return 125

}
