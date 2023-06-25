package main

import "fmt"

type km float64
type mile float64

func main() {
	type age int        //int is its underline type
	type oldAge age     //int is its underline type
	type veryOldAGe age //int is its underline type

	type speed uint
	var s1 speed = 10

	var s2 speed = 20

	fmt.Println(s2 - s1) // 20

	var x uint
	// x = s1 //return error
	x = uint(s1)
	_ = x

	// var s3 speed  = x //error
	var s3 speed = speed(x) //wow convert number to definded type!!
	_ = s3

	var parisToLandon km = 465
	var distanceInMiles mile

	distanceInMiles = mile(parisToLandon) / 0.621 //748.792270531401
	fmt.Println(distanceInMiles)
}
