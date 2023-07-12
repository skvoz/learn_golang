package main

import "fmt"

func main() {
	r := 'ă'
	fmt.Printf("%T\n", r) //int32

	s1, s2 := "ma", "m"

	s3 := s1 + s2 + string(r)

	fmt.Println(s3)
}
