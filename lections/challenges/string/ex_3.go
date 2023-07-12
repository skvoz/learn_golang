package main

import "fmt"

func main() {
	s1 := "țară means country in Romanian"

	for i, _ := range s1 {
		fmt.Printf("%v", s1[i])
	}
	fmt.Println("") //20097114196321091019711011532991111171101161141213210511032821111099711010597110

	for i, r := range s1 {
		fmt.Printf("byte index: %d -> rune: %c\n", i, r)
	}

	fmt.Println()
}
