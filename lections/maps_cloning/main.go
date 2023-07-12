package main

import "fmt"

func main() {
	friends := map[string]int{
		"Dan":   40,
		"Maria": 25,
	}

	neighbors := friends

	friends["Dan"] = 50

	fmt.Println(neighbors) //map[Dan:50 Maria:25]
	//wtf? we work with map header, value store not in var
	//so when we copy, var, or set new var, we have one value
	//so when change value, this  value will changed for all
	//vars

	people := make(map[string]int)
	 // colleagues = friends // -> ERROR, illegal with maps!
	//for create NEW var we need clone not copy
	for k, v := range friends {
		people[k] = v
	}

	people["Anne"] = 18
	fmt.Printf("%#v---------------------%#v\n", people, friends)
	delete(friends, "Dan")
	fmt.Println(friends) //map[Maria:25]

}
