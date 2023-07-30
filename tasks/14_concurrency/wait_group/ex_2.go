package main

import (
	"fmt"
	"sync"
)

func sum(a float32, b float32, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%f\n", a+b)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go sum(1, 2, &wg)
	go sum(4, 4, &wg)
	go sum(5, 5, &wg)
	wg.Wait()
}
