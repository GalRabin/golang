package main

import "fmt"

func returnsFunc() func() {
	return func() {
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}
	}
}

func main() {
	x := returnsFunc()
	fmt.Printf("%T", x)
	x()
}
