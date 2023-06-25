package main

import "fmt"

func main() {
	slice1 := []string{"foo", "bar", "baz"}

	for i, v := range slice1 {
		fmt.Printf("index el %d, value el %v\n", i, v)
	}
}
