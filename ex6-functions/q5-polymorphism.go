package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type square struct {
	length float64
	width float64
}

func (s square) area() float64 {
	return s.length * s.width
}

type shape interface {
	area() float64
}

func printArea(s ...shape)  {
	for _, value := range s {
		fmt.Println(value.area())
	}
}

func main() {
	x := circle{radius: 3}
	y :=square{
		length: 5,
		width:  6,
	}
	printArea(x, y)
}