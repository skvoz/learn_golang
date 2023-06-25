package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type object interface {
	volume() float64
}

type geometry interface {
	shape
	object
	getColor() string
}

type cube struct {
	edge  float64
	color string
}

func (c cube) area() float64 {
	return 6 * (c.edge * c.edge)
}

func (c cube) getColor() string {
	return c.color
}

func (c cube) volume() float64 {
	return math.Pow(c.edge, 3)
}
func measure(g geometry) (float64, float64) {
	a := g.area()
	v := g.volume()
	return a, v
}

func main() {
	//exists circule-embedding interface
	//int1->int2->int2
	//will genreate error
	c := cube{edge: 2}
	a, v := measure(c)
	fmt.Printf("area :%v, volume %v\n", a, v)
}
