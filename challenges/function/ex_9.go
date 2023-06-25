package main

import "fmt"

func foo(i int) {
	fmt.Println(i)
}

func main() {
	a := foo
	a(1) //return 1
}
