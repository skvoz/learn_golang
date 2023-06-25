package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func f1(wg *sync.WaitGroup) {
	fmt.Println("f1 gorutine execution started")
	for i := 0; i < 3; i++ {
		fmt.Println("f1, i=", i)
		time.Sleep(time.Second)
	}
	fmt.Println("f1 execution finished")
	wg.Done()
}

func f2() {
	fmt.Println("f2 gorutine execution started")
	for i := 5; i < 8; i++ {
		fmt.Println("f2, i=", i)
		// time.Sleep(time.Second)
	}
	fmt.Println("f2 execution finished")
	// wg.Done()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	fmt.Println("main execution started")
	fmt.Println("No. of CPUs:", runtime.NumCPU())
	fmt.Println("No. of Gorutines:", runtime.NumGoroutine())

	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Arch:", runtime.GOARCH)
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))

	go f1(&wg)

	fmt.Println("No. Gorutines after go f1():", runtime.NumGoroutine())

	f2()

	//withot this we don't see f1  result
	//look an gorutine mechanism
	// time.Sleep(time.Second * 2)
	wg.Wait()
	fmt.Println("main execution stopped")
	//** EXPECTED OUTPUT: **//
	// main execution started
	// No. of Goroutines: 2
	// f2 execution started
	// f1(goroutine) execution started
	// f1, i= 0
	// f2(), i= 5
	// f2(), i= 6
	// f2(), i= 7
	// f2 execution finished
	// f1, i= 1
	// f1, i= 2
	// f1 execution finished
	// main execution stopped
}
