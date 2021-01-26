package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen1(q)

	receive1(c, q)

	fmt.Println("about to exit")
}

func receive1(c, q<-chan int)  {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case v := <-q:
			if v == 1 {
				fmt.Println("Exit")
				return
			}
		}
	}
	}


func gen1(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q<-1
		close(c)
	}()

	return c
}
