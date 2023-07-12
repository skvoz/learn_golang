package main

import "fmt"

func main() {
	for i := 1; i <= 50; i++ {
		if i%7 != 0 {
			continue
		}

		fmt.Println(i)

		if i%21 == 0 {
			break
		}
	}
}
