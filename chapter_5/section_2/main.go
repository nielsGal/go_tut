package main

import "fmt"

type person struct {
	first, last string //you can multi assign here too
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	s1 := secretAgent{
		person: person{
			first: "james",
			last:  "bond",
		},
		ltk: true,
	}
	fmt.Println(s1)
	fmt.Println(s1.first) //we auto destructre the struct
}
