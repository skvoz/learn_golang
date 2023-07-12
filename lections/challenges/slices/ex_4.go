// run this program in a terminal with arguments
// ex: go run main.go 5 7.1 9.9 10

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

//purpouse from course author
func main() {

	if len(os.Args) < 2 { //if not run with arguments
		log.Fatal("Please run the program with arguments (2-10 numbers)!")

	}

	//taking the arguments in a new slice
	args := os.Args[1:]

	// declaring and initializing sum and product of type float64
	sum, product := 0., 1.

	if len(args) < 2 || len(args) > 10 {
		fmt.Println("Please enter between 2 and 10 numbers!")
	} else {

		for _, v := range args {
			num, err := strconv.ParseFloat(v, 64)
			if err != nil {
				continue //if it can't convert string to float64, continue with the next number
			}
			sum += num
			product *= num
		}
	}

	fmt.Printf("Sum: %v, Product: %v\n", sum, product)

}

//my work need more check, look in rquire, main construction work well
// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func main() {
// 	args := os.Args[1:]
// 	fmt.Printf("%#v\n", args)
// 	sum := 0
// 	prod := 1
// 	for _, v := range args {
// 		fmt.Println(v)
// 		intV, err := strconv.Atoi(v)
// 		if err != nil {
// 			fmt.Println(err)
// 			break
// 		} else {
// 			sum += intV
// 			prod *= intV
// 		}
// 	}
// 	fmt.Println(sum, prod)
// }
