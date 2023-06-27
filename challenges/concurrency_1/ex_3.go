package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func(f float64, wg *sync.WaitGroup) {
		x := math.Sqrt(f)
		fmt.Printf("Square root of %.2f is %.4f\n", f, x)
		wg.Done()
	}(4.4, &wg)
	wg.Wait()
}
