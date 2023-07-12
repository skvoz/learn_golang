package main

import (
	"fmt"
	"strconv"
)

func myFunc(s string) int {
	//need add error handler like:
	//if err != nill return error
	foo, _ := strconv.Atoi(s)
	bar, _ := strconv.Atoi(s + s)
	baz, _ := strconv.Atoi(s + s + s)
	return foo + bar + baz
}

func main() {
	fmt.Println(myFunc("1")) //return 123
}
