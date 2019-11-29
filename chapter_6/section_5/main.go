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

func main() {
	sa1 := secretAgent{
		person: person{
			first: "james",
			last:  "bond",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			first: "miss",
			last:  "Moneypenny",
		},
		ltk: false,
	}
	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()
}
