package main

import "fmt"

func main() {
	//could use in gorutine etc
	func(msg string) {
		fmt.Println(msg)
	}("I'm an anonymous function!") // calling the anonymous function
}
