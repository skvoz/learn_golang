package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	//a[start:stop] - element stop not output new  slice the end befor stop element

	b := a[1:3] //get part slice
	fmt.Printf("%v, %T\n", b, b)

	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1[1:3]
	fmt.Println(s2)
	fmt.Println(s1[2:]) //s1[2:len(s1)]
	fmt.Println(s1[:3]) //s1[0:3]
	// fmt.Println(s1[:100]) //error out of range

	s1 = append(s1[:4], 100)

	fmt.Println(s1) //[1 2 3 4 100]]
	s1 = append(s1[:4], 200)
	fmt.Println(s1) //[1 2 3 4 200]]
}
