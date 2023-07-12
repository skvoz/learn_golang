package main

import (
	"fmt"
)

func searchItem(sSlice []string, s string) bool {
	for _, v := range sSlice {
		if v == s {
			return true
		}
	}
	return false
}

func main() {
	animals := []string{"lion", "tiger", "bear"}
	result := searchItem(animals, "bear")
	fmt.Println(result) // => true
	result = searchItem(animals, "pig")
	fmt.Println(result) // => false
}
