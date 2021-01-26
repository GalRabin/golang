package main

import "fmt"

// Will cause deadlock
//func main() {
//	c := make(chan int)
//
//	c <- 42
//
//	fmt.Println(<-c)
//}


// Buffer solution
//func main() {
//	c := make(chan int, 2)
//
//	c <- 42
//
//	fmt.Println(<-c)
//}


// Executing in anonymous function
func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()


	fmt.Println(<-c)
}


