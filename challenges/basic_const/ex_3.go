package main

import "fmt"

func main() {
	const secPerDay int32 = 24 * 60 * 60
	const daysYear int32 = 365
	fmt.Printf("sec per year %d\n", secPerDay*daysYear)
}
