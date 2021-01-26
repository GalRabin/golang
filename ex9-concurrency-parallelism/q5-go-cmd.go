package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("OS:\t\t", runtime.GOOS)
	fmt.Println("Arch:\t\t", runtime.GOARCH)
	fmt.Println("CPUs:\t\t", runtime.NumCPU())
	fmt.Println("Goroutine:\t", runtime.NumGoroutine())
}
