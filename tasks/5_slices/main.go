package main

import "fmt"

func main() {
	fmt.Println("Ex1#############################")
	arr := []string{"foo", "bar", "baz"}
	for index, v := range arr {

		fmt.Printf("index %d, val %q\n", index, v)
	}
	/**
	  index 0, val "foo"
	  index 1, val "bar"
	  index 2, val "baz"
	  **/

	fmt.Println("Ex2#############################")

	nums := []float64{1.1, 2.2, 3.3}
	nums = append(nums, 10.1)
	nums = append(nums, 4.1, 5.5, 6.6)
	fmt.Printf("%#v\n", nums)
	//[]float64{1.1, 2.2, 3.3, 10.1, 4.1, 5.5, 6.6}

	fmt.Println("Ex3#############################")
	nums1 := []int{5, -1, 9, 10, 1100, 6, -1, 6}

	nums11 := nums1[2 : len(nums1)-2]

	for _, v := range nums11 {
		fmt.Print(v, " ")
	}
	fmt.Println()
	//3.3 10.1 4.1 5.5

	fmt.Println("Ex4#############################")
	friends := []string{"Marry", "John", "Paul", "Diana"}

	var friends1 = make([]string, len(friends))
	copy(friends1, friends)
	friends[0] = "foo"
	fmt.Printf("src array %#v \ndst array %#v\n", friends, friends1)

	fmt.Println("Ex5#############################")
	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
	newYears := []int{}
	newYears = append(years[:3], years[len(years)-3:]...)
	fmt.Printf("%#v\n", newYears)
}
