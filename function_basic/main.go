package main

import (
	"fmt"
	"math"
)

func main() {
	//not exist overload
	//not exist default params
	f1()                            //this is f1() function
	f2(1, 2)                        //sum: 3
	f3(4, 5, 6, 4.4, 7.8, "golang") //4 5 6 4.4 7.8 golang
	fmt.Println(f4(4))              //256
	fmt.Println(f5(4, 2))           //6 8
	fmt.Println(sum(4, 8))          //12

}

func f1() {
	fmt.Println("this is f1() function")
}

func f2(a int, b int) {
	fmt.Println("sum:", a+b)
}

func f3(a, b, c int, d, e float64, s string) {
	fmt.Println(a, b, c, d, e, s)
}

func f4(a float64) float64 {
	return math.Pow(a, a)
}

func f5(a, b int) (int, int) {
	return a + b, a * b
}

//predefined output
func sum(a, b int) (s int) {
	fmt.Println(s) // s-> 0
	s = a + b
	return
}
