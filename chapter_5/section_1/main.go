package main

import "fmt"

type person struct {
	first string
	last  string
}

func main() {
	p1 := person{
		"james",
		"bond",
	}
	p2 := person{
		"miss",
		"moneypenny",
	}
	fmt.Println(p1.first, p2.last)
}
