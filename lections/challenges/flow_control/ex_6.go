package main

import "fmt"

func main() {
	switch age := -9; true {
	case age < 0 || age > 100:
		fmt.Println("Invalid Age")
	case age < 18:
		fmt.Println("You are minor!")
	case age == 18:
		fmt.Println("Congratulations! You've just become major!")
	default:
		fmt.Println("You are major!")
	}
}
