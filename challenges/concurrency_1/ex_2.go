package main

import (
	"fmt"
	"sync"
)

func sum(f1 float64, f2 float64, w *sync.WaitGroup) {
	fmt.Printf("%.2f\n", f1+f2)
	w.Done()
}
func main() {
	var wg sync.WaitGroup

	wg.Add(3)
	for i := 0; i <= 2; i++ {
		go sum(1.1, (0.1 + float64(i)), &wg)
	}
	wg.Wait()
}
