package main

import "fmt"

func foo() {
	fmt.Println("funciton foo call first")
}

func bar() {
	fmt.Println("function bar call second")
}
func main() {
	//call in the end of script
	defer foo()

	bar()
	/**
	function bar call second
	funciton foo call first
	**/
}
