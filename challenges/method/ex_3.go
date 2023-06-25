package main

import "fmt"

type book struct {
	price float64
	title string
}

func (b book) vat() float64 {
	return b.price * 0.09
}

func (b *book) discount() {
	(*b).price = b.price * 0.9
}

func main() {
	book := book{title: "title1", price: 100}
	vat := book.vat()
	book.discount()
	fmt.Printf("Vat : %v, discount: %v", vat, book)
}
