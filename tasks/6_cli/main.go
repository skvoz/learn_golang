package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 { //if not run with arguments
		log.Fatal("Please run the program with arguments (2-10 numbers)!")

	}

	args := os.Args[1:]
	sum := 0.
	prod := 1.
	if len(args) < 2 || len(args) > 10 {
		fmt.Println("Please enter between 2 and 10 numbers!")
	} else {
		for _, v := range args {
			num, err := strconv.ParseFloat(v, 64)
			if err != nil {
				continue //if it can't convert string to float64, continue with the next number
			}
			sum += num
			prod *= num
		}
	}
	fmt.Printf("sum %v, prod %v\n", sum, prod)
	// go run main.go 1 2 3 4
	// sum 10, prod 24
}
