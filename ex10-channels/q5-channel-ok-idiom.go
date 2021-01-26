package main

import "fmt"

func main()  {
	c := make(chan int)
	go func() {
		for i := 0; i < 2; i++ {
			c <- i
		}
		close(c)
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	v, ok = <-c
	fmt.Println(v, ok)

	v, ok = <-c
	fmt.Println(v, ok)
}