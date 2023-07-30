package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func(x float64, wg *sync.WaitGroup) {
		defer wg.Done()
		f := math.Sqrt(x)
		fmt.Println(f)
	}(20.0, &wg)
	wg.Wait()
}
