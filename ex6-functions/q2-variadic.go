package main

import "fmt"

func foo(param ...int) int {
	total := 0
	for _, value := range param {
		total += value
	}

	return total
}

func bar(param []int) int {
	total := 0
	for _, value := range param {
		total += value
	}

	return total
}

func main() {
	ii := []int{1,2,3,4}
	fmt.Println(foo(ii...))
	fmt.Println(bar(ii))
}
