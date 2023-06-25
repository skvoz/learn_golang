package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// s1 := []int{10, 20, 30, 40, 50}

	// s3, s4 := s1[0:2], s1[1:3]

	// s3[1] = 600

	// fmt.Println(s1)
	// fmt.Println(s4)

	// arr1 := [4]int{10, 20, 30, 40}

	// slice1, slice2 := arr1[0:2], arr1[1:3]
	// arr1[1] = 2
	// fmt.Println(arr1, slice1, slice2)

	cars := []string{"ford", "honda", "audi", "range rover"}
	newCars := []string{}

	newCars = append(newCars, cars[1:3]...)
	cars[0] = "Nissan"
	fmt.Println(cars, newCars)

	s1 := []int{10, 20, 30, 40, 50}
	newSlice := s1[0:3]
	fmt.Println(len(newSlice), cap(newSlice))

	newSlice = s1[2:5] //{30, 40, 50}
	fmt.Println(len(newSlice), cap(newSlice))

	fmt.Printf("%p\n", &s1)

	fmt.Printf("%p %p \n", &s1, &newSlice)

	newSlice[0] = 1000
	fmt.Println("s1: ", s1)
	/**
	[Nissan honda audi range rover] [honda audi]
	3 5
	3 3
	0xc000010060
	0xc000010060 0xc000010078
	s1:  [10 20 1000 40 50]
	*/
	a := [5]int{1, 2, 3, 4, 5}
	s := []int{1, 2, 3, 4, 5}

	fmt.Printf("array's size in bytes: %d \n", unsafe.Sizeof(a))
	fmt.Printf("slices's size in bytes: %d \n", unsafe.Sizeof(s))
}
