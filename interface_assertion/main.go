/////////////////////////////////
// Implementing Interfaces in Go
// Go Playground: https://play.golang.org/p/SMjFrOYL5f3
/////////////////////////////////

package main

import (
	"fmt"
	"math"
)

// declaring an interface type called shape
// an interface contains only the signatures of the methods, but not their implementation
type shape interface {
	area() float64
	perimeter() float64
}

// declaring 2 struct types that represent geometrical shapes: rectangle and circle

type rectangle struct {
	width, height float64
}
type circle struct {
	radius float64
}

// method that calculates circle's area
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// method that calculates rectangle's area
func (r rectangle) area() float64 {
	return r.height * r.width
}

// method that calculates circle's perimeter
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

// method that calculates rectangle's perimeter
func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

// any type that implements the interface is also of type of the interface
// rectangle and circle values are also of type shape
func print(s shape) {
	fmt.Printf("Shape: %#v\n", s)
	fmt.Printf("Area: %v\n", s.area())
	fmt.Printf("Perimeter: %v\n", s.perimeter())
}

func main() {
	//here we can use only method assertion in shape
	//even if circle has more method
	var s shape = circle{radius: 2.5}
	fmt.Printf("%T\n", s)
	// s.volume() //error
	//<interface>.(<type assertion>).method
	//s.(circle).volume()//correct  because we have mthod
	//associated with circle
	// s.(circle).volume()//correct example
	ball, ok := s.(circle)

	if ok == true {
		fmt.Printf("Ball Volume  : %v\n", ball.volume()) //Ball Volume  : 49.087385212340514

	}

	switch value := s.(type) {
	case circle:
		fmt.Printf("%#v has circle type \n", value) //correct, main.circle{radius:2.5} has circle type
	case rectangle:
		fmt.Printf("%#v has rectangle type \n", value)
	}
	// s = rectangle{width: 3.4, height: 2.2}

	// switch value := s.(type) {
	// case circle:
	// 	fmt.Printf("%#v has circle type \n", value)
	// case rectangle:
	// 	fmt.Printf("%#v has rectangle type \n", value)  //main.rectangle{width:3.4, height:2.2} has rectangle type

	// }

}
