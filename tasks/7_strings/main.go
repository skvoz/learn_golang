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
	/**
		Ex1#######################
	name is Constantine,
	country is TrampampamCity
	a) He says: "Hello"
	b) C:\Users\a.txt
	Ex2#######################
	type is int32, value is 259
	"mamă"
	Ex3#######################
	200 155 97 114 196 131 32 109 101 97 110 115 32 99 111 117 110 116 114 121 32 105 110 32 82 111 109 97 110 105 97 110
	țară means country in Romanian
	Ex4#######################
	228 189 160 229 165 189 32 71 111 33
	[]byte{0xe4, 0xbd, 0xa0, 0xe5, 0xa5, 0xbd, 0x20, 0x47, 0x6f, 0x21}
	Ex4#######################
	[]int32{20320, 22909, 32, 71, 111, 33}
	index 0, value  你
	index 1, value  好
	index 2, value
	index 3, value  G
	index 4, value  o
	index 5, value  !
	**/
}
