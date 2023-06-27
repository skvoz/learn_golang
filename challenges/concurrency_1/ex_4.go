package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(50)
	for i := 100; i < 150; i++ {
		go func(f float64, wg *sync.WaitGroup) {
			x := math.Sqrt(f)
			fmt.Printf("Square root of %.2f is %.4f\n", f, x)
			wg.Done()
		}(float64(i), &wg)
	}

	wg.Wait()
}
