package main

import "fmt"

func main() {
	//slice bytes example
	s1 := "I love golang!"

	fmt.Println(s1[2:5]) //return "lov"

	s2 := "中文维基是世界上"
	fmt.Println(s2[0:2]) //return � , return bytes

	//slice rune example
	rs := []rune(s2)
	fmt.Printf("%T\n", rs) // []int32

	fmt.Println(string(rs[0:6])) //中文维
}
