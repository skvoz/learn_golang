package main

import "fmt"

func main() {
	var distanceMilKm float64 = 149.6
	var distance float64 = distanceMilKm * 1000000 * 1000
	var speed int64 = 299792458

	time := int64(distance) / speed

	fmt.Println("Speed :", time)

	/* correct solution
	const (
		distance = 149_600_000 * 1000 // distance from the Sun to Earth in m
		// (it's allowed to use underscores in numbers for a better readability)

		speed = 299_792_458 // speed of light in m / s
	)

	const time = distance / speed // time in seconds

	fmt.Printf("The Sunlight reaches Earth in %v seconds.\n", time)
	*/
}
