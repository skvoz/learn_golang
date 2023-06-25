package main

import (
	"fmt"
	"strings"
)

func searchItem(sSlice []string, s string) bool {
	for _, v := range sSlice {
		if strings.ToLower(v) == strings.ToLower(s) {
			return true
		}
	}
	return false
}

func main() {
	animals := []string{"Lion", "tiger", "bear"}
	result := searchItem(animals, "beaR")
	fmt.Println(result) // => true
	result = searchItem(animals, "lion")
	fmt.Println(result) // => true
}
