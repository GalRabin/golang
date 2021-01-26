package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	c := make(chan int)
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(c chan<- int) {
			for j := 0; j < 2; j++ {
				c <- j
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}
