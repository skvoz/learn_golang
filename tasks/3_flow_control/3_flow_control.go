package main

import "fmt"

func main() {
	fmt.Println("ex1#################")
	ex1()
	/**
	ex1#################
	7
	14
	21
	28
	**/
}

func ex1() {
	for i := 1; i <= 50; i++ {
		if i == 31 {
			break
		}
		if i%7 == 0 {
			fmt.Println(i)
		}
	}
}
