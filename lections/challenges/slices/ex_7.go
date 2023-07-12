package main

import "fmt"

//purpouse from course author
func main() {

	friends := []string{"Marry", "John", "Paul", "Diana"}
	friendsNew := []string{}
	friendsNew = append(friendsNew, friends...)
	friendsNew[0] = "1"
	fmt.Println(friends, friendsNew)

}
