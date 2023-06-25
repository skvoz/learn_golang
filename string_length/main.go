package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s1 := "golang"
	fmt.Println(len(s1)) //6 ascii leters - 6 bytes

	name := "Codurţa"
	fmt.Println(len(name)) //8, ţ encoded with 2 bytes

	fmt.Println(utf8.RuneCountInString(name))                                                     //7
	fmt.Printf("countr bytes in rune: %d, count rune:%d", len("ţ"), utf8.RuneCountInString("ţ")) //7

}
