package main

import "fmt"

func swap(x *float64, y *float64) {
	y1 := *y
	x1 := *x
	*x = y1
	*y = x1
	//in course more short
	//*a, *b = *b, *a
}
func main() {
	x, y := 5.5, 8.8
	fmt.Printf("x: %f, y:%f\n", x, y)
	swap(&x, &y)
	fmt.Printf("x: %f, y:%f\n", x, y)
}
