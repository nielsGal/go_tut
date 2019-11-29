package main

import "fmt"

type person struct {
	first, last string
	age         int
}

func main() {
	p1 := person{
		first: "james",
		last:  "bond",
		age:   32,
	}

	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}

func changeMe(pptr *person) {
	pptr.first = "a new first name"
}
