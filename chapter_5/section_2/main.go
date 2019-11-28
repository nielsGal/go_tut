package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	s1 := secretAgent{
		person{
			"james",
			"bond",
		},
		true,
	}
	fmt.Println(s1)
	fmt.Println(s1.first)
}
