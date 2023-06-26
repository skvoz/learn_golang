package main

import (
	"fmt"
	"runtime"
	"time"
)

func f1() {
	fmt.Println("f1 gorutine execution started")
	for i := 0; i < 3; i++ {
		fmt.Println("f1, i=", i)
	}
	fmt.Println("f1 execution finished")
}

func f2() {
	fmt.Println("f2 gorutine execution started")
	for i := 5; i < 8; i++ {
		fmt.Println("f2, i=", i)
	}
	fmt.Println("f2 execution finished")
}

func main() {
	fmt.Println("main execution started")
	fmt.Println("No. of CPUs:", runtime.NumCPU())
	fmt.Println("No. of Gorutines:", runtime.NumGoroutine())

	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Arch:", runtime.GOARCH)
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))

	go f1()

	fmt.Println("No. Gorutines after go f1():", runtime.NumGoroutine())

	f2()

	//withot this we don't see f1  result
	//look an gorutine mechanism
	time.Sleep(time.Second * 2)
	fmt.Println("main execution stopped")

}
