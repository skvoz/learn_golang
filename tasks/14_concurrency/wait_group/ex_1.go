package main

import (
	"fmt"
	"sync"
)

func sayHello(wg *sync.WaitGroup, n string) {
	defer wg.Done()
	fmt.Printf("Hello, %s!\n", n)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go sayHello(&wg, "Mr. Wick")
	wg.Wait()
}
