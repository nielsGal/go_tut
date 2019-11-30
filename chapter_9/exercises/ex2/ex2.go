package main

import "fmt"

type person struct {
	first, last string
	age         int
}

type human interface {
	speak()
}

func main() {
	p1 := person{
		first: "James",
		last:  "bond",
		age:   32,
	}
	saySomething(p1)
}

func (p *person) speak() {
	fmt.Println("hi i am", p.first, p.last)
}
func saySomething(p person) {
	p.speak()
}
