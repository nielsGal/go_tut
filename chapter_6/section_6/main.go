package main

import "fmt"

type person struct {
	first, last string
}

type secretAgent struct {
	person person
	ltk    bool
}

func (s secretAgent) speak() {
	fmt.Printf("the name is %s , %s %s\n", s.person.last, s.person.first, s.person.last)
}

func (p person) speak() {
	fmt.Printf("the name is %s , %s %s\n", p.last, p.first, p.last)
}

func bar(h human) {
	switch h.(type) {
	case secretAgent:
		fmt.Println(h, "is a secret agent", h.(secretAgent).person.first)
	case person:
		fmt.Println(h, "is a person", h.(person).first)
	default:
		fmt.Println(h, " was passed into bar")
	}

}

type human interface {
	speak()
}

func main() {
	var sa1 human
	sa1 = secretAgent{
		person: person{
			first: "james",
			last:  "bond",
		},
		ltk: true,
	}
	p1 := person{
		first: "doctor",
		last:  "no",
	}
	fmt.Println(sa1)
	fmt.Println(p1)
	bar(p1)
	bar(sa1)
	sa1 = p1
	fmt.Println(sa1)
}
