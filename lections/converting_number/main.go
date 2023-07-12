package main

import "fmt"

func main() {
	var x = 3
	var y = 3.1

	// x = x * y //imposible operation strong typisation
	x = x * int(y) //9

	fmt.Println(x)
	fmt.Printf("%T\n", y) //float64

	x = int(float64(x) * y)
	fmt.Println(x) //27

	y = float64(x) * y
	fmt.Println(y) //83.7

	var a = 5 //int type

	fmt.Printf("%T\n", a) //int

	var b int64 = 2
	// a=b //return error  beacose diff types
	_ = b
}
