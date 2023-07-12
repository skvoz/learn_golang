package main

import (
	"fmt"
)

func f1(n uint) (int, int) {

	factN := 1
	sumN := 0
	for i := 1; i <= int(n); i++ {
		factN *= i
		sumN += i
	}

	return factN, sumN
}

func main() {
	fmt.Println(f1(4))
}
