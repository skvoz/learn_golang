package main

import "fmt"

func main() {
	m := map[int]bool{1: true, 2: false, 3: false}

	delete(m, 2)

	for k, v := range m {
		fmt.Printf("key %d value %#v\n", k, v)
	}

	fmt.Println("")

}
