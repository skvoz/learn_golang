package main

import "fmt"

type money float64

func (m money) print() {
	fmt.Printf("%.2f\n", m)
}

func main() {
	var cost money = 1.2345
	fmt.Println(cost) //1.2345

	cost.print() //1.23
}
