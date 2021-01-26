package main

import "fmt"

func funcAsArgument(f func() int)  {
	fmt.Println(f())
}

func main() {
	x := func() int {
		return 4
	}
	funcAsArgument(x)
}