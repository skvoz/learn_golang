package main

import "fmt"

func main() {
	var a = 1
	var b = 1.1

	//a = b return error, a and b has diff types
	a = int(b) //change b type to int
	_, _ = a, b

	/**
	* go zero values
	* - numeric type: 0
	* - bool type : false
	* - string type: "",
	* - pointer type : nil
	**/

	// The zero-value mechanism of Go ensures that a variable always holds a well defined value of its type
	var value int                         // initialized with 0
	var price float64                     // initialized with 0.0
	var name string                       // initialized with empty string -> ""
	var done bool                         // initialized with false
	fmt.Println(value, price, name, done) // -> 0 0.0 ""  false
}
