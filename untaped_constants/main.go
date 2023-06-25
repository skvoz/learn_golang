package main

import "fmt"

func main() {
	const a float64 = 5.1 //typed const

	const b = 6.7 //untyped const

	const c float64 = a * b

	const str = "hello " + "GO!"

	const d = 5 > 10

	fmt.Println(d)

	// const x int = 5
	//  const y float64 = 2.2 * 5  //can't do this , can't multyplay float by int

	const x = 5 //untyped const

	const y = 2.2 * x //so float * untyped = float :)

	fmt.Printf("%T\n", y) // return float64

	var i int = x //x changes to int

	var j float64 = x //var j float64=float64(x)

	var p byte = x //changes to byte

	fmt.Println(i, j, p) // return 5 5 5

	const r = 5

	var rr = r

	fmt.Printf("%T\n", rr) //return int, becase variable can't be untyped
	// so get type from value
}
