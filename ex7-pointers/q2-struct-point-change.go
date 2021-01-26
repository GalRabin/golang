package main

import "fmt"

type person struct {
	name string
	lastName string
	age int
}

func changeMe(p *person)  {
	p.name = "Miryam"
	p.lastName = "Berdugo"
}

func main() {
	p := person{
		name:     "Gal",
		lastName: "Rabin",
		age:      27,
	}
	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)
}
