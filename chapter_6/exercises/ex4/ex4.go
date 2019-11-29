package main

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
	p1.speak()
}

func (p person) speak() {
	println("hello i am", p.first, p.last)
}
