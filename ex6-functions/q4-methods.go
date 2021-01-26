package main

import "fmt"

type person struct {
	name string
	lastName string
	age int
}

func (p person) speak() string {
	return fmt.Sprintf("My name is %v, age %v", p.name, p.age)
}

func main() {
	p := person{
		name:     "Gal",
		lastName: "Rabin",
		age:      27,
	}
	fmt.Println(p.speak())
}