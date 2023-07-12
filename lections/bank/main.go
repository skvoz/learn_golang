package main

import (
	"fmt"
	"mypackages/numbers"
)

func main() {
	var x uint = 41
	fmt.Printf("%d is even: %t\n", x, numbers.Even(x))
	fmt.Println(numbers.IsPrime(13), numbers.IsPrime(100))
	fmt.Println(numbers.OddAndPrime(25))
	fmt.Println(numbers.OddAndPrime(13))


}
