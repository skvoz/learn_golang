package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := string(99) //return int as ascii code

	fmt.Println(s)

	// s1 := string(44.2) //cannot convert 44.2 (untyped float constant) to type string
	// fmt.Println(s1)

	var myStr = fmt.Sprintf("%f", 44.2)
	fmt.Println(myStr) //"44.200000"

	var myStr1 = fmt.Sprintf("%d", 34234)
	fmt.Println(myStr1)        //"34234"
	fmt.Println(string(34234)) //薺

	s1 := "3.123"
	fmt.Printf("%T\n", s1) //string

	//MAGIC function to convert string to another types
	var f1, err = strconv.ParseFloat(s1, 64)

	_ = err

	fmt.Println(f1) //3.123

	//atoi = ascii to int
	i, err := strconv.Atoi("-50")
	s2 := strconv.Itoa(20)
	fmt.Printf("i type is %T, i value is %v\n", i, i) //i type is int, i value is -50
	fmt.Printf("s type is %T, s value is %q", s2, s2) //s type is string, s value is "20"
}
