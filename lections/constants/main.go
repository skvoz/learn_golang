package main

import "fmt"

func main() {
	const foo int = 1

	//const foo int  // will generate error
	_ = foo

	const n1, m1 = 1, 2
	_, _ = n1, m1

	const (
		foo1 = 1
		bar1
		baz1
	) //lolwhat , first value will set in all const

	fmt.Println(foo1, bar1, baz1)

	//const rules
	//1. can't change
	//2. can't initiate at runtime (math.pow - runtime function , len - not runtime function )
	//3. can't initiate with variables
}
