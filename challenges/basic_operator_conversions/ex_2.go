package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i = 3
	var f = 3.2
	var s1, s2 = "3.14", "5"

	mystr, err := strconv.Atoi(s2)
	_ = err
	//myerror need
	//if err == nil {
	// fmt.Printf("i type is %T, i value is %v\n", is, is)
	// } else {
	// 	fmt.Println("Can not convert string to int.")
	// }
	myint := strconv.Itoa(i)

	fmt.Printf("type i is %T value %s, type s2 is %T value %d\n",
		myint, myint, mystr, mystr)

	myFloatToString := fmt.Sprintf("%f", f)
	fmt.Printf("type f is %T value %v\n",
		myFloatToString, myFloatToString)

	stringToFloat, err := strconv.ParseFloat(s1, 64)

	fmt.Printf("type s1 is %T value %v\n",
		stringToFloat, stringToFloat)
	//my error handle err
	/* 	if err1 == nil {
		fmt.Printf("f1's type: %T, f1's value: %v\n", f1, f1)
	} else {
		fmt.Println("Cannot convert string to float64.")
	}
	*/
}
