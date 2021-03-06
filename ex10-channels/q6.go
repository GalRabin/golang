package main

import "fmt"

func push(c chan<- int)  {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func pull(c <-chan int)  {
	for v := range c {
		fmt.Println(v)
	}
}

func main() {
	c := make(chan int)
	go push(c)
	pull(c)
}