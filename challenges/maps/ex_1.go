package main

import "fmt"

func main() {
	m1 := map[int]int{}
	fmt.Printf("type: %T, value: %#v\n", m1, m1) //type: map[int]int, value: map[int]int{}

	m2 := map[int]string{}
	m2[10] = "Abba"
	fmt.Printf("exists key: %v, unexsits key %v\n", m2[10], m2[20]) //exists key: Abba, unexsits key

}
