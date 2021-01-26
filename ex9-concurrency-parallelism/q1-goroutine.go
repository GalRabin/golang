package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Goroutines", runtime.NumGoroutine())

	numOfGoRoutines := 2
	wg.Add(numOfGoRoutines)
	go func() {
		for i :=0; i < 5; i++ {
			fmt.Println("From goroutine 1 - Index", i)
		}
		wg.Done()
	}()

	go func() {
		for i :=0; i < 5; i++ {
			fmt.Println("From goroutine 2 - Index", i)
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("About to exit")
}
