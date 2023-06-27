package main

import (
	"fmt"
	"sync"
)

/**
my solution
func sayHello(n string) {
	fmt.Printf("Hello, %s!\n", n)
}

func main() {
	defer sayHello("Mr. Wick")
}
**/
//course example
func sayHello(n string, w *sync.WaitGroup) {
	fmt.Printf("Hello, %s!\n", n)
	w.Done()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	go sayHello("Mr. Wick", &wg)

	wg.Wait()
}
