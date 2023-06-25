package main

import "fmt"

func main() {
	//can compare only with nil
	a := map[string]string{"A": "X"}
	b := map[string]string{"B": "Y"}

	// fmt.Println(a == b) //invalid operation: a == b (map can only be compared to nil)

	fmt.Println(a == nil) //false
	fmt.Println(b == nil) //false

	s1 := fmt.Sprintf("%s", a)
	s2 := fmt.Sprintf("%s", b)
	fmt.Println(s1, s2)
	if s1 == s2 {
		fmt.Println("Maps are equal")
	} else {
		fmt.Println("Maps are not equal")
	}
}
