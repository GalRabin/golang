package main

import "fmt"

type person struct {
	name string
	last string
	age int
}

func (p person) talk()  {
	p.age = 8
}

func (p *person) speak()  {
	p.age = 8
}

func main() {
	p := person{
		name: "Gal",
		last: "Rabin",
		age:  20,
	}
	p.talk()
	fmt.Println(p)
}