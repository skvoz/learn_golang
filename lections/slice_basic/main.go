package main

import "fmt"

func main() {
	var cities []string //empty slice nil

	fmt.Println("cities is  equal to nil:", cities == nil)
	fmt.Printf("cities %#v\n", cities)
	fmt.Println(len(cities))

	numbers := []int{2, 3, 4, 5}
	fmt.Println(numbers)

	nums := make([]int, 2)
	fmt.Printf("cities %#v\n", nums)

	type names []string
	friends := names{"Dan", "maria"}
	fmt.Println(friends)

	myFriend := friends[0]
	fmt.Println("My best frien is ", myFriend)

	for index, value := range numbers {
		fmt.Printf("index %v, value %v\n", index, value)
	}

	var n []int

	n = numbers
	fmt.Println(n)
}
