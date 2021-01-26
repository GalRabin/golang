package main

import "fmt"

func deferExample()  {
	defer func() {
		fmt.Println("First - But last")
	}()
	fmt.Println("Second")
}

func main() {
	defer deferExample()
	fmt.Println("Third - But first")
}