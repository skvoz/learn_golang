package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ex1#######################")
	var name string = "Constantine"
	var country string = "TrampampamCity"

	fmt.Printf(`name is %s,
country is %s`, name, country)

	fmt.Println()

	fmt.Println(`a) He says: "Hello"
b) C:\Users\a.txt`)

	fmt.Println("Ex2#######################")
	var r rune = 'ă'

	fmt.Printf("type is %T, value is %v\n", r, r)
	ma, m := "ma", "m"
	newString := ma + m + string(r)
	fmt.Printf("%#v\n", newString)

	fmt.Println("Ex3#######################")
	s1 := "țară means country in Romanian"

	for i := 0; i < len(s1); i++ {
		fmt.Print(s1[i], " ")
	}

	fmt.Println()

	for _, v := range s1 {
		fmt.Printf("%c", v)
	}
	fmt.Println()

	fmt.Println("Ex4#######################")
	s := "你好 Go!"
	for i := 0; i < len(s); i++ {
		fmt.Print(s[i], " ")
	}
	fmt.Println()
	b := []byte(s)
	fmt.Printf("%#v\n", b)

	fmt.Println("Ex4#######################")

	r1 := []rune(s)
	fmt.Printf("%#v\n", r1)
	for i, v := range r1 {
		fmt.Printf("index %d, value  %c\n", i, v)
	}
}
