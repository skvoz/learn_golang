package main

import "fmt"

func swap(a, b *float64) {
	*a, *b = *b, *a
}

func main() {
	fmt.Println("Ex1#################################")
	x := 10.10
	fmt.Println("address of x:=10.10 is", &x)

	ptr := &x
	fmt.Printf("type ptr=%T, value ptr=%v\n", ptr, ptr)

	fmt.Printf("The address of ptr: %p\n", &ptr)
	fmt.Printf("The value of x through the pointer:%f\n", *ptr)

	fmt.Println("Ex2#################################")

	x1, y1 := 10, 2
	ptrx1, ptry1 := &x1, &y1

	z := *ptrx1 / *ptry1
	fmt.Printf("x pointer / y pointer = %v\n", z)

	fmt.Println("Ex3#################################")
	x2, y2 := 5.5, 8.8
	fmt.Printf("x=%v, y=%v\n", x2, y2)
	swap(&x2, &y2)
	fmt.Printf("x=%v, y=%v\n", x2, y2)

}
