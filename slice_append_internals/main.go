package main

import "fmt"

func main() {
	//how capacity work ? when back array overflow, back array will extend on
	//2,4,8,16 ....n^2
	var nums []int

	fmt.Printf("%#v\n", nums)
	fmt.Printf("Length : %d, Capacity %d \n", len(nums), cap(nums))

	nums = append(nums, 1, 2)
	fmt.Printf("Length : %d, Capacity %d \n", len(nums), cap(nums))

	nums = append(nums, 3)
	fmt.Printf("Length : %d, Capacity %d \n", len(nums), cap(nums))

	nums = append(nums, 4)
	fmt.Printf("Length : %d, Capacity %d \n", len(nums), cap(nums))

	nums = append(nums, 100)
	fmt.Printf("Length : %d, Capacity %d \n", len(nums), cap(nums))

	nums = append(nums[0:4], 200, 300, 400, 500, 600)
	fmt.Printf("Length : %d, Capacity %d \n", len(nums), cap(nums))
	fmt.Println(nums)

	/**
		how capacity work ? when back array overflow, back array will extend on
		2,4,8,16 ....n^2
		[]int(nil)
		Length : 0, Capacity 0
		Length : 2, Capacity 2
		Length : 3, Capacity 4
		Length : 4, Capacity 4
		Length : 5, Capacity 8
		Length : 9, Capacity 16
		[1 2 3 4 200 300 400 500 600]
	**/

	letters := []string{"a", "b", "c", "d", "e", "f"}
	letters = append(letters[0:1], "x", "y")

	fmt.Println(letters, len(letters), cap(letters)) //output  [a x y] 3 6

	// fmt.Println(letters[4]) //error index out of range
	fmt.Println(letters[3:6]) //look in backing array :) lol kek
	//output DEF

	n1 := []int{10, 20, 30, 40, 50}
	n2 := n1[:3]
	n2[0] = 66
	fmt.Println(n1) //[66 20 30 40 50] wtf ?
	//n1 and n2 the same backing array

	n11 := []int{10, 20, 30, 40, 50}
	n22 := append(n11, 100) //another back array, append function create new backing array
	n22[0] = 66
	fmt.Println(n11)
}
