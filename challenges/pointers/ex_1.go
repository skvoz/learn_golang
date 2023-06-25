package main

import "fmt"

func main() {
	x := 10.10

	fmt.Printf("%p\n", &x)
	ptr := &x

	fmt.Printf("Type ptr is:%T, Value ptr is:%+v\n", ptr, ptr)
	fmt.Printf("address ptr is:%+v, Value ptr through  dereferencing operator) is:%+v\n",
		ptr, *ptr)
	//here my error , need &ptr, *ptr , because need address ptr (address of address :))
	/**
		0xc0001a4000
	Type ptr is:*float64, Value ptr is:0xc0001a4000
	Value ptr is:0xc0001a4000, Value ptr through  dereferencing operator) is:10.1
	**/

}
