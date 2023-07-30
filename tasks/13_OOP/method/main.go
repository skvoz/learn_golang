package main

import (
	"fmt"
)

type money float64
type book struct {
	price float64
	title string
}

func (b *book) vat() float64 {
	return b.price * 0.09
}

func (b *book) discount() {
	(*b).price = b.price * 0.9
}

//composition method with type
func (m money) print() {
	fmt.Printf("%.2f\n", m)
}

func (m money) printStr() string {
	return fmt.Sprintf("%.2f", m)

}

func main() {
	fmt.Println("Ex1###########################################")
	var salary money = 1.5 * 5.503
	salary.print()

	fmt.Println("Ex2###########################################")
	s := salary.printStr()
	fmt.Println(s)

	fmt.Println("Ex3###########################################")
	bestBook := book{title: "The Trial", price: 9.9}

	// calling the methods
	vat := bestBook.vat()
	fmt.Printf("Vat:%v\n", vat)

	bestBook.discount()
	fmt.Printf("%#v\n", bestBook)
	/**
	  Ex1###########################################
	  8.25
	  Ex2###########################################
	  8.25
	  Ex3###########################################
	  Vat:0.891
	  main.book{price:8.91, title:"The Trial"}
	  **/
}
