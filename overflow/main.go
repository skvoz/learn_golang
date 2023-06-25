package main

import (
	"fmt"
	"math"
)

func main() {
	var x uint8 = 255

	x++ //overflow , x is 0

	fmt.Println(x) //return 0

	// a := int8(255 + 1)
	var b int8 = 127
	fmt.Printf("%d\n", b+1) //return -128
	b = -128

	b--
	fmt.Println(b) //127

	f := float32(math.MaxFloat32)

	fmt.Println(f) //3.4028235e+38
	f = f * 1.2
	fmt.Println(f) //+Inf (infinite) beskonechnost

	// const i int8 = 300 //found error in compile time

	//if use very big numbers(arbitrary-precision arithmetic) need package math/big

}
