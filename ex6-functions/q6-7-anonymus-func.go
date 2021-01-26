package main

import "fmt"

func main() {
	// Will be used when running parallel
	x := func () int {
		return 4
	}
	fmt.Printf("%T", x)
	fmt.Println(x())
}
