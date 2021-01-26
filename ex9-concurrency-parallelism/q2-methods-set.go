package main

import "fmt"

type person struct {
	name string
	last string
	age int
}

func (p *person) speak()  {
	fmt.Println("My name is", p.name, p.last, ", Age", p.age)
}

type human interface {
	speak()
}

func saySomething(h human)  {
	h.speak()
}

func main() {
	p1 := person{
		name: "Gal",
		last: "Rabin",
		age:  30,
	}

	// You cannot do this
	//saySomething(p1)
	saySomething(&p1)
}