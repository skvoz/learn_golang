package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(50)
	for i := 100.; i < 150; i++ {

		go func(x float64, wg *sync.WaitGroup) {
			defer wg.Done()
			f := math.Sqrt(x)
			fmt.Println(f)
		}(i, &wg)
	}

	wg.Wait()
}
