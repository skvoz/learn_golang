package main

import (
	"fmt"
)

//purpouse from course author
func main() {

	friends := []string{"Marry", "John", "Paul", "Diana"}
	friendsNew := make([]string, len(friends))
	nn := copy(friendsNew, friends)
	_ = nn
	fmt.Printf("%#v\n", friends)
	friendsNew[0] = "1"
	fmt.Printf("%#v\n", friendsNew)
}
