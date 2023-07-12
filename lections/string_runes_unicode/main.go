package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//diff between stirng and runes
	var a, b = 'a', 'b'
	fmt.Printf("Type: %T, Value:%d \n", a, a)
	var c, d = "c", "d"
	fmt.Printf("Type: %T, Value:%s \n", c, c)
	_, _ = b, d

	str := "ţară"
	fmt.Println(len(str))
	fmt.Println("byte (not rune) at position 1:", str[1])

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}
	fmt.Println("\n")
	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("%c", r)
		i += size
	}
	fmt.Println("\n")

	for _, r := range str {
		fmt.Printf("%c", r)
	}
}
