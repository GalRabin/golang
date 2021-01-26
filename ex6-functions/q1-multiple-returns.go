package main

import "fmt"

func firstFunc() int {
	return 8
}

func secondFunc() (string, int) {
	return "Gal", 9
}

func main() {
	fmt.Println(firstFunc())
	fmt.Println(secondFunc())
}