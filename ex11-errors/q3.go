package main

import (
	"fmt"
	"log"
)

type myBuiltInError struct {
	err string
}

func (e *myBuiltInError) Error() string {
	return e.err
}

func getFile() (string, error) {
	return "Gal", &myBuiltInError{"What?????"}
}

func main() {
	s, err := getFile()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)
}
