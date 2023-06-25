package main

import "fmt"

func main() {
	//lol ex
	myArray := [4]float64{1.2, 5.6}

	myArray[2] = float64(6)

	a := 10
	myArray[0] = float64(a)

	myArray[3] = 10.10

	fmt.Println(myArray)

}
