package main

import "fmt"

func main() {
	//map is a hash, associative array?
	var employees map[string]string

	fmt.Printf("%#v\n", employees) //map[string]string(nil)

	fmt.Printf("No of paris: %d\n", len(employees))                     //No of paris: 0
	fmt.Printf("No value  for key %q is %q\n", "Dan", employees["Dan"]) //No value  for key "Dan" is ""

	var accounts map[string]float64
	fmt.Printf("%#v\n", accounts["0x323"]) //0

	// var m1 map[[]int]string //error not possible
	var m1 map[[5]int]string
	_ = m1

	// employees["Dan"] = "Programmer" //panic: assignment to entry in nil map

	people := map[string]float64{}
	people["Johne"] = 21.4
	people["Marry"] = 10
	fmt.Println(people) //map[Johne:21.4 Marry:10]

	map1 := make(map[int]int)
	map1[4] = 8

	balances := map[string]float64{
		"USD": 323.11,
		"EUR": 432.22,
		// 50:50will error
	}

	fmt.Println(balances) //map[EUR:432.22 USD:323.11]

	m := map[int]int{1: 10, 2: 20, 3: 30}
	_ = m

	balances["USD"] = 500.2
	balances["GBR"] = 10
	fmt.Println(balances)
	fmt.Println(balances["RON"]) //0

	balances["RON"] = 1
	v, ok := balances["RON"] // v value, ok - true if value exists, false if not

	if ok {
		fmt.Println("The RON balance is : ", v) //The RON balance is :  1
	} else { 
		fmt.Println("The RON key doesn't exist in the map!")
	}

	for k, v := range balances {
		fmt.Printf("Key: %#v, value: %#v\n", k, v)
	}
	/**
	
Key: "USD", value: 500.2
Key: "EUR", value: 432.22
Key: "GBR", value: 10
Key: "RON", value: 1
map[EUR:432.22 GBR:10 RON:1]
**/

	delete(balances, "USD")
	fmt.Println(balances) //map[EUR:432.22 GBR:10 RON:1]

}
