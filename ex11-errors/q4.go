package main

import (
	"errors"
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}


func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, sqrtError{"234", "3423", errors.New("Whats going on???")}
	}
	return 42, nil
}


func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

