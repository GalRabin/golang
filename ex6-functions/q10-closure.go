package main

import "fmt"

func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	a := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
}

