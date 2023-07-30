package main

import (
	"fmt"
	"sync"
)

func deposit(b *int, n int, wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	defer m.Unlock()
	*b += n
	wg.Done()
}

func withdraw(b *int, n int, wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	defer m.Unlock()
	*b -= n
	wg.Done()
}

func main() {
	var m sync.Mutex
	var wg sync.WaitGroup

	wg.Add(200)

	balance := 100

	for i := 0; i < 100; i++ {
		go deposit(&balance, i, &wg, &m)
		go withdraw(&balance, i, &wg, &m)
	}
	wg.Wait()

	fmt.Println("Final balance value:", balance)
}
